package respwriter

import "time"

type Module struct {
	start time.Time
}

type Response struct {
	ProcessTime float64     `json:"processTime,omitempty"`
	IsSuccess   bool        `json:"success"`
	Data        interface{} `json:"data"`
	Error       interface{} `json:"error"`
}

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Traces  []string `json:"traces,omitempty"`
}
