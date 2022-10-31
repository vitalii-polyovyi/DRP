package models

import (
	"drp/logger/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	AppStatusActive   string = "active"
	AppStatusDisabled string = "disabled"
)

type App struct {
	Id         primitive.ObjectID `bson:"_id"`
	Dt         int64              `bson:"dt" validate:"required"`
	Status     string             `bson:"status" validate:"required,oneof: active disabled"`
	Key        string             `bson:"key" validate:"required"`
	AppField   `bson:",inline"`
	Timestamps `bson:",inline"`
}

func (m *App) GetCollection() *mongo.Collection {
	return database.GetDatabase().Collection("apps")
}
