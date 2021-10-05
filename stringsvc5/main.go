package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	var logger log.Logger = log.NewLogfmtLogger(os.Stderr)
	var svc StringService = stringService{}
	var h http.Handler = MakeHTTPHandler(svc)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", h))
}
