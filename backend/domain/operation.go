package domain

import "time"

type Operation struct {
	ID        int
	Number1   float64
	Number2   float64
	Operator  string
	Result    float64
	CreatedAt *time.Time
}
