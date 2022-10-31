package models

import (
	"drp/logger/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventLog struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Dt         int64              `bson:"dt" json:"dt" validate:"required,gte=0"`
	User       string             `bson:"user" json:"user" validate:"required,min=1,max=256"`
	Event      string             `bson:"event" json:"event" validate:"required,min=1,max=256"`
	Context    bson.M             `bson:"context" json:"context" validate:"required"`
	Tags       []string           `bson:"tags" json:"tags" validate:"required"`
	AppField   `bson:",inline"`
	Timestamps `bson:",inline"`
}

func (m *EventLog) GetCollection() *mongo.Collection {
	return database.GetDatabase().Collection("event_log")
}
