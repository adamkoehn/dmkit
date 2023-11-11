package attributes

import (
	"fmt"
	"math"
)

type Name string

type Attribute uint

func (attr Attribute) Modifier() int {
	return int(math.Floor((float64(attr) - 10.0) / 2.0))
}

func (attr Attribute) String() string {
	mod := attr.Modifier()
	if mod < 0 {
		return fmt.Sprintf("%d (-%d)", uint(attr), mod)
	}

	return fmt.Sprintf("%d (+%d)", uint(attr), mod)
}

type ChallengeRating struct {
	Rating     float32 `json:"Rating"`
	Experience uint    `json:"EXP"`
}

func (cr ChallengeRating) String() string {
	if cr.Rating < 1 {
		var rating string
		if cr.Rating == 0.125 {
			rating = "1/8"
		}

		return fmt.Sprintf("%s (%d XP)", rating, cr.Experience)
	}

	return fmt.Sprintf("%d (%d XP)", uint(cr.Rating), cr.Experience)
}

type Action struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func (act Action) String() string {
	return fmt.Sprintf("%s. %s", act.Name, act.Description)
}

type Reaction struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func (react Reaction) String() string {
	return fmt.Sprintf("%s. %s", react.Name, react.Description)
}

type Trait struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func (trait Trait) String() string {
	return fmt.Sprintf("%s. %s", trait.Name, trait.Description)
}

type Speed uint

func (speed Speed) String() string {
	return fmt.Sprintf("%d ft.", uint(speed))
}

type PassivePerception uint

func (senses PassivePerception) String() string {
	return fmt.Sprintf("passive Perception %d", uint(senses))
}

type Skill uint

func (skill Skill) String() string {
	return fmt.Sprintf("+%d", uint(skill))
}

type ArmorClass uint

func (ac ArmorClass) String() string {
	return fmt.Sprintf("%d", uint(ac))
}

func (ac ArmorClass) CanHit(roll uint) bool {
	return roll >= uint(ac)
}

type HitPoints uint
type Description string
