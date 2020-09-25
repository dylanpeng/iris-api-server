package router

import (
	"github.com/kataras/iris"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"juggernaut/common"
	ctrl "juggernaut/gateway/logic/control"
	"juggernaut/gateway/logic/middleware"
	"juggernaut/lib/http"
)

func (r *router) RegHttpHandler(app *iris.Application) {
	app.Use(common.HttpInterceptor)
	app.Any("/health", ctrl.Health)

	app.Options("/{route:path}", middleware.CrossDomain)
	app.Get("/metrics", http.UnGzip, iris.FromStd(promhttp.Handler()))

	// user group
	userParty := app.Party("/user")
	{
		userParty.Post("/login", ctrl.User.Login)
	}

	// user group
	kafkaParty := app.Party("/kafka")
	{
		kafkaParty.Post("/push", ctrl.Message.PushKafkaMessage)
	}

	// test grpc group
	testGrpcParty := app.Party("/test")
	{
		testGrpcParty.Post("/grpc/push", ctrl.TestGrpc.Push)
	}
}
