package password_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	testapi "github.com/fujitsu-cf/cli/cf/api/fakes"
	"github.com/fujitsu-cf/cli/cf/net"
	testconfig "github.com/fujitsu-cf/cli/testhelpers/configuration"
	testnet "github.com/fujitsu-cf/cli/testhelpers/net"
	testterm "github.com/fujitsu-cf/cli/testhelpers/terminal"

	. "github.com/fujitsu-cf/cli/cf/api/password"
	. "github.com/fujitsu-cf/cli/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CloudControllerPasswordRepository", func() {
	It("updates your password", func() {
		req := testapi.NewCloudControllerTestRequest(testnet.TestRequest{
			Method:   "PUT",
			Path:     "/Users/my-user-guid/password",
			Matcher:  testnet.RequestBodyMatcher(`{"password":"new-password","oldPassword":"old-password"}`),
			Response: testnet.TestResponse{Status: http.StatusOK},
		})

		passwordUpdateServer, handler, repo := createPasswordRepo(req)
		defer passwordUpdateServer.Close()

		apiErr := repo.UpdatePassword("old-password", "new-password")
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(apiErr).NotTo(HaveOccurred())
	})
})

func createPasswordRepo(req testnet.TestRequest) (passwordServer *httptest.Server, handler *testnet.TestHandler, repo PasswordRepository) {
	passwordServer, handler = testnet.NewServer([]testnet.TestRequest{req})

	configRepo := testconfig.NewRepositoryWithDefaults()
	configRepo.SetUaaEndpoint(passwordServer.URL)
	gateway := net.NewCloudControllerGateway(configRepo, time.Now, &testterm.FakeUI{})
	repo = NewCloudControllerPasswordRepository(configRepo, gateway)
	return
}
