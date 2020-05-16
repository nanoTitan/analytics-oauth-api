package rest_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/nanoTitan/analytics-oauth-api/src/repository/rest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var _ = Describe("TestLoginUser", func() {
	It("tests that the user login errors with internal server error", func() {
		msg := "error when trying to unmarshal users response"
		body := `{"Status": 500, "Message": ` + msg + `}`
		responder := httpmock.NewStringResponder(0, body)
		httpmock.RegisterResponder(
			"POST",
			rest.GetBaseURL()+"/users/login",
			responder)

		repository := rest.New()
		user, err := repository.LoginUser("email@gmail.com", "password")

		Expect(user).To(BeNil())
		Expect(err).NotTo(BeNil())
		Expect(err.Status).To(Equal(http.StatusInternalServerError))
		Expect(err.Message).To(Equal(msg))
	})

	It("tests that the user login errors with invalid interface", func() {
		responder := httpmock.NewStringResponder(300, ``)
		httpmock.RegisterResponder(
			"POST",
			rest.GetBaseURL()+"/users/login",
			responder)

		repository := rest.New()
		user, err := repository.LoginUser("email@gmail.com", "password")

		Expect(user).To(BeNil())
		Expect(err).NotTo(BeNil())
		Expect(err.Status).To(Equal(http.StatusInternalServerError))
		Expect(err.Message).To(Equal("invalid error interface when trying to login user"))
	})

	It("tests returns error when status > 299 and can't unmarshall request body", func() {
		responder := httpmock.NewStringResponder(300, ``)
		httpmock.RegisterResponder(
			"POST",
			rest.GetBaseURL()+"/users/login",
			responder)

		repository := rest.New()
		user, err := repository.LoginUser("email@gmail.com", "password")

		fmt.Println(user)
		fmt.Println(err)

		Expect(user).To(BeNil())
		Expect(err).NotTo(BeNil())
		Expect(err.Status).To(Equal(http.StatusInternalServerError))
		Expect(err.Message).To(Equal("invalid error interface when trying to login user"))
	})

	It("tests that the user login errors during unmarshal with invalid User object", func() {
		responder := httpmock.NewStringResponder(200, ``)
		httpmock.RegisterResponder(
			"POST",
			rest.GetBaseURL()+"/users/login",
			responder)

		repository := rest.New()
		user, err := repository.LoginUser("email@gmail.com", "password")

		Expect(user).To(BeNil())
		Expect(err).NotTo(BeNil())
		Expect(err.Status).To(Equal(http.StatusInternalServerError))
		Expect(err.Message).To(Equal("error when trying to unmarshal users response"))
	})

	It("tests returns invalid login credentials", func() {
		body := `{"message": "invalid login credentials", "status": 404, "error": "not_found"}`
		responder := httpmock.NewStringResponder(404, body)
		httpmock.RegisterResponder(
			"POST",
			rest.GetBaseURL()+"/users/login",
			responder)

		repository := rest.New()
		user, err := repository.LoginUser("email@gmail.com", "password")

		Expect(user).To(BeNil())
		Expect(err).NotTo(BeNil())
		Expect(err.Status).To(Equal(http.StatusNotFound))
		Expect(err.Message).To(Equal("invalid login credentials"))
	})

	It("tests that the user is created successfully", func() {
		response := `
		{
			"id": 123,
			"first_name": "John",
			"last_name": "Doe",
			"email": "johndoe@test.com",
			"date_created": "01/01/2000",
			"status": "active",
			"password": "abc123"
		}`
		responder := httpmock.NewStringResponder(200, response)
		httpmock.RegisterResponder("POST", rest.GetBaseURL()+"/users/login", responder)

		repository := rest.New()
		user, err := repository.LoginUser("email@gmail.com", "password")
		Expect(err).To(BeNil())
		userId := int64(123)
		Expect(user.Id).To(Equal(userId))
		Expect(user.FirstName).To(Equal("John"))
		Expect(user.LastName).To(Equal("Doe"))
		Expect(user.Email).To(Equal("johndoe@test.com"))
		Expect(user.Status).To(Equal("active"))
	})
})
