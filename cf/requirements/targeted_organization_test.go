package requirements_test

import (
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	"github.com/fujitsu-cf/cli/cf/models"

	testassert "github.com/fujitsu-cf/cli/testhelpers/assert"
	testconfig "github.com/fujitsu-cf/cli/testhelpers/configuration"
	testterm "github.com/fujitsu-cf/cli/testhelpers/terminal"

	. "github.com/fujitsu-cf/cli/cf/requirements"
	. "github.com/fujitsu-cf/cli/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TargetedOrganizationRequirement", func() {
	var (
		ui     *testterm.FakeUI
		config core_config.ReadWriter
	)

	BeforeEach(func() {
		ui = new(testterm.FakeUI)
		config = testconfig.NewRepositoryWithDefaults()
	})

	Context("when the user has an org targeted", func() {
		It("succeeds", func() {
			req := NewTargetedOrgRequirement(ui, config)
			success := req.Execute()
			Expect(success).To(BeTrue())
		})
	})

	Context("when the user does not have an org targeted", func() {
		It("fails", func() {
			config.SetOrganizationFields(models.OrganizationFields{})

			testassert.AssertPanic(testterm.QuietPanic, func() {
				NewTargetedOrgRequirement(ui, config).Execute()
			})

			Expect(ui.Outputs).To(ContainSubstrings(
				[]string{"FAILED"},
				[]string{"No org targeted"},
			))
		})
	})
})
