package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/gommon/random"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func (a *API) registerRoutes() {
	a.web.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			id := random.String(32)
			req.Header.Set(echo.HeaderXRequestID, id)
			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			return err
		}
	})
	a.web.Pre(middleware.RemoveTrailingSlash())
	a.web.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = random.String(32)
			}
			res.Header().Set(echo.HeaderXRequestID, id)

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			p := req.URL.Path
			bytesIn := req.Header.Get(echo.HeaderContentLength)

			a.log.WithFields(logrus.Fields{
				"id":            id,
				"remote_ip":     c.RealIP(),
				"host":          req.Host,
				"uri":           req.RequestURI,
				"method":        req.Method,
				"path":          p,
				"referer":       req.Referer(),
				"user_agent":    req.UserAgent(),
				"status":        res.Status,
				"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
				"latency_human": stop.Sub(start).String(),
				"bytes_in":      bytesIn,
				"bytes_out":     strconv.FormatInt(res.Size, 10),
			}).Info()
			return err
		}
	})
	a.web.Use(middleware.CORS())
	a.web.Use(middleware.Recover())

	g := a.web.Group("/v1")
	g.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})

	category := g.Group("/categories")
	category.GET("/:id", a.GetCategory)
}
