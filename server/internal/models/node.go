package models

import "time"

type Node struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Identifier   string    `json:"identifier"`
	IsActive     bool      `json:"is_active"`
	IsOnline     bool      `json:"is_online"`
	LastOnlineAt time.Time `json:"last_online_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type NodeAPIKey struct {
	ID          int       `json:"id"`
	ApiKey      string    `json:"api_key"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NodeApiKeyValue struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type NodeWithAPIKeys struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Identifier   string       `json:"identifier"`
	IsActive     bool         `json:"is_active"`
	IsOnline     bool         `json:"is_online"`
	LastOnlineAt time.Time    `json:"last_online_at"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	APIKeys      []NodeAPIKey `json:"api_keys"`
}
