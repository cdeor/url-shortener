package main

import "time"

type Request struct {
	URL       string        `json:"url"`
	CustomURL string        `json:"short"`
	Expiry    time.Duration `json:"expiry"`
}

type Response struct {
	URL             string        `json:"url"`
	CustomURL       string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}
