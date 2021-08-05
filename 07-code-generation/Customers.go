
package domain
type Customers []Customer

func (items *Customers) IndexOf(item Customer) int {
	for idx, p := range *items {
		if p == item {
			return idx
		}
	}
	return -1
}

func (items *Customers) Includes(item Customer) bool {
	return items.IndexOf(item) != -1
}

func (items *Customers) Any(criteria func(Customer) bool) bool {
	for _, item := range *items {
		if criteria(item) {
			return true
		}
	}
	return false
}

func (items *Customers) All(criteria func(Customer) bool) bool {
	for _, item := range *items {
		if !criteria(item) {
			return false
		}
	}
	return true
}

func (items *Customers) Filter(criteria func(Customer) bool) *Customers {
	result := &