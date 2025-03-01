package tools

// Response Standar
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// type ErrorResponse struct {
// 	Message string `json:"message"`
// 	Status  string `json:"status"`

// }
