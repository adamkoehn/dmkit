package generator

import "math/rand"

func GetAvailable10GPGemstones() []string {
	return []string{
		"Azurite (opaque mottled deep blue)",
		"Banded agate (translucent striped brown, blue, white, or red)",
		"Blue quarts (transparent pale blue)",
		"Eye agate (translucent circles of gray, white, brown, blue, or green)",
		"Hematite (opaque gray-black)",
		"Lapis lazuli (opaque light and dark blue with yellow flecks)",
		"Malachite (opaque striated light and dark green)",
		"Moss agate (translucent pink or yellow-white with mossy gray or green markings)",
		"Obsidian (opaque black)",
		"Rhodochrosite (opaaque light pink)",
		"Tiger eye (translucent brown with golden center)",
		"Turquoise (opaque light blue-green)",
	}
}

func Get10GPGemstone() string {
	return GetAvailable10GPGemstones()[rand.Intn(12)]
}

func GetAvailable50GPGemstones() []string {
	return []string{
		"Bloodstone (opaque dark gray with red flecks)",
		"Carnelian (opaque orange to red-brown)",
		"Chalcedony (opaque white)",
		"Chrysoprase (translucent green)",
		"Citrine (transparent pale yellow-brown)",
		"Jasper (opaque blue, black, or brown)",
		"Moonstone (translucent white with pale blue glow)",
		"Onyx (opaque bands of black and white, or pure black or white)",
		"Quartz (transparent white, smoky gray, or yellow)",
		"Sardonyx (opaque band of red and white)",
		"Star rose quartz (translucent rosy stone with white star-shaped center)",
		"Zircon (transparent pale blue-green)",
	}
}

func Get50GPGemstone() string {
	return GetAvailable50GPGemstones()[rand.Intn(12)]
}
