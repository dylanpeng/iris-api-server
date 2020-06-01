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
}
