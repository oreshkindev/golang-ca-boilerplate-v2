package postgres

import "time"

// Option -.
type Option func(*Database)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *Database) {
		c.maxPoolSize = size
	}
}

// ConnAttempts -.
func ConnAttempts(attempts int) Option {
	return func(c *Database) {
		c.connAttempts = attempts
	}
}

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *Database) {
		c.connTimeout = timeout
	}
}
