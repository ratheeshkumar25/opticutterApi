package dto

type Item struct {
	ItemName    string `json:"item_name"`
	MaterialID  uint   `json:"material_id"`
	Length      uint   `json:"length"`
	Width       uint   `json:"width"`
	FixedSizeID uint   `json:"fixed_size_id"`
	IsCustom    bool   `json:"iscutom" default:"false"` // Flag indicating if size is custom
}
