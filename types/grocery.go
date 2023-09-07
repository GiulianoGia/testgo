package types

type Grocery struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Done     bool   `json:"done"`
}
