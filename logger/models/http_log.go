package models

import (
	"drp/logger/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HttpLog struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Dt           int64              `bson:"dt" json:"dt" validate:"required,gte=0"`
	User         string             `bson:"user" json:"user" validate:"required,min=1,max=256"`
	Method       string             `bson:"method" json:"method" validate:"required,oneof=GET POST PUT PATCH DELETE HEAD OPTIONS"`
	Uri          string             `bson:"uri" json:"uri" validate:"required,min=1,max=65536"`
	Body         string             `bson:"body" json:"body"`
	ResponseCode uint16             `bson:"response_code" json:"response_code" validate:"required,gte=100,lte=599"`
	Response     string             `bson:"response" json:"response"`
	RemoteIp     string             `bson:"remote_ip" json:"remote_ip" validate:"required,ip"`
	RefererUrl   string             `bson:"referer_url" json:"referer_url" validate:"required,min=1,max=65536"`
	RemoteAgent  string             `bson:"remote_agent" json:"remote_agent" validate:"required,min=1,max=65536"`
	ExecTime     float64            `bson:"exec_time" json:"exec_time" validate:"required,numeric,gte=0"`
	AppField     `bson:",inline"`
	Timestamps   `bson:",inline"`
}

func (m *HttpLog) GetCollection() *mongo.Collection {
	return database.GetDatabase().Collection("http_log")
}
