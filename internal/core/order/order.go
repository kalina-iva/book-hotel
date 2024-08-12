package order

import (
	"time"
)

type Repository interface {
	CreateOrder(newOrder Order) error
}

type HotelRepository interface {
	GetRoomID(hotel, RoomType string) (int, error)
	GetAvailableDays(roomID int, from, to time.Time) ([]RoomAvailability, error)
	UpdateAvailability(updates []RoomAvailability) error
}

type Service struct {
	orderRepo Repository
	hotelRepo HotelRepository
}

func NewService(orderRepo Repository, hotelRepo HotelRepository) Service {
	return Service{
		orderRepo: orderRepo,
		hotelRepo: hotelRepo,
	}
}

func (s *Service) CreateOrder(
	hotel, room, userEmail string, from, to time.Time,
) error {
	roomID, err := s.hotelRepo.GetRoomID(hotel, room)
	if err != nil {
		return err
	}

	newOrder := Order{
		RoomID:    roomID,
		UserEmail: userEmail,
		From:      from,
		To:        to,
	}
	err = s.bookRoom(newOrder)
	if err != nil {
		return err
	}

	return s.orderRepo.CreateOrder(newOrder)
}

func (s *Service) bookRoom(newOrder Order) error {
	Availability, err := s.hotelRepo.GetAvailableDays(newOrder.RoomID, newOrder.From, newOrder.To)
	if err != nil {
		return err
	}

	daysToBook := daysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	updates := make([]RoomAvailability, 0, len(Availability))
	for _, dayToBook := range daysToBook {
		for i, availability := range Availability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			updates = append(updates, availability)
			Availability[i] = availability
			delete(unavailableDays, dayToBook)
			break
		}
	}

	if len(unavailableDays) != 0 {
		return ErrHotelRoomIsNotAvailable
	}

	return s.hotelRepo.UpdateAvailability(updates)
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := ToDay(from); !d.After(ToDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func ToDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
