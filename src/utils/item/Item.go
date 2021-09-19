package item

type Item struct {
	Id               uint8            `json:"id"`
	Name             string           `json:"name"`
	PreparationTime  uint             `json:"preparation-time"`
	Complexity       uint8            `json:"complexity"`
	CookingApparatus CookingApparatus `json:"cooking-apparatus"`
}
