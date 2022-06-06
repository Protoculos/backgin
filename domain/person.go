package domain

import "time"

type Recipe struct {
	Name         string    `json:"name,omitempty"`
	Tags         []string  `json:"tags,omitempty"`
	Ingredients  []string  `json:"ingredients,omitempty"`
	Instructions []string  `json:"instructions,omitempty"`
	PublichedAt  time.Time `json:"publiched_at,omitempty"`
}
