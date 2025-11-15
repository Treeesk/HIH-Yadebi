package models

import "time"

type App struct {
	ID        int       `json:"id"`
	Rating    float64   `json:"rating"`
	Downloads int       `json:"downloads"`
	CreatedAt time.Time `json:"created_at"`
}
