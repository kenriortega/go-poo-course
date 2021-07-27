package invoiceitem

// Item contains of information of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

// New return a new Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

type Items []Item

func NewItems(items ...Item) Items {
	var is Items
	for _, it := range items {
		is = append(is, it)
	}
	return is
}
func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}
	return total
}
