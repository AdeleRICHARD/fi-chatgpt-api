package model

import "encoding/json"

type Classifieds struct {
	Property struct {
		Description string `json:"description"`
	} `json:"property"`
	ID string `json:"id"`
}

type Hit struct {
	Source json.RawMessage `json:"_source"`
}

type ClassifiedsList struct {
	Hits struct {
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}
