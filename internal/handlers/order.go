package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"book_hotel/internal/core/order"
	"book_hotel/internal/pkg/logger"
)

type Handler struct {
	orderService order.Service
}

func NewHandler(orderService order.Service) Handler {
	return Handler{
		orderService: orderService,
	}
}

type Order struct {
	HotelID   string `json:"hotel_id"`
	RoomID    string `json:"room_id"`
	UserEmail string `json:"email"`
	From      string `json:"from"`
	To        string `json:"to"`
}

func (h *Handler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var newOrder Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		setError(w, fmt.Sprintf("unable decode body: %s", err), http.StatusInternalServerError)
		return
	}

	// validate email
	from, to, err := validateTime(newOrder)
	if err != nil {
		setError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.orderService.CreateOrder(
		newOrder.HotelID, newOrder.RoomID, newOrder.UserEmail, *from, *to,
	)
	if err != nil {
		setError(w, fmt.Sprintf("Unable create order: %s", err), getStatusCodeByError(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	logger.LogInfo("Order successfully created: %v", newOrder)
}

func validateTime(newOrder Order) (*time.Time, *time.Time, error) {
	from, err := time.Parse(time.DateOnly, newOrder.From)
	if err != nil {
		return nil, nil, fmt.Errorf("wrong from date: %v", err)
	}
	to, err := time.Parse(time.DateOnly, newOrder.To)
	if err != nil {
		return nil, nil, fmt.Errorf("wrong to date: %v", err)
	}
	if from.After(to) {
		return nil, nil, errors.New("date after")
	}
	return &from, &to, nil
}

func getStatusCodeByError(err error) int {
	code := http.StatusInternalServerError
	if errors.Is(err, order.ErrHotelRoomIsNotAvailable) {
		code = http.StatusConflict
	} else if errors.Is(err, order.ErrUnknownHotel) || errors.Is(err, order.ErrUnknownRoomType) {
		code = http.StatusBadRequest
	}
	return code
}
