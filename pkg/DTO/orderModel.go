package dto

type Order struct {
	ItemID   uint `json:"item_ID"`  // The item being ordered
	Quantity int  `json:"quantity"` // Order quantity

}
