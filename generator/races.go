package generator

import "math/rand"

func GetAvailableRaces() []string {
	return []string{
		"Dwarf",
		"Elf",
		"Halfling",
		"Human",
		"Dragonborn",
		"Gnome",
		"Half-Elf",
		"Half-Orc",
		"Tiefling",
	}
}

func GenerateRace() string {
	return GetAvailableRaces()[rand.Intn(9)]
}
