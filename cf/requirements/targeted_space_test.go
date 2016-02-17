package requirements_test

import (
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	"github.com/fujitsu-cf/cli/cf/models"
	. "github.com/fujitsu-cf/cli/cf/requirements"
	testassert "github.com/fujitsu-cf/cli/testhelpers/assert"
	testconfig "github.com/fujitsu-cf/cli/testhelpers/configuration"
	testterm "github.com/fujitsu-cf/cli/testhelpers/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fujitsu-cf/cli/testhelpers/matchers"
)

var _ = Describe("TargetedSpaceRequirement", func() {
	var (
		ui     *testterm.FakeUI
		config core_config.ReadWriter
	)

	BeforeEach(func() {
		ui = new(testterm.FakeUI)
		config = testconfig.NewRepositoryWithDefaults()
	})

	Context("when the user has targeted a space", func() {
		It("succeeds", func() {
			req := NewTargetedSpaceRequirement(ui, config)
			Expect(req.Execute()).To(BeTrue())
		})
	})

	Context("when the user does not have a space targeted", func() {
		It("fails", func() {
			config.SetSpaceFields(models.SpaceFields{})

			testassert.AssertPanic(testterm.QuietPanic, func() {
				NewTargetedSpaceRequirement(ui, config).Execute()
			})

			Expect(ui.Outputs).To(ContainSubstrings(
				[]string{"FAILED"},
				[]string{"No space targeted"},
			))
		})
	})
})
