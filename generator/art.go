package generator

import "math/rand"

func GetAvailable25GPArtObjects() []string {
	return []string{
		"Silver ewer",
		"Carved bone statuette",
		"Small gold bracelet",
		"Cloth-of-gold vestments",
		"Black velvet mask stitched with silver thread",
		"Copper chalice with silver filigree",
		"Pair of engraved bone dice",
		"Small mirror set in a painted wooden frame",
		"Embroidered silk handkerchief",
		"Gold locket with a painted portrait inside",
	}
}

func Get25GPArtwork() string {
	return GetAvailable25GPArtObjects()[rand.Intn(10)]
}
