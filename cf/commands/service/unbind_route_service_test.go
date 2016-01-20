package service_test

import (
	"net/http"

	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/commands/service"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/errors"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/flags"

	testapi "github.com/cloudfoundry/cli/cf/api/fakes"
	fakerequirements "github.com/cloudfoundry/cli/cf/requirements/fakes"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"

	. "github.com/cloudfoundry/cli/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UnbindRouteService", func() {
	var (
		ui                      *testterm.FakeUI
		configRepo              core_config.Repository
		routeRepo               *testapi.FakeRouteRepository
		routeServiceBindingRepo *testapi.FakeRouteServiceBindingRepository

		cmd         command_registry.Command
		deps        command_registry.Dependency
		factory     *fakerequirements.FakeFactory
		flagContext flags.FlagContext

		fakeDomain models.DomainFields

		loginRequirement           requirements.Requirement
		domainRequirement          *fakerequirements.FakeDomainRequirement
		serviceInstanceRequirement *fakerequirements.FakeServiceInstanceRequirement
	)

	BeforeEach(func() {
		ui = &testterm.FakeUI{}

		configRepo = testconfig.NewRepositoryWithDefaults()
		routeRepo = &testapi.FakeRouteRepository{}
		repoLocator := deps.RepoLocator.SetRouteRepository(routeRepo)

		routeServiceBindingRepo = &testapi.FakeRouteServiceBindingRepository{}
		repoLocator = repoLocator.SetRouteServiceBindingRepository(routeServiceBindingRepo)

		deps = command_registry.Dependency{
			Ui:          ui,
			Config:      configRepo,
			RepoLocator: repoLocator,
		}

		cmd = &service.UnbindRouteService{}
		cmd.SetDependency(deps, false)

		flagContext = flags.NewFlagContext(cmd.MetaData().Flags)

		factory = &fakerequirements.FakeFactory{}

		loginRequirement = &passingRequirement{Name: "login-requirement"}
		factory.NewLoginRequirementReturns(loginRequirement)

		domainRequirement = &fakerequirements.FakeDomainRequirement{}
		factory.NewDomainRequirementReturns(domainRequirement)

		fakeDomain = models.DomainFields{
			Guid: "fake-domain-guid",
			Name: "fake-domain-name",
		}
		domainRequirement.GetDomainReturns(fakeDomain)

		serviceInstanceRequirement = &fakerequirements.FakeServiceInstanceRequirement{}
		factory.NewServiceInstanceRequirementReturns(serviceInstanceRequirement)
	})

	Describe("Requirements", func() {
		Context("when not provided exactly two args", func() {
			BeforeEach(func() {
				flagContext.Parse("domain-name")
			})

			It("fails with usage", func() {
				Expect(func() { cmd.Requirements(factory, flagContext) }).To(Panic())
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"FAILED"},
					[]string{"Incorrect Usage. Requires DOMAIN and SERVICE_INSTANCE as arguments"},
				))
			})
		})

		Context("when provided exactly two args", func() {
			BeforeEach(func() {
				flagContext.Parse("domain-name", "service-instance")
			})

			It("returns a LoginRequirement", func() {
				actualRequirements, err := cmd.Requirements(factory, flagContext)
				Expect(err).NotTo(HaveOccurred())
				Expect(factory.NewLoginRequirementCallCount()).To(Equal(1))
				Expect(actualRequirements).To(ContainElement(loginRequirement))
			})

			It("returns a DomainRequirement", func() {
				actualRequirements, err := cmd.Requirements(factory, flagContext)
				Expect(err).NotTo(HaveOccurred())
				Expect(factory.NewLoginRequirementCallCount()).To(Equal(1))
				Expect(actualRequirements).To(ContainElement(loginRequirement))
			})

			It("returns a ServiceInstanceRequirement", func() {
				actualRequirements, err := cmd.Requirements(factory, flagContext)
				Expect(err).NotTo(HaveOccurred())
				Expect(factory.NewServiceInstanceRequirementCallCount()).To(Equal(1))
				Expect(actualRequirements).To(ContainElement(serviceInstanceRequirement))
			})
		})
	})

	Describe("Execute", func() {
		BeforeEach(func() {
			err := flagContext.Parse("domain-name", "service-instance")
			Expect(err).NotTo(HaveOccurred())
			_, err = cmd.Requirements(factory, flagContext)
			Expect(err).NotTo(HaveOccurred())
		})

		It("tries to find the route", func() {
			ui.Inputs = []string{"n"}
			cmd.Execute(flagContext)
			Expect(routeRepo.FindCallCount()).To(Equal(1))
			host, domain, _, path := routeRepo.FindArgsForCall(0)
			Expect(host).To(Equal(""))
			Expect(domain).To(Equal(fakeDomain))
			Expect(path).To(Equal(""))
		})

		Context("when given a hostname", func() {
			BeforeEach(func() {
				flagContext = flags.NewFlagContext(cmd.MetaData().Flags)
				err := flagContext.Parse("domain-name", "service-instance", "-n", "the-hostname")
				Expect(err).NotTo(HaveOccurred())
			})

			It("tries to find the route with the given hostname", func() {
				ui.Inputs = []string{"n"}
				cmd.Execute(flagContext)
				Expect(routeRepo.FindCallCount()).To(Equal(1))
				host, _, _, _ := routeRepo.FindArgsForCall(0)
				Expect(host).To(Equal("the-hostname"))
			})
		})

		Context("when the route can be found", func() {
			BeforeEach(func() {
				routeRepo.FindReturns(models.Route{Guid: "route-guid"}, nil)
			})

			It("asks the user to confirm", func() {
				ui.Inputs = []string{"n"}
				cmd.Execute(flagContext)
				Expect(ui.Prompts).To(ContainSubstrings(
					[]string{"Unbinding may leave apps mapped to route", "Do you want to proceed?"},
				))
			})

			Context("when the user confirms", func() {
				JustBeforeEach(func() {
					defer func() { recover() }()
					ui.Inputs = []string{"y"}
					cmd.Execute(flagContext)
				})

				It("does not warn", func() {
					Expect(func() []string { return ui.Outputs }).NotTo(ContainSubstrings(
						[]string{"Unbind cancelled"},
					))
				})

				It("tells the user it is unbinding the route service", func() {
					Expect(ui.Outputs).To(ContainSubstrings(
						[]string{"Unbinding route", "from service instance"},
					))
				})

				It("tries to unbind the route service", func() {
					Expect(routeServiceBindingRepo.UnbindCallCount()).To(Equal(1))
				})

				Context("when unbinding the route service succeeds", func() {
					BeforeEach(func() {
						routeServiceBindingRepo.UnbindReturns(nil)
					})

					It("says OK", func() {
						Expect(ui.Outputs).To(ContainSubstrings(
							[]string{"OK"},
						))
					})
				})

				Context("when unbinding the route service fails because it was not bound", func() {
					BeforeEach(func() {
						routeServiceBindingRepo.UnbindReturns(errors.NewHttpError(http.StatusOK, errors.ROUTE_WAS_NOT_BOUND, "http-err"))
					})

					It("says OK", func() {
						Expect(ui.Outputs).To(ContainSubstrings(
							[]string{"OK"},
						))
					})

					It("warns", func() {
						Expect(ui.Outputs).To(ContainSubstrings(
							[]string{"Route", "was not bound to service instance"},
						))
					})
				})

				Context("when unbinding the route service fails for any other reason", func() {
					BeforeEach(func() {
						routeServiceBindingRepo.UnbindReturns(errors.New("unbind-err"))
					})

					It("fails with the error", func() {
						Expect(ui.Outputs).To(ContainSubstrings(
							[]string{"FAILED"},
							[]string{"unbind-err"},
						))
					})
				})
			})

			Context("when the user does not confirm", func() {
				BeforeEach(func() {
					ui.Inputs = []string{"n"}
					cmd.Execute(flagContext)
				})

				It("warns", func() {
					Expect(ui.Outputs).To(ContainSubstrings(
						[]string{"Unbind cancelled"},
					))
				})

				It("does not bind the route service", func() {
					Expect(routeServiceBindingRepo.UnbindCallCount()).To(Equal(0))
				})
			})

			Context("when the -f flag has been passed", func() {
				BeforeEach(func() {
					flagContext = flags.NewFlagContext(cmd.MetaData().Flags)
					flagContext.Parse("domain-name", "-f")
				})

				It("does not ask the user to confirm", func() {
					cmd.Execute(flagContext)
					Expect(ui.Prompts).NotTo(ContainSubstrings(
						[]string{"Unbinding may leave apps mapped to route", "Do you want to proceed?"},
					))
				})
			})
		})

		Context("when finding the route results in an error", func() {
			BeforeEach(func() {
				routeRepo.FindReturns(models.Route{Guid: "route-guid"}, errors.New("find-err"))
			})

			It("fails with error", func() {
				defer func() { recover() }()
				cmd.Execute(flagContext)
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"FAILED"},
					[]string{"find-err"},
				))
			})
		})
	})
})
