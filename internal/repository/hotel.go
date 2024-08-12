package repository

import (
	"time"

	"book_hotel/internal/core/order"
	"book_hotel/internal/storage"
)

type HotelRepo struct {
	db storage.DB
}

func NewRepo(db storage.DB) HotelRepo {
	return HotelRepo{
		db: db,
	}
}

func (r *HotelRepo) GetAvailableDays(roomID int, from, to time.Time) ([]order.RoomAvailability, error) {
	res := make([]order.RoomAvailability, 0)
	for _, r := range r.db.Availability {
		if r.RoomID != roomID {
			continue
		}
		if r.Date.Equal(from) || r.Date.Equal(to) || r.Date.After(from) && r.Date.Before(to) {
			res = append(res, r)
		}
	}
	return res, nil
}

func (r *HotelRepo) UpdateAvailability(updates []order.RoomAvailability) error {
	for _, n := range updates {
		for i, a := range r.db.Availability {
			if a.RoomID == n.RoomID && a.Date.Equal(n.Date) {
				r.db.Availability[i].Quota = n.Quota
				break
			}
		}
	}

	return nil
}

func (r *HotelRepo) GetRoomID(hotel, RoomType string) (int, error) {
	hotelID := 0
	for _, h := range r.db.Hotels {
		if h.Name == hotel {
			hotelID = h.ID
			break
		}
	}
	if hotelID <= 0 {
		return 0, order.ErrUnknownHotel
	}

	roomID := 0
	for _, r := range r.db.Rooms {
		if r.HotelID == hotelID && r.Type == RoomType {
			roomID = r.ID
			break
		}
	}
	var err error
	if roomID <= 0 {
		err = order.ErrUnknownRoomType
	}
	return roomID, err
}
