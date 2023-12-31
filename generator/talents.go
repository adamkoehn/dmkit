package generator

import "math/rand"

func GetAvailableTalents() []string {
	return []string{
		"Plays a musical instrument",
		"Speaks several languages fluently",
		"Unbelievably lucky",
		"Perfect memory",
		"Great with animals",
		"Great with children",
		"Great at solving puzzles",
		"great at one game",
		"Great at impersonations",
		"Draws beautifully",
		"Paints beautifully",
		"Sings beautifully",
		"Drinks everyone under the table",
		"Exert carpenter",
		"Expert cook",
		"Expert dart thrower and rock skipper",
		"Expert juggler",
		"Skilled actor and master of disguise",
		"Skilled dancer",
		"Knows thieves' cant",
	}
}

func GenerateTalent() string {
	return GetAvailableTalents()[rand.Intn(20)]
}
