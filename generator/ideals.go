package generator

import (
	"math/rand"
)

type Ideals struct {
	Morality string `json:"morality"`
	Society  string `json:"society"`
}

type Generator interface {
	GenerateTrait() string
}

type GoodMoralGenerator struct{}
type EvilMoralGenerator struct{}
type LawfulSocietyGenerator struct{}
type ChaoticSocietyGenerator struct{}
type NeutralIdealGenerator struct{}

func GetAvailableGoodMoralIdeals() []string {
	return []string{
		"Beauty",
		"Charity",
		"Greater good",
		"Life",
		"Respect",
		"Self-sacrifice",
	}
}

func GetAvailableEvilMoralIdeals() []string {
	return []string{
		"Domination",
		"Greed",
		"Might",
		"Pain",
		"Retribution",
		"Slaughter",
	}
}

func GetAvailableLawfulSocietyIdeals() []string {
	return []string{
		"Community",
		"Fairness",
		"Honor",
		"Logic",
		"Responsibility",
		"Tradition",
	}
}

func GetAvailableChaoticSocietyIdeals() []string {
	return []string{
		"Change",
		"Creativity",
		"Freedom",
		"Independence",
		"No limits",
		"Whimsy",
	}
}

func GetAvailableNeutralIdeals() []string {
	return []string{
		"Balance",
		"Knowledge",
		"Live and let live",
		"Moderation",
		"Neutrality",
		"People",
		"Aspiration",
		"Discovery",
		"Glory",
		"Nation",
		"Redemption",
		"Self-knowledge",
	}
}

func (g *GoodMoralGenerator) GenerateTrait() string {
	return GetAvailableGoodMoralIdeals()[rand.Intn(6)]
}

func (e *EvilMoralGenerator) GenerateTrait() string {
	return GetAvailableEvilMoralIdeals()[rand.Intn(6)]
}

func (l *LawfulSocietyGenerator) GenerateTrait() string {
	return GetAvailableLawfulSocietyIdeals()[rand.Intn(6)]
}

func (c *ChaoticSocietyGenerator) GenerateTrait() string {
	return GetAvailableChaoticSocietyIdeals()[rand.Intn(6)]
}

func (n *NeutralIdealGenerator) GenerateTrait() string {
	return GetAvailableNeutralIdeals()[rand.Intn(12)]
}

func GenerateIdeals() Ideals {
	moralGenerators := []Generator{
		&GoodMoralGenerator{},
		&NeutralIdealGenerator{},
		&EvilMoralGenerator{},
	}

	societyGenerators := []Generator{
		&LawfulSocietyGenerator{},
		&NeutralIdealGenerator{},
		&ChaoticSocietyGenerator{},
	}

	return Ideals{
		Morality: moralGenerators[rand.Intn(3)].GenerateTrait(),
		Society:  societyGenerators[rand.Intn(3)].GenerateTrait(),
	}
}
