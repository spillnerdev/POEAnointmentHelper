package main

type Oil string
type Oils []Oil

const Golden Oil = "Golden Oil"
const Silver Oil = "Silver Oil"
const Opalescent Oil = "Opalescent Oil"
const Black Oil = "Black Oil"
const Crimson Oil = "Crimson Oil"
const Violet Oil = "Violet Oil"
const Indigo Oil = "Indigo Oil"
const Azure Oil = "Azure Oil"
const Teal Oil = "Teal Oil"
const Verdant Oil = "Verdant Oil"
const Amber Oil = "Amber Oil"
const Sepia Oil = "Sepia Oil"
const Clear Oil = "Clear Oil"

var IndexToOil = Oils{
	Golden,
	Silver,
	Opalescent,
	Black,
	Crimson,
	Violet,
	Indigo,
	Azure,
	Teal,
	Verdant,
	Amber,
	Sepia,
	Clear,
}

func (oil Oil) GetIndex() int {
	return OilToIndex[oil]
}

var OilToIndex = IndexToOil.toIndexMap()

func (o Oils) toIndexMap() map[Oil]int {
	var result = make(map[Oil]int)
	for i, oil := range o {
		result[oil] = i
	}

	return result
}

var ringAnointments map[string][]Oil = nil

func GetRingAnointments() map[string][]Oil {

	if ringAnointments == nil {
		ringAnointments = loadRingAnointments()
	}

	return ringAnointments
}

func loadRingAnointments() map[string][]Oil {
	result := make(map[string][]Oil)
	return *ReadJson("./ring_anoints.json", &result)
}

var amuletAnointments map[string][]Oil = nil

func GetAmuletAnointments() map[string][]Oil {

	if amuletAnointments == nil {
		amuletAnointments = loadAmuletAnointments()
	}

	return amuletAnointments
}

func loadAmuletAnointments() map[string][]Oil {
	result := make(map[string][]Oil)
	return *ReadJson("./amulet_anoints.json", &result)
}
