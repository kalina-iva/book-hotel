package order

import "errors"

var (
	ErrHotelRoomIsNotAvailable = errors.New("hotel room is not available for selected dates")
	ErrUnknownHotel            = errors.New("unknown hotel")
	ErrUnknownRoomType         = errors.New("unknown room type")
)
