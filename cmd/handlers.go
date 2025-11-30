package main

import "github.com/momomo0206/go-pizza-tracker/internal/models"

type Hanndler struct {
	orders *models.OrderModel
}

func NewHanndler(dbModel *models.DBModel) *Hanndler {
	return &Hanndler{
		orders: &dbModel.Order,
	}
}
