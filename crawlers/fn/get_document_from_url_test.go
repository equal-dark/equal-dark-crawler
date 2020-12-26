package fn_test

import (
	"equal_dark_crawler/crawlers/fn"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetDocument_WhenFailedGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	_, actual := fn.GetDocumentFromURL("http://failed_test")

	assert.NotNil(t, actual)
}

func Test_GetDocument_WhenStatusNot200(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	_, actual := fn.GetDocumentFromURL(server.URL)

	assert.NotNil(t, actual)
}

func Test_GetDocument_WithoutError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(nil)
	}))
	defer server.Close()

	_, actual := fn.GetDocumentFromURL(server.URL)

	assert.Nil(t, actual)
}
