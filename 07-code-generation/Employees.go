
type Employees []Employee

func (items *Employees) IndexOf(item Employee) int {
	for idx, p := range *items {
		if p == item {
			return idx
		}
	}
	return -1
}

func (items *Employees) Includes(item Employee) bool {
	return items.IndexOf(item) != -1
}

func (items *Employees) Any(criteria func(Employee) bool) bool {
	for _, item := range *items {
		if criteria(item) {
			return true
		}
	}
	return false
}

func (items *Employees) All(criteria func(Employee) bool) bool {
	for _, item := range *items {
		if !criteria(item) {
			return false
		}
	}
	return true
}

func (items *Employees) Filter(criteria func(Employee) bool) *Employees {
	result := &Employees{}
	for _, item := range *items {
		if criteria(item) {
			*result = append(*result, item)
		}
	}
	return result
}
