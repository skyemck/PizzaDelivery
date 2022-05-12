package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type DeliveryRoute struct {
	Vertical       int
	Horizontal     int
	GoatVeritcal   int
	GoatHorizontal int
	VisitedHouses  map[int]map[int]int
	HousesVisited  int
	GoatsTurn      bool
}

func Init() *DeliveryRoute {
	return &DeliveryRoute{
		Vertical:       0,
		Horizontal:     0,
		GoatVeritcal:   0,
		GoatHorizontal: 0,
		VisitedHouses:  map[int]map[int]int{},
		HousesVisited:  0,
		GoatsTurn:      false,
	}
}

func (dr *DeliveryRoute) CheckSymbol(value string) bool {
	switch {
	case value == "^":
		dr.Vertical += 1
		return true
	case value == "v":
		dr.Vertical -= 1
		return true
	case value == ">":
		dr.Horizontal += 1
		return true
	case value == "<":
		dr.Horizontal -= 1
		return true
	}
	return false
}

func (dr *DeliveryRoute) CheckSymbolGoat(value string) bool {
	switch {
	case value == "^":
		dr.GoatVeritcal += 1
		return true
	case value == "v":
		dr.GoatVeritcal -= 1
		return true
	case value == ">":
		dr.GoatHorizontal += 1
		return true
	case value == "<":
		dr.GoatHorizontal -= 1
		return true
	}
	return false
}

func (dr *DeliveryRoute) CheckMap() bool {
	if _, ok := dr.VisitedHouses[dr.Vertical]; ok {
		if _, okH := dr.VisitedHouses[dr.Vertical][dr.Horizontal]; okH {
			dr.VisitedHouses[dr.Vertical][dr.Horizontal] += 1
			return false
		} else {
			dr.VisitedHouses[dr.Vertical][dr.Horizontal] = 1
			dr.HousesVisited += 1
			return true
		}
	} else {
		dr.VisitedHouses[dr.Vertical] = make(map[int]int)
		dr.VisitedHouses[dr.Vertical][dr.Horizontal] = 1
		dr.HousesVisited += 1
		return true
	}
}

func (dr *DeliveryRoute) CheckMapGoat() bool {
	if _, ok := dr.VisitedHouses[dr.GoatVeritcal]; ok {
		if _, okH := dr.VisitedHouses[dr.GoatVeritcal][dr.GoatHorizontal]; okH {
			dr.VisitedHouses[dr.GoatVeritcal][dr.GoatHorizontal] += 1
			return false
		} else {
			dr.VisitedHouses[dr.GoatVeritcal][dr.GoatHorizontal] = 1
			dr.HousesVisited += 1
			return true
		}
	} else {
		dr.VisitedHouses[dr.GoatVeritcal] = make(map[int]int)
		dr.VisitedHouses[dr.GoatVeritcal][dr.GoatHorizontal] = 1
		dr.HousesVisited += 1
		return true
	}
}

func (dr *DeliveryRoute) Delivery(route string) int {
	routeList := strings.Split(route, "")
	count := 0
	//inital point is 0 0
	ok := dr.CheckMap()
	if !ok {
		fmt.Println("error when setting origin")
		return 0
	}
	ok = dr.CheckMapGoat()
	if ok {
		fmt.Println("error when setting origin on goat")
		return 0
	}
	count += 1
	for _, next := range routeList {
		if dr.GoatsTurn {
			ok = dr.CheckSymbolGoat(next)
			if !ok {
				fmt.Println("error when checking list of symbols")
				return 0
			}
			ok = dr.CheckMapGoat()
			if ok {
				count += 1
			}
			dr.GoatsTurn = false
		} else {
			ok = dr.CheckSymbol(next)
			if !ok {
				fmt.Println("error when checking list of symbols")
				return 0
			}
			ok = dr.CheckMap()
			if ok {
				count += 1
			}
			dr.GoatsTurn = true
		}
	}
	return count
}

func main() {
	//Test 1
	deliverer := Init()
	sample := "^v"
	count := deliverer.Delivery(sample)
	if count == deliverer.HousesVisited && count == 3 {
		fmt.Println("success Test 1")
	}
	//Test 2
	deliverer = Init()
	sample = "^>v<"
	count = deliverer.Delivery(sample)
	if count == deliverer.HousesVisited && count == 3 {
		fmt.Println("success Test 2")
	}
	//Test 3
	deliverer = Init()
	sample = "^v^v^v^v^v"
	count = deliverer.Delivery(sample)
	if count == deliverer.HousesVisited && count == 11 {
		fmt.Println("success Test 3")
	}
	//Test 4
	deliverer = Init()
	buf, err := ioutil.ReadFile("PizzaDeliveryInput.txt")
    if err != nil {
        fmt.Print(err)
    }
	sample = string(buf)
	count = deliverer.Delivery(sample)
	if count == deliverer.HousesVisited {
		fmt.Println("success actual Test count: %d", count)
	}
}
