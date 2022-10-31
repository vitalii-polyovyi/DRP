package routes

import (
	"drp/logger/controllers"

	"github.com/go-chi/chi/v5"
)

type EntityLogRoutes struct{}

var entityLogController controllers.EntityLogController

func (rs EntityLogRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", entityLogController.List)
	r.Post("/", entityLogController.Create)
	r.Post("/bulk", entityLogController.CreateBulk)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(WithIdCtx)
		r.Get("/", entityLogController.Get)
		r.Put("/", entityLogController.Update)
		r.Delete("/", entityLogController.Delete)
	})

	return r
}
