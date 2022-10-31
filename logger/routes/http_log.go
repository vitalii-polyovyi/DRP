package routes

import (
	"drp/logger/controllers"

	"github.com/go-chi/chi/v5"
)

type HttpLogRoutes struct{}

var httpLogController controllers.HttpLogController

func (rs HttpLogRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", httpLogController.List)
	r.Post("/", httpLogController.Create)
	r.Post("/bulk", httpLogController.CreateBulk)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(WithIdCtx)
		r.Get("/", httpLogController.Get)
		r.Put("/", httpLogController.Update)
		r.Delete("/", httpLogController.Delete)
	})

	return r
}
