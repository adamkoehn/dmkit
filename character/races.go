package character

import "fmt"

type BaseRace string
type SubRace string

type AbilityIncrease struct {
	Set    Abilities
	Choice [6]int
}

type Race interface {
	String() string
	GetAbilityIncrease() AbilityIncrease
}

type Dwarf struct {
	SubRace SubRace
}

func (dwarf Dwarf) String() string {
	return string(dwarf.SubRace)
}

func (dwarf Dwarf) GetAbilityIncrease() AbilityIncrease {
	abilityIncrease := AbilityIncrease{Set: Abilities{Constitution: 2}}

	switch dwarf.SubRace {
	case SUB_RACE_HILL_DWARF:
		abilityIncrease.Set.Wisdom = 1
	case SUB_RACE_MOUNTAIN_DWARF:
		abilityIncrease.Set.Strength = 2
	}

	return abilityIncrease
}

type Elf struct {
	SubRace SubRace
}

func (elf Elf) String() string {
	return string(elf.SubRace)
}

func (elf Elf) GetAbilityIncrease() AbilityIncrease {
	abilityIncrease := AbilityIncrease{Set: Abilities{Dexterity: 2}}

	switch elf.SubRace {
	case SUB_RACE_HIGH_ELF:
		abilityIncrease.Set.Intelligence = 1
	case SUB_RACE_WOOD_ELF:
		abilityIncrease.Set.Wisdom = 1
	case SUB_RACE_DROW:
		abilityIncrease.Set.Charisma = 1
	}

	return abilityIncrease
}

type Halfling struct {
	SubRace SubRace
}

func (halfling Halfling) String() string {
	return fmt.Sprintf("%s %s", string(halfling.SubRace), RACE_HALFLING)
}

func (halfling Halfling) GetAbilityIncrease() AbilityIncrease {
	abilityIncrease := AbilityIncrease{Set: Abilities{Dexterity: 2}}

	switch halfling.SubRace {
	case SUB_RACE_LIGHTFOOT:
		abilityIncrease.Set.Charisma = 1
	case SUB_RACE_STOUT:
		abilityIncrease.Set.Constitution = 1
	}

	return abilityIncrease
}

type Human struct {
	SubRace SubRace
}

func (human Human) String() string {
	return fmt.Sprintf("%s (%s)", RACE_HUMAN, string(human.SubRace))
}

func (human Human) GetAbilityIncrease() AbilityIncrease {
	return AbilityIncrease{Set: Abilities{
		Strength:     1,
		Dexterity:    1,
		Constitution: 1,
		Intelligence: 1,
		Wisdom:       1,
		Charisma:     1,
	}}
}

type Dragonborn struct {
}

func (dragonborn Dragonborn) String() string {
	return string(RACE_DRAGONBORN)
}

func (dragonborn Dragonborn) GetAbilityIncrease() AbilityIncrease {
	return AbilityIncrease{Set: Abilities{Strength: 2, Charisma: 1}}
}

type Gnome struct {
	SubRace SubRace
}

func (gnome Gnome) String() string {
	return string(gnome.SubRace)
}

func (gnome Gnome) GetAbilityIncrease() AbilityIncrease {
	abilityIncrease := AbilityIncrease{Set: Abilities{Intelligence: 2}}

	switch gnome.SubRace {
	case SUB_RACE_FOREST_GNOME:
		abilityIncrease.Set.Dexterity = 1
	case SUB_RACE_ROCK_GNOME:
		abilityIncrease.Set.Constitution = 1
	}

	return abilityIncrease
}

type HalfElf struct {
}

func (halfElf HalfElf) String() string {
	return string(RACE_HALF_ELF)
}

func (halfElf HalfElf) GetAbilityIncrease() AbilityIncrease {
	return AbilityIncrease{Set: Abilities{Charisma: 2}, Choice: [6]int{1, 1, 0, 0, 0, 0}}
}

type HalfOrc struct {
}

func (halfOrc HalfOrc) String() string {
	return string(RACE_HALF_ORC)
}

func (halfOrc HalfOrc) GetAbilityIncrease() AbilityIncrease {
	return AbilityIncrease{Set: Abilities{Strength: 2, Constitution: 1}}
}

type Tiefling struct {
}

func (tiefling Tiefling) String() string {
	return string(RACE_TIEFLING)
}

func (tiefling Tiefling) GetAbilityIncrease() AbilityIncrease {
	return AbilityIncrease{Set: Abilities{Intelligence: 1, Charisma: 2}}
}

const (
	// Dwarf
	RACE_DWARF              BaseRace = "Dwarf"
	SUB_RACE_HILL_DWARF     SubRace  = "Hill Dwarf"
	SUB_RACE_MOUNTAIN_DWARF SubRace  = "Mountain Dwarf"

	// Elf
	RACE_ELF          BaseRace = "Elf"
	SUB_RACE_HIGH_ELF SubRace  = "High Elf"
	SUB_RACE_WOOD_ELF SubRace  = "Wood Elf"
	SUB_RACE_DROW     SubRace  = "Drow"

	// Halfling
	RACE_HALFLING      BaseRace = "Halfling"
	SUB_RACE_LIGHTFOOT SubRace  = "Lightfoot"
	SUB_RACE_STOUT     SubRace  = "Stout"

	// Human
	RACE_HUMAN          BaseRace = "Human"
	SUB_RACE_CALISHITE  SubRace  = "Calishite"
	SUB_RACE_CHONDATHAN SubRace  = "Chondathan"
	SUB_RACE_DAMARAN    SubRace  = "Damaran"
	SUB_RACE_ILLUSKAN   SubRace  = "Illuskan"
	SUB_RACE_MULAN      SubRace  = "Mulan"
	SUB_RACE_RASHEMI    SubRace  = "Rashemi"
	SUB_RACE_SHOU       SubRace  = "Shou"
	SUB_RACE_TETHYRIAN  SubRace  = "Tethyrian"
	SUB_RACE_TURAMI     SubRace  = "Turami"

	// Dragonborn
	RACE_DRAGONBORN BaseRace = "Dragonborn"

	// Gnome
	RACE_GNOME            BaseRace = "Gnome"
	SUB_RACE_FOREST_GNOME SubRace  = "Forest Gnome"
	SUB_RACE_ROCK_GNOME   SubRace  = "Rock Gnome"

	// Half-Elf
	RACE_HALF_ELF BaseRace = "Half-Elf"

	// Half-Orc
	RACE_HALF_ORC BaseRace = "Half-Orc"

	// Tiefling
	RACE_TIEFLING BaseRace = "Tiefling"
)
