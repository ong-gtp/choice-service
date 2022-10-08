package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/ong-gtp/choice-service/logging"
	"github.com/ong-gtp/choice-service/service"
	"github.com/ong-gtp/choice-service/transport"
)

func main() {
	var logger log.Logger
	port := "8081"
	httpTimeout := 60 * time.Second
	env := "development"
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "env", env, "listen", port, "caller", log.DefaultCaller)

	svc := logging.NewLoggingMiddleware(logger, service.NewService())
	r := transport.NewHttpServer(svc, logger)
	rWithTimeout := http.TimeoutHandler(r, httpTimeout, "Timeout!")

	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(":"+port, rWithTimeout))
}
