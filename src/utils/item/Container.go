package item

type Container struct {
	items []Item
}

func NewContainer(items []Item) *Container {
	return &Container{items: items}
}

func (container *Container) Get(id uint8) (*Item, bool) {
	for i := range container.items {
		if container.items[i].Id == id {
			return &container.items[i], true
		}
	}
	return nil, false
}
