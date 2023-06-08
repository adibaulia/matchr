package domain

import "time"

type (
	Swipe struct {
		ID             string    `json:"id"`
		SwiperID       string    `json:"swiper_id"`
		SwipedID       string    `json:"swiped_id"`
		SwipeDirection bool      `json:"swipe_direction"`
		SwipeDate      time.Time `json:"swipe_date"`
		CreatedDate    time.Time `json:"created_date"`
		UpdatedDate    time.Time `json:"updated_date"`
	}

	SwipeUsecase interface {
		
	}
)
