package character

import (
	"math/rand"
	"sort"
	"time"

	"github.com/adamkoehn/dmkit/dice"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func CreateRandomCharacter() Character {
	race := RandomRace()
	class := RandomClass()
	abilities := RandomAbilities(race, class)
	hitDie := class.GetHitDie()
	hitPoints := hitDie.Max() + abilities.Constitution.Modifier()

	return Character{
		// CharacterName
		// PlayerName
		Class: class,
		Level: 1,
		// Background
		Race:         race,
		Alignment:    RandomAlignment(),
		Experience:   0,
		Inspiration:  false,
		Proficiency:  2,
		SavingThrows: class.GenerateSavingThrows(abilities, 2),
		Abilities:    abilities,
		Skills:       abilities.GenerateSkills(), // TODO: add proficiencies
		// Armor Class
		// Initiative
		// Speed
		MaxHitPoints:       HitPoints(hitPoints),
		HitPoints:          HitPoints(hitPoints),
		TemporaryHitPoints: HitPoints(0),
		HitDiceCount:       1,
		HitDie:             hitDie,
		DeathSaves:         DeathSaves{Success: 0, Failures: 0},
		PassiveWisdom:      PassiveWisdom(10 + abilities.Wisdom.Modifier()),
	}
}

func RandomAbility() int {
	die := dice.D6{}
	rolls := []int{
		die.Roll(),
		die.Roll(),
		die.Roll(),
		die.Roll(),
	}

	min := rolls[0]
	for i := 1; i < 4; i++ {
		if rolls[i] < min {
			min = rolls[i]
		}
	}

	for i := 0; i < 4; i++ {
		if rolls[i] == min {
			rolls[i] = 0
			break
		}
	}

	sum := 0
	for i := 0; i < 4; i++ {
		sum += rolls[i]
	}

	return sum
}

func RandomAbilities(race Race, class Class) Abilities {
	randomAbilities := []int{
		RandomAbility(),
		RandomAbility(),
		RandomAbility(),
		RandomAbility(),
		RandomAbility(),
		RandomAbility(),
	}
	increase := race.GetAbilityIncrease()

	// low to high
	sort.Ints(randomAbilities)

	preferences := []string{
		"str",
		"dex",
		"con",
		"int",
		"wis",
		"cha",
	}

	switch class.ClassName() {
	case CLASS_BARBARIAN:
		preferences = []string{"str", "con", "dex", "wis", "cha", "int"}
	case CLASS_BARD:
		preferences = []string{"cha", "dex", "int", "con", "wis", "str"}
	case CLASS_CLERIC:
		preferences = []string{"wis", "con", "str", "dex", "int", "cha"}
	case CLASS_DRUID:
		preferences = []string{"wis", "con", "dex", "str", "int", "cha"}
	case CLASS_FIGHTER:
		preferences = []string{"str", "con", "dex", "int", "cha", "wis"}
	case CLASS_MONK:
		preferences = []string{"dex", "wis", "con", "int", "cha", "str"}
	case CLASS_PALADIN:
		preferences = []string{"str", "cha", "con", "dex", "wis", "int"}
	case CLASS_RANGER:
		preferences = []string{"dex", "wis", "con", "str", "cha", "int"}
	case CLASS_ROGUE:
		preferences = []string{"dex", "int", "cha", "con", "str", "wis"}
	case CLASS_SORCERER:
		preferences = []string{"cha", "con", "wis", "dex", "int", "str"}
	case CLASS_WARLOCK:
		preferences = []string{"cha", "con", "str", "dex", "wis", "int"}
	case CLASS_WIZARD:
		preferences = []string{"int", "con", "dex", "wis", "cha", "str"}
	}

	choise := 0
	abilities := Abilities{}
	for i := 0; i < 6; i++ {
		preference := preferences[i]
		ability := randomAbilities[5-i]
		switch preference {
		case "str":
			abilities.Strength = Ability(ability) + increase.Set.Strength
			if abilities.Strength < 20 {
				abilities.Strength += Ability(increase.Choice[choise])
				choise++
			}
		case "dex":
			abilities.Dexterity = Ability(ability) + increase.Set.Dexterity
			if abilities.Dexterity < 20 {
				abilities.Dexterity += Ability(increase.Choice[choise])
				choise++
			}
		case "con":
			abilities.Constitution = Ability(ability) + increase.Set.Constitution
			if abilities.Constitution < 20 {
				abilities.Constitution += Ability(increase.Choice[choise])
				choise++
			}
		case "int":
			abilities.Intelligence = Ability(ability) + increase.Set.Intelligence
			if abilities.Intelligence < 20 {
				abilities.Intelligence += Ability(increase.Choice[choise])
				choise++
			}
		case "wis":
			abilities.Wisdom = Ability(ability) + increase.Set.Wisdom
			if abilities.Wisdom < 20 {
				abilities.Wisdom += Ability(increase.Choice[choise])
				choise++
			}
		case "cha":
			abilities.Charisma = Ability(ability) + increase.Set.Charisma
			if abilities.Charisma < 20 {
				abilities.Charisma += Ability(increase.Choice[choise])
				choise++
			}
		}
	}

	return abilities
}

func RandomRace() Race {
	baseRace := RandomBaseRace()

	switch baseRace {
	case RACE_DWARF:
		return Dwarf{SubRace: RandomDwarfSubRace()}
	case RACE_ELF:
		return Elf{SubRace: RandomElfSubRace()}
	case RACE_HALFLING:
		return Halfling{SubRace: RandomHalflingSubRace()}
	case RACE_HUMAN:
		return Human{SubRace: RandomHumanSubRace()}
	case RACE_DRAGONBORN:
		return Dragonborn{}
	case RACE_GNOME:
		return Gnome{SubRace: RandomGnomeSubRace()}
	case RACE_HALF_ELF:
		return HalfElf{}
	case RACE_HALF_ORC:
		return HalfOrc{}
	case RACE_TIEFLING:
		return Tiefling{}
	}

	panic("Cannot determine race")
}

func RandomBaseRace() BaseRace {
	return RandomFromList[BaseRace]([]BaseRace{
		RACE_DRAGONBORN,
		RACE_DWARF,
		RACE_ELF,
		RACE_GNOME,
		RACE_HALFLING,
		RACE_HALF_ELF,
		RACE_HALF_ORC,
		RACE_HUMAN,
		RACE_TIEFLING,
	})
}

func RandomDwarfSubRace() SubRace {
	return RandomFromList[SubRace]([]SubRace{
		SUB_RACE_HILL_DWARF,
		SUB_RACE_MOUNTAIN_DWARF,
	})
}

func RandomElfSubRace() SubRace {
	return RandomFromList[SubRace]([]SubRace{
		SUB_RACE_HIGH_ELF,
		SUB_RACE_WOOD_ELF,
		SUB_RACE_DROW,
	})
}

func RandomHalflingSubRace() SubRace {
	return RandomFromList[SubRace]([]SubRace{
		SUB_RACE_LIGHTFOOT,
		SUB_RACE_STOUT,
	})
}

func RandomHumanSubRace() SubRace {
	return RandomFromList[SubRace]([]SubRace{
		SUB_RACE_CHONDATHAN,
		SUB_RACE_DAMARAN,
		SUB_RACE_ILLUSKAN,
		SUB_RACE_MULAN,
		SUB_RACE_RASHEMI,
		SUB_RACE_SHOU,
		SUB_RACE_TETHYRIAN,
		SUB_RACE_TURAMI,
	})
}

func RandomGnomeSubRace() SubRace {
	return RandomFromList[SubRace]([]SubRace{
		SUB_RACE_FOREST_GNOME,
		SUB_RACE_ROCK_GNOME,
	})
}

func RandomClass() Class {
	class := RandomFromList[ClassName]([]ClassName{
		CLASS_BARBARIAN,
		CLASS_BARD,
		CLASS_CLERIC,
		CLASS_DRUID,
		CLASS_FIGHTER,
		CLASS_MONK,
		CLASS_PALADIN,
		CLASS_RANGER,
		CLASS_ROGUE,
		CLASS_SORCERER,
		CLASS_WARLOCK,
		CLASS_WIZARD,
	})

	switch class {
	case CLASS_BARBARIAN:
		return Barbarian{}
	case CLASS_BARD:
		return Bard{}
	case CLASS_CLERIC:
		return Cleric{}
	case CLASS_DRUID:
		return Druid{}
	case CLASS_FIGHTER:
		return Fighter{}
	case CLASS_MONK:
		return Monk{}
	case CLASS_PALADIN:
		return Paladin{}
	case CLASS_RANGER:
		return Ranger{}
	case CLASS_ROGUE:
		return Rogue{}
	case CLASS_SORCERER:
		return Sorcerer{}
	case CLASS_WARLOCK:
		return Warlock{}
	case CLASS_WIZARD:
		return Wizard{}
	}

	panic("Cannot determine class")
}

func RandomAlignment() Alignment {
	return Alignment{
		Attitude: RandomFromList[Attitude]([]Attitude{ATTITUDE_LAWFUL, ATTITUDE_NEUTRAL, ATTITUDE_CHAOTIC}),
		Morality: RandomFromList[Morality]([]Morality{MORALITY_GOOD, MORALITY_NEUTRAL, MORALITY_EVIL}),
	}
}

func RandomFromList[T interface{}](list []T) T {
	return list[rand.Intn(len(list))]
}
