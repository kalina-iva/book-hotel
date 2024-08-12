package storage

import (
	"time"

	"book_hotel/internal/core/order"
)

type DB struct {
	Hotels       []order.Hotel
	Rooms        []order.Room
	Availability []order.RoomAvailability
	Orders       []order.Order
}

func NewDB() DB {
	return DB{
		Hotels:       generateHotels(),
		Rooms:        generateRooms(),
		Availability: generateRoomAvailability(),
		Orders:       []order.Order{},
	}
}

func generateHotels() []order.Hotel {
	return []order.Hotel{
		{1, "reddison"},
	}
}

func generateRooms() []order.Room {
	return []order.Room{
		{1, 1, "lux"},
	}
}

func generateRoomAvailability() []order.RoomAvailability {
	return []order.RoomAvailability{
		{1, date(2024, 1, 1), 1},
		{1, date(2024, 1, 2), 1},
		{1, date(2024, 1, 3), 0},
		{1, date(2024, 1, 4), 1},
		{1, date(2024, 1, 5), 0},
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
