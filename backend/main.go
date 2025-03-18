package main

import (
	"context"
	"fmt"
	"net/http"
	"tokenSystem/controllers"
	"tokenSystem/routers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions, err := mongo.NewClient(options.Client().ApplyURI(controllers.ConnectionString))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = clientOptions.Connect(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	routers.Router()
	err = http.ListenAndServe(":8080", routers.Router())
	if err != nil {
		fmt.Println(err)
	}
}
