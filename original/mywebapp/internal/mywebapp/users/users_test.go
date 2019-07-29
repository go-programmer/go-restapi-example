package users_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"mywebapp/internal/mywebapp"
	"mywebapp/internal/mywebapp/users"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMyWebApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "test suite for mywebapp")
}

var _ = Describe("my web app", func() {
	var server *httptest.Server
	BeforeEach(func() {
		server = httptest.NewServer(mywebapp.NewController())
	})
	AfterEach(func() {
		server.Close()
	})

	Describe("the standard user lifecycle", func() {
		var user = users.User{
			Id:    string(rand.Int31()),
			Email: "thisisanemail@domain.com",
		}
		userJson, _ := json.Marshal(user)

		Describe("creating a user", func() {
			var (
				resp *http.Response
				err  error
			)

			BeforeEach(func() {
				resp, err = http.Post(fmt.Sprintf("%s/users", server.URL), "application/json", bytes.NewReader(userJson))
			})

			It("should return a 201", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(BeIdenticalTo(http.StatusCreated))
			})
		})

		Describe("getting the user", func() {
			var (
				resp *http.Response
				err  error
			)

			BeforeEach(func() {
				resp, err = http.Get(fmt.Sprintf("%s/users/%s", server.URL, user.Id))
			})

			It("should return a 200", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(BeIdenticalTo(http.StatusOK))
			})

			It("should return the JSON encoded user as its body", func() {
				decodedUser := users.User{}
				err := json.NewDecoder(resp.Body).Decode(&decodedUser)
				Expect(err).ToNot(HaveOccurred())
				Expect(decodedUser).To(BeIdenticalTo(user))
			})
		})

		Describe("deleting the user", func() {
			var (
				resp *http.Response
				err  error
			)

			BeforeEach(func() {
				req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/users/%s", server.URL, user.Id), nil)
				Expect(err).ToNot(HaveOccurred())
				resp, err = http.DefaultClient.Do(req)
			})

			It("should return a 204", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(BeIdenticalTo(http.StatusNoContent))
			})
		})
	})
})
