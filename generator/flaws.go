package generator

import "math/rand"

func GetAvailableFlaws() []string {
	return []string{
		"Forbidden love or susceptibility to romance",
		"Enjoys decadent pleasures",
		"Arrogance",
		"Envies another creature's possessions or station",
		"Overpowering greed",
		"Prone to rage",
		"Has a powerful enemy",
		"Specific phobia",
		"Shameful or scandalous history",
		"Secret crime or misdeed",
		"Posession of forbidden lore",
		"Foolhardy bravery",
	}
}

func GenerateFlaw() string {
	return GetAvailableFlaws()[rand.Intn(12)]
}
