package domain

type (
	Response struct {
		StatusCode int         `json:"status_code"`
		Status     string      `json:"status,omitempty"`
		Message    string      `json:"message,omitempty"`
		Data       interface{} `json:"data,omitempty"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)
