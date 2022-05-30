package main

import (
	"fmt"
)

func Ternary[T any](check bool, a, b T) T {
	if check {
		return a
	} else {
		return b
	}
}

func main() {

	creds := ReadCreds()
	stash, err := GetUserStash(creds)

	if err != nil {
		panic(err)
	}

	tab := stash.Tabs[0]
	rings, _ := GetAccessories(creds, tab)
	anointments := GetRingAnointments()

	threshold := Golden
	contestants := filter(rings, hasAnyOf(threshold, anointments))

	fmt.Printf("Results:\nThreshold: %s\nContestants:\n%v\n", threshold, Items(contestants).ToString())

}

type Items []Item

func (i Items) ToString() string {

	out := ""

	for _, item := range i {
		out += item.String() + "\n"
	}

	return out
}

func hasAnyOf(threshold Oil, anointments map[string][]Oil) func(item Item) bool {

	return func(item Item) bool {

		for _, enchant := range item.EnchantMods {
			oils := anointments[enchant]
			for _, oil := range oils {
				if oil.GetIndex() <= threshold.GetIndex() {
					return true
				}
			}
		}
		return false
	}
}

func hasAllOf(threshold Oil, anointments map[string][]Oil) func(item Item) bool {
	return func(item Item) bool {

		for _, enchant := range item.EnchantMods {
			oils := anointments[enchant]
			for _, oil := range oils {
				if oil.GetIndex() > threshold.GetIndex() {
					return false
				}
			}
		}
		return true
	}
}
