package server_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/go-http-testing/examples/client"
	"github.com/ScarletTanager/go-http-testing/examples/server"
)

var _ = Describe("Server", func() {
	var (
		req     *http.Request
		handler http.Handler
		writer  *httptest.ResponseRecorder
	)

	Describe("HandleGET", func() {
		BeforeEach(func() {
			handler = http.HandlerFunc(server.HandleGET)
			req = httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Add(server.HEADER_KEY_X_ACCOUNT, "myaccount")
			writer = httptest.NewRecorder()
		})

		It("Processes a GET request successfully", func() {
			handler.ServeHTTP(writer, req)
			Expect(writer.Code).To(Equal(http.StatusOK))
		})

		Context("When we get the wrong method", func() {
			JustBeforeEach(func() {
				req.Method = http.MethodPost
			})

			It("Returns an internal server error", func() {
				handler.ServeHTTP(writer, req)
				Expect(writer.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})

	Describe("Server-side Request handling", func() {
		var (
			c   *client.MyApplicationClient
			srv *httptest.Server
		)

		BeforeEach(func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.Path).To(Equal("/resource"))
				w.WriteHeader(http.StatusCreated)
			})
		})

		JustBeforeEach(func() {
			srv = httptest.NewServer(handler)
			c = client.NewApplicationClient(srv.Client())
		})

		AfterEach(func() {
			srv.CloseClientConnections()
			srv.Close()
		})

		It("Checks the resource path", func() {
			c.PerformQuery()
		})

		Context("Use a different handler", func() {
			BeforeEach(func() {
				handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					Expect(r.Method).To(Equal(http.MethodPost))
					w.WriteHeader(http.StatusCreated)
				})
			})

			It("Checks the request method", func() {
				c.PerformQuery()
			})
		})
	})
})
