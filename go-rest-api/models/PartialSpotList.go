package models

type PartialSpotList struct {
    Records []PartialRecord `json:"records"`
}

type PartialRecord struct {
    Destination string          `json:"Destination"`
	PhotoURL     string          `json:"photo_de_chat"`
    Address     string          `json:"Address"`
}
