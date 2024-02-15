package models

// Структура для запиту
type Request struct {
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

// Структура для відповіді
type Response struct {
	Greeting string  `json:"greeting"`
	Request  Request `json:"request"`
}
