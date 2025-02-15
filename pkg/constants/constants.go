package constants

import "time"

const (
	// DefaultPage Pagination Default
	DefaultPage  = 1
	DefaultLimit = 10

	// JWTSecretKey JWT Settings
	JWTSecretKey = "my_secret_key" // Replace with a secure key in production
	JWTDuration  = 24 * time.Hour  // Token expiration duration

	// RedisCacheDuration Redis Settings
	RedisCacheDuration = 5 * time.Minute // Default cache duration

	// TimeoutDuration Timeout Duration Middleware
	TimeoutDuration = 60 * time.Second // The Duration for server timeout

)
