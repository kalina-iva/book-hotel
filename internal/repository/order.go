package repository

import (
	"book_hotel/internal/core/order"
	"book_hotel/internal/storage"
)

type OrderRepo struct {
	db storage.DB
}

func NewOrderRepo(db storage.DB) OrderRepo {
	return OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) CreateOrder(newOrder order.Order) error {
	r.db.Orders = append(r.db.Orders, newOrder)
	return nil
}
