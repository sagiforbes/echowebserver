package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Use(httpEchoMiddle())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func httpEchoMiddle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return reqHandler
	}
}

type echoRespons struct {
	ReqTime     time.Time
	CallerAddr  string
	Proto       string
	Method      string
	Path        string
	Uri         string
	Queryvalues url.Values
	Headers     http.Header
}

func reqHandler(c echo.Context) error {
	var req *http.Request = c.Request()
	ret := &echoRespons{
		ReqTime:     time.Now(),
		Proto:       req.Proto,
		CallerAddr:  req.RemoteAddr,
		Method:      req.Method,
		Path:        req.URL.Path,
		Uri:         req.RequestURI,
		Queryvalues: req.URL.Query(),
		Headers:     req.Header,
	}
	req.URL.Query()

	return c.JSONPretty(http.StatusOK, ret, "  ")
}
