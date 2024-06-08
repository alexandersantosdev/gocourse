package repository

import (
	"context"
	"log"
	"msgb/model"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUrl string
var databaseName string
var collectionName string

func newMongoClient() *mongo.Client {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	err = mongoClient.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Pinged MongoDB!")

	return mongoClient

}

func TestCaseMongoOperations(t *testing.T) {

	initEnvs()

	mongoClient := newMongoClient()

	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(databaseName).Collection(collectionName)

	budRepo := BudgetRepo{MongoCollection: coll}

	budgetId := uuid.New().String()

	t.Run("Insert budget", func(t *testing.T) {
		budget1Model := model.Budget{
			Client:      "Teste update",
			Amount:      500,
			Phone:       "123456789",
			Observation: "Teste",
			Status:      "open",
			BudgetId:    budgetId,
			CreatedAt:   time.Now(),
			ExpiresAt:   time.Now().Add(10 * 24 * time.Hour),
			Items: []model.Item{
				{
					Name:     "Item 1",
					Price:    100,
					Quantity: 1,
				},

				{
					Name:     "Item 2",
					Price:    200,
					Quantity: 2,
				},
			},
		}

		result, err := budRepo.InsertBudget(&budget1Model)

		if err != nil {
			t.Fatal("Error inserting budget", err)
		}

		t.Log("Inserted budget", result)
	})

	t.Run("Get budget", func(t *testing.T) {

		budget, err := budRepo.GetBudget(budgetId)

		if err != nil {
			t.Fatal("Error getting budget", err)
		}

		t.Log("Get budget", budget.Client)
	})

	t.Run("Update budget", func(t *testing.T) {

		budget, err := budRepo.GetBudget(budgetId)

		if err != nil {
			t.Fatal("Error getting budget", err)
		}

		budget.Status = "closed"

		update, err := budRepo.UpdateBudget(budgetId, budget)

		if err != nil {
			t.Fatal("Error updating budget", err)
		}

		t.Log("Updated budget", update)

	})

	t.Run("Delete budget", func(t *testing.T) {

		delete, err := budRepo.DeleteBudget(budgetId)

		if err != nil {
			t.Fatal("Error deleting budget", err)
		}

		t.Log("Deleted budget", delete)
	})

	t.Run("Delete All budgets", func(t *testing.T) {

		delete, err := budRepo.DeleteBudgets()

		if err != nil {
			t.Fatal("Error deleting budget", err)
		}

		t.Log("Deleted budget", delete)

	})

}

func initEnvs() {
	godotenv.Load("../.env")
	mongoUrl = os.Getenv("MONGO_URL")
	databaseName = os.Getenv("DATABASE_NAME")
	collectionName = os.Getenv("TEST_COLLECTION_NAME")
}
