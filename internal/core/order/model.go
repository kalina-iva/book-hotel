package order

import "time"

type Order struct {
	RoomID    int
	UserEmail string
	From      time.Time
	To        time.Time
}

type Hotel struct {
	ID   int
	Name string
}

type Room struct {
	ID      int
	HotelID int
	Type    string
}

type RoomAvailability struct {
	RoomID int
	Date   time.Time
	Quota  int
}
