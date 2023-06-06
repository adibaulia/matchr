package domain

type (
	UserRequest struct {
		Username    string `json:"username"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		Name        string `json:"name"`
		DateOfBirth string `json:"date_of_birth"`
		Gender      string `json:"gender"`
		Bio         string `json:"bio"`
	}
)
