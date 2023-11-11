package generator

import (
	"math/rand"
)

type Abilities struct {
	High string `json:"high"`
	Low  string `json:"low"`
}

func GetAbilitiesHigh() []string {
	return []string{
		"Strength - powerful, brawny, strong as an ox",
		"Dexterity - lithe, agile, graceful",
		"Constitution - hardy, hale, healthy",
		"Intelligence - studious, learned, inquisitive",
		"Wisdom - perceptive, spiritual, insightful",
		"Charisma - persuasive, forceful, born leader",
	}
}

func GetAbilitiesLow() []string {
	return []string{
		"Strength - feeble, scrawny",
		"Dexterity - clumsy, fumbling",
		"Constitution - sickly, pale",
		"Intellegence - dim-witted, slow",
		"Wisdom - oblivious, absentminded",
		"Charisma - dull, boring",
	}
}

func GenerateAbilities() Abilities {
	return Abilities{
		High: GetAbilitiesHigh()[rand.Intn(6)],
		Low:  GetAbilitiesLow()[rand.Intn(6)],
	}
}
