package main

import (
	"math/rand"
	"time"

	prometheusMiddleware "github.com/iris-contrib/middleware/prometheus"
	"github.com/kataras/iris"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	app := iris.New()
	m := prometheusMiddleware.New("serviceName", 300, 1200, 5000)

	app.Use(m.ServeHTTP)

	s := NewCollector()
	prometheus.MustRegister(s)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		m.ServeHTTP(ctx)
		ctx.Writef("Not Found")
	})

	app.Get("/", func(ctx iris.Context) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ctx.Writef("Slept for %d milliseconds", sleep)
	})

	app.Get("/metrics", iris.FromStd(prometheus.Handler()))

	app.Run(iris.Addr(":8080"))
}
