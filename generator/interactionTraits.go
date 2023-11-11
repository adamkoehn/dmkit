package generator

import "math/rand"

func GetAvailableInteractionTraits() []string {
	return []string{
		"Argumentative",
		"Arrogant",
		"Blustering",
		"Rude",
		"Curious",
		"Friendly",
		"Honest",
		"Hot temptered",
		"Irritable",
		"Ponderous",
		"Quiet",
		"Suspicious",
	}
}

func GenerateInteractionTrait() string {
	return GetAvailableInteractionTraits()[rand.Intn(12)]
}
