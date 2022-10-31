package middlewares

import (
	"context"
	"drp/logger/helpers"
	"drp/logger/models"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var app models.App
		key := r.Header.Get("Authorization")
		condition := bson.M{"key": key, "status": models.AppStatusActive}
		err := new(models.App).GetCollection().FindOne(context.Background(), condition).Decode(&app)
		if err != nil {
			b, _ := json.Marshal(&helpers.ErrorResponse{Error: "Unauthorized"})
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(b)

			log.Println(err)

			return
		}

		ctx := context.WithValue(r.Context(), "app", app.App)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
