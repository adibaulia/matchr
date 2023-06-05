// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package generated

import (
	"time"
)

const TableNameUserStatus = "user_status"

// UserStatus mapped from table <user_status>
type UserStatus struct {
	ID          string     `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string     `gorm:"column:name;type:character varying(255);not null" json:"name"`
	Description *string    `gorm:"column:description;type:text" json:"description"`
	Price       float64    `gorm:"column:price;type:numeric(10,2);not null" json:"price"`
	Premium     bool       `gorm:"column:premium;type:boolean;not null" json:"premium"`
	CreatedDate *time.Time `gorm:"column:created_date;type:timestamp with time zone;default:now()" json:"created_date"`
	UpdatedDate *time.Time `gorm:"column:updated_date;type:timestamp with time zone;default:now()" json:"updated_date"`
}

// TableName UserStatus's table name
func (*UserStatus) TableName() string {
	return TableNameUserStatus
}