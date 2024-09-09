package web

type WebSuccessResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload T      `json:"payload"`
}
