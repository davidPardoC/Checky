package common

type message struct {
	Message string
}

type CustomError struct {
	StatusCode int
	Message    string
}
