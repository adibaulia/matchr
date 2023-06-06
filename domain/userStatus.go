package domain

import "strings"

type UserStatus string

var (
	FREE   UserStatus = "FREE"
	MALE   string     = "M"
	FEMALE string     = "F"
)

func GetGender(gender string) string {
	switch strings.ToLower(gender) {
	case "male":
		return MALE
	case "female":
		return FEMALE
	default:
		return gender
	}
}
