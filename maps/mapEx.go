package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12,
		"small_gross": 120, "gross": 144, "great_gross": 1728}
	return units

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if !exists {
		return false
	}
	_, exists2 := bill[item]
	if exists2 {
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit]
	}
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	_, exists2 := bill[item]
	if !exists || !exists2 {
		return false
	}
	newQua := bill[item] - units[unit]
	if newQua < 0 {
		return false
	} else if newQua == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQua
	}
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]

	if !exists {
		return 0, false
	} else {
		return value, true
	}

}
