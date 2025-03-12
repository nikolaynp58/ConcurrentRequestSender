package models

type RequestResult struct {
	URL        string
	StatusCode int
	Error      error
	Number     int32
}
