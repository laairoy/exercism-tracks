package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
  units := map[string]int{
  "quarter_of_a_dozen": 3,
    "half_of_a_dozen": 6,
    "dozen": 12,
    "small_gross": 120,
    "gross": 144,
    "great_gross": 1728,
  }


  return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
  return make(map[string]int) 
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
  qtd, exists := units[unit]
  if exists == false {
    return false
  }

  bill[item] += qtd
  return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
  qtd, exists := units[unit]
  if exists == false {
    return false
  }
  
  total, exists := bill[item]
  if exists == false {
    return false
  }

  if total - qtd < 0 {
    return false
  } else if total - qtd == 0 {
    delete(bill, item)
    return true
  }

  bill[item] -= qtd
  return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
  total, exists := bill[item]
  return total, exists
}
