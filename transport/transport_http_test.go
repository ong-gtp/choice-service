package transport

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/log"
	"github.com/ong-gtp/choice-service/service"
)

type serviceRequest struct {
	method, url, body string
	want              int
}

func TestHTTP(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "env", "testing", "caller", log.DefaultCaller)

	s := service.NewService()
	r := NewHttpServer(s, logger)
	srv := httptest.NewServer(r)

	srs := []serviceRequest{
		{method: "GET", url: srv.URL + "/", body: "", want: http.StatusOK},
		{method: "GET", url: srv.URL + "/health", body: "", want: http.StatusOK},
		{method: "GET", url: srv.URL + "/v1", body: "", want: http.StatusOK},
		{method: "GET", url: srv.URL + "/v1/choices", body: "", want: http.StatusOK},
		{method: "POST", url: srv.URL + "/v1/choices", body: "", want: http.StatusMethodNotAllowed},
		{method: "POST", url: srv.URL + "/v1/choice", body: "", want: http.StatusMethodNotAllowed},
		{method: "GET", url: srv.URL + "/v1/choice", body: "", want: http.StatusOK},
	}

	for _, testcase := range srs {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, _ := http.DefaultClient.Do(req)
		if testcase.want != resp.StatusCode {
			t.Errorf("%s %s: want %d have %d", testcase.method, testcase.url, testcase.want, resp.StatusCode)
		}
	}
}
