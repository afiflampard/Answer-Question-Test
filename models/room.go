package models

//Room is ....
type Room struct {
	InventoriID int64     `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt uint64    `json:"purchased_at"`
	Placements  Placement `json:"placement"`
}

//Placement is ...
type Placement struct {
	RoomID uint64 `json:"room_id"`
	Name   string `json:"name"`
}
