package models

type Request struct {
	Symbol            string `json:"symbol"`
	Price             uint64 `json:"price"`
	Timestamp         uint64 `json:"timestamp"`
	RetryAttempt      int    `json:"retry_attempt"`
	RetryDuration     int    `json:"retry_duration"`
	RetryAttemptHTTP  int    `json:"retry_attempt_http"`
	RetryDurationHTTP int    `json:"retry_duration_http"`
	Timeout           int    `json:"timeout"`
}
