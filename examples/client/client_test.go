package client_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/go-http-testing/examples/client"
	"github.com/ScarletTanager/go-http-testing/examples/client/clientfakes"
)

var _ = Describe("Client", func() {
	Describe("NewApplicationClient", func() {
		Context("When using a stock golang http client", func() {
			It("Successfully creates a new client", func() {
				Expect(client.NewApplicationClient(&http.Client{})).NotTo(BeNil())
			})
		})

		Context("When using a testing http client", func() {
			var (
				testingClient *http.Client
			)

			BeforeEach(func() {
				testingClient = httptest.NewServer(http.HandlerFunc(func(
					w http.ResponseWriter,
					r *http.Request) {
					// Noop implementation
				})).Client()
			})

			It("Successfully creates a new client", func() {
				Expect(client.NewApplicationClient(testingClient)).NotTo(BeNil())
			})
		})

		Context("When using a fake http client", func() {
			It("Successfully creates a new client", func() {
				Expect(client.NewApplicationClient(
					&clientfakes.FakeMyHttpClient{})).NotTo(BeNil())
			})
		})
	})
})
