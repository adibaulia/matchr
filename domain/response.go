package domain

type (
	Response struct {
		StatusCode int         `json:"status_code"`
		Status     string      `json:"status"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
	}
)
