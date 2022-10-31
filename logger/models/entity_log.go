package models

import (
	"drp/logger/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	EntityLogActionCreate string = "create"
	EntityLogActionUpdate string = "update"
	EntityLogActionDelete string = "delete"
)

type EntityLog struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Dt         int64              `bson:"dt" json:"dt" validate:"required,gte=0"`
	User       string             `bson:"user" json:"user" validate:"required,min=1,max=256"`
	Action     string             `bson:"action" json:"action" validate:"required,oneof=create update delete"`
	Entity     string             `bson:"entity" json:"entity" validate:"required,min=1,max=256"`
	RowId      string             `bson:"row_id" json:"row_id" validate:"required"`
	Field      string             `bson:"field" json:"field" validate:"required,min=1,max=256"`
	OldValue   string             `bson:"old_value" json:"old_value"`
	NewValue   string             `bson:"new_value" json:"new_value" validate:"required"`
	RequestId  string             `bson:"request_id" json:"request_id" validate:"required"`
	AppField   `bson:",inline"`
	Timestamps `bson:",inline"`
}

func (m *EntityLog) GetCollection() *mongo.Collection {
	return database.GetDatabase().Collection("entity_log")
}
