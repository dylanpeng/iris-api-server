package router

import (
	"github.com/kataras/iris"
	ctrl "juggernaut/gateway/logic/control"
	"juggernaut/gateway/logic/middleware"
)

func (r *router) RegHttpHandler(app *iris.Application) {
	app.Any("/health", ctrl.Health)

	app.Options("/{route:path}", middleware.CrossDomain)

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
