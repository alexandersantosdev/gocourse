package main

import (
	"context"
	"fmt"
	"log"
	"msgb/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	err = mongoClient.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Pinged MongoDB!")

}

func main() {
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	budgetService := usecase.BudgetService{MongoCollection: coll}

	r := mux.NewRouter()

	r.HandleFunc("/ping", pingHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/budgets", budgetService.GetBudgets).Methods(http.MethodGet)
	r.HandleFunc("/budgets/{budgetId}", budgetService.GetBudget).Methods(http.MethodGet)
	r.HandleFunc("/budgets", budgetService.InsertBudget).Methods(http.MethodPost)
	r.HandleFunc("/budgets/{budgetId}", budgetService.UpdateBudget).Methods(http.MethodPut)
	r.HandleFunc("/budgets/{budgetId}", budgetService.DeleteBudget).Methods(http.MethodDelete)
	r.HandleFunc("/budgets", budgetService.DeleteBudgets).Methods(http.MethodDelete)

	PORT := os.Getenv("PORT")

	log.Println("Listening on port " + PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
