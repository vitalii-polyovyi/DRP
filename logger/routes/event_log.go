package routes

import (
	"drp/logger/controllers"

	"github.com/go-chi/chi/v5"
)

type EventLogRoutes struct{}

var eventLogController controllers.EventLogController

func (el EventLogRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", eventLogController.List)
	r.Post("/", eventLogController.Create)
	r.Post("/bulk", eventLogController.CreateBulk)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(WithIdCtx)
		r.Get("/", eventLogController.Get)
		r.Put("/", eventLogController.Update)
		r.Delete("/", eventLogController.Delete)
	})

	return r
}
