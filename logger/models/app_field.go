package models

type AppField struct {
	App string `bson:"app" json:"app"`
}

func (a *AppField) SetApp(app string) {
	a.App = app
}
