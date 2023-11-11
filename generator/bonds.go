package generator

import "math/rand"

func GetAvailableBonds() []string {
	return []string{
		"Dedicated to fulfilling a personal life goal",
		"Protective of close family members",
		"Protective of colleages or compatriots",
		"Loyal to a benefactor, patron, or employer",
		"Captivated by a romantic interest",
		"Drawn to a special place",
		"Protective of sentimental keepsake",
		"Protective of a valuable possession",
		"Out for revenge",
	}
}

func GenerateBond() string {
	return GetAvailableBonds()[rand.Intn(9)]
}
