package models

import "go.mongodb.org/mongo-driver/mongo"

type Model interface {
	GetCollection() *mongo.Collection
	SetOnCreate()
	SetOnUpdate()
	SetApp(app string)
}
