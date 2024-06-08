package repository

import (
	"context"
	"msgb/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BudgetRepo struct {
	MongoCollection *mongo.Collection
}

func (r *BudgetRepo) InsertBudget(budget *model.Budget) (interface{}, error) {

	result, err := r.MongoCollection.InsertOne(context.Background(), budget)

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *BudgetRepo) UpdateBudget(budgetId string, budget *model.Budget) (int64, error) {

	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "budget_id", Value: budgetId}}, bson.D{{Key: "$set", Value: budget}})

	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil

}

func (r *BudgetRepo) GetBudget(budgetId string) (*model.Budget, error) {
	var budget *model.Budget

	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "budget_id", Value: budgetId}}).Decode(&budget)

	if err != nil {
		return nil, err
	}

	return budget, nil
}

func (r *BudgetRepo) GetBudgets() ([]*model.Budget, error) {

	var budgets []*model.Budget

	results, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	results.All(context.Background(), &budgets)

	return budgets, nil
}

func (r *BudgetRepo) DeleteBudget(budgetId string) (int64, error) {

	result, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "budget_id", Value: budgetId}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (r *BudgetRepo) DeleteBudgets() (int64, error) {

	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
