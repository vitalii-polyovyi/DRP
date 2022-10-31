package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
)

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("POSTGRES_DATABASE"))

	web.Router("/", &MainController{})
	web.Run()
}
