package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectPostsDB() *mongo.Collection {
	err := godotenv.Load("helper/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	clientOptions := options.Client().ApplyURI("mongodb+srv://Admin:" + os.Getenv("DBKEY") + "@cluster0.rre0y.mongodb.net/Database?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Post Collection. ")

	collection := client.Database("go_rest_api").Collection("posts")

	return collection
}

func ConnectUsersDB() *mongo.Collection {

	clientOptions := options.Client().ApplyURI("mongodb+srv://Admin:" + os.Getenv("DBKEY") + "@cluster0.rre0y.mongodb.net/Database?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to User Collection. ")

	collection := client.Database("go_rest_api").Collection("users")

	return collection
}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
