package models

type DevicesResponse struct {
	Data []*DevicesData `json:"data"`
}

type DevicesData struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Latitude  string   `json:"latitude"`
	Longitude string   `json:"longitude"`
}
