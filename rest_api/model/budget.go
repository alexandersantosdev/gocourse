package model

import "time"

type Budget struct {
	BudgetId    string    `json:"budget_id,omitempty" bson:"budget_id"`
	Client      string    `json:"client,omitempty" bson:"client"`
	Phone       string    `json:"phone,omitempty" bson:"phone"`
	Observation string    `json:"observation,omitempty" bson:"observation"`
	Amount      int       `json:"amount,omitempty" bson:"amount"`
	Status      string    `json:"status,omitempty" bson:"status"`
	Items       []Item    `json:"items,omitempty" bson:"items"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at"`
	ExpiresAt   time.Time `json:"expires_at,omitempty" bson:"expires_at"`
}

type Item struct {
	Name     string `json:"name,omitempty" bson:"name"`
	Price    int    `json:"price,omitempty" bson:"price"`
	Quantity int    `json:"quantity,omitempty" bson:"quantity"`
}
