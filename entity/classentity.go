package entity

import (
	_ "encoding/json"
)

type Class struct {
	School    string `json:"school"`
	Apartment string `json:"apartment"`
	Floor     string `json:"floor"`
}
