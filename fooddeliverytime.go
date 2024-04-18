package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	food1 := food{}
	orderarray := []string{"burger", "chips", "nuggets"}
	if order == orderarray[0] {
		food1.preptime = 15
	} else if order == orderarray[1] {
		food1.preptime = 10
	} else if order == orderarray[2] {
		food1.preptime = 12
	} else {
		food1.preptime = 404
	}
	return food1.preptime
}
