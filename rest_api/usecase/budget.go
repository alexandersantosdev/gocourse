package usecase

import (
	"encoding/json"
	"log"
	"msgb/model"
	"msgb/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type BudgetService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (s *BudgetService) GetBudgets(w http.ResponseWriter, r *http.Request) {}
func (s *BudgetService) GetBudget(w http.ResponseWriter, r *http.Request)  {}
func (s *BudgetService) InsertBudget(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	var budget model.Budget

	err := json.NewDecoder(r.Body).Decode(&budget)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}

	budget.BudgetId = uuid.New().String()
	budget.CreatedAt = time.Now()
	budget.ExpiresAt = time.Now().Add(10 * 24 * time.Hour)

	var amount int

	for _, item := range budget.Items {
		amount += item.Price * item.Quantity
	}

	budget.Amount = amount

	repo := repository.BudgetRepo{MongoCollection: s.MongoCollection}

	result, err := repo.InsertBudget(&budget)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error inserting budget", err)
		res.Error = err.Error()
		return
	}

	res.Data = budget.BudgetId
	w.WriteHeader(http.StatusCreated)
	log.Println("Budget inserted", result, budget)
}

func (s *BudgetService) UpdateBudget(w http.ResponseWriter, r *http.Request)  {}
func (s *BudgetService) DeleteBudget(w http.ResponseWriter, r *http.Request)  {}
func (s *BudgetService) DeleteBudgets(w http.ResponseWriter, r *http.Request) {}
