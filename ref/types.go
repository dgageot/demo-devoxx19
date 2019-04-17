package main

// Hexagon describes an hexagon.
type Hexagon struct {
	Image       string `json:"image"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Info gives information about the hexagons.
type Info struct {
	Count int `json:"count"`
}
