package requirements_test

import (
	testapi "github.com/fujitsu-cf/cli/cf/api/fakes"
	"github.com/fujitsu-cf/cli/cf/models"
	. "github.com/fujitsu-cf/cli/cf/requirements"
	testassert "github.com/fujitsu-cf/cli/testhelpers/assert"
	testterm "github.com/fujitsu-cf/cli/testhelpers/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BuildpackRequirement", func() {
	var (
		ui *testterm.FakeUI
	)

	BeforeEach(func() {
		ui = new(testterm.FakeUI)
	})

	It("succeeds when a buildpack with the given name exists", func() {
		buildpack := models.Buildpack{Name: "my-buildpack"}
		buildpackRepo := &testapi.FakeBuildpackRepository{FindByNameBuildpack: buildpack}

		buildpackReq := NewBuildpackRequirement("my-buildpack", ui, buildpackRepo)

		Expect(buildpackReq.Execute()).To(BeTrue())
		Expect(buildpackRepo.FindByNameName).To(Equal("my-buildpack"))
		Expect(buildpackReq.GetBuildpack()).To(Equal(buildpack))
	})

	It("fails when the buildpack cannot be found", func() {
		buildpackRepo := &testapi.FakeBuildpackRepository{FindByNameNotFound: true}

		testassert.AssertPanic(testterm.QuietPanic, func() {
			NewBuildpackRequirement("foo", ui, buildpackRepo).Execute()
		})
	})
})
