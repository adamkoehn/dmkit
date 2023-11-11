package generator

import (
	"math/rand"
)

func GetAvailableFeatures() []string {
	return []string{
		"Distinctive jewelry: earrings, necklace, circlet, bracelets",
		"Piercings",
		"Flamboyant or outlandish clothes",
		"Formal, clean clothes",
		"Ragged, dirty clothes",
		"Pronounced scar",
		"Missing teeth",
		"Missing fingers",
		"Unusual eye color (or two different colors)",
		"Tattoos",
		"Birthmark",
		"Unusual skin color",
		"Bald",
		"Braided beard or hair",
		"Unusual hair color",
		"Nervous eye twitch",
		"Distinctive nose",
		"Distinctive posture (crooked or rigid)",
		"Exceptionally beautiful",
		"Exceptionally ugly",
	}
}

func GenerateAppearance() string {
	return GetAvailableFeatures()[rand.Intn(20)]
}
