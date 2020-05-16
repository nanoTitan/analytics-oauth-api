package rest_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/nanoTitan/analytics-oauth-api/src/repository/rest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	// block all HTTP requests
	client := rest.GetClient()
	httpmock.ActivateNonDefault(client.GetClient())
})

var _ = BeforeEach(func() {
	// remove any mocks
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})

func TestRest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rest Suite")
}
