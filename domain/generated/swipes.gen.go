// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package generated

import (
	"time"
)

const TableNameSwipe = "swipes"

// Swipe mapped from table <swipes>
type Swipe struct {
	ID             string     `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	SwiperID       *string    `gorm:"column:swiper_id;type:uuid" json:"swiper_id"`
	SwipedID       *string    `gorm:"column:swiped_id;type:uuid" json:"swiped_id"`
	SwipeDirection *bool      `gorm:"column:swipe_direction;type:boolean" json:"swipe_direction"`
	SwipeDate      *time.Time `gorm:"column:swipe_date;type:timestamp with time zone;default:now()" json:"swipe_date"`
	CreatedDate    *time.Time `gorm:"column:created_date;type:timestamp with time zone;default:now()" json:"created_date"`
	UpdatedDate    *time.Time `gorm:"column:updated_date;type:timestamp with time zone;default:now()" json:"updated_date"`
}

// TableName Swipe's table name
func (*Swipe) TableName() string {
	return TableNameSwipe
}
