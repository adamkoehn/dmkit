package character

import (
	"fmt"
	"math"

	"github.com/adamkoehn/dmkit/dice"
)

type CharacterName string
type PlayerName string
type Ability int
type Modifier int
type Inspiration bool
type Proficiency int
type ClassName string
type Level uint
type Background string
type Morality string
type Attitude string
type Experience uint
type SkillName string
type ArmorClass uint
type Initiative uint
type Speed uint
type HitPoints uint
type HitDiceCount uint
type PassiveWisdom uint

const (
	MORALITY_GOOD    Morality  = "Good"
	MORALITY_NEUTRAL Morality  = "Neutral"
	MORALITY_EVIL    Morality  = "Evil"
	ATTITUDE_LAWFUL  Attitude  = "Lawful"
	ATTITUDE_NEUTRAL Attitude  = "Neutral"
	ATTITUDE_CHAOTIC Attitude  = "Chaotic"
	CLASS_BARBARIAN  ClassName = "Barbarian"
	CLASS_BARD       ClassName = "Bard"
	CLASS_CLERIC     ClassName = "Cleric"
	CLASS_DRUID      ClassName = "Druid"
	CLASS_FIGHTER    ClassName = "Fighter"
	CLASS_MONK       ClassName = "Monk"
	CLASS_PALADIN    ClassName = "Paladin"
	CLASS_RANGER     ClassName = "Ranger"
	CLASS_ROGUE      ClassName = "Rogue"
	CLASS_SORCERER   ClassName = "Sorcerer"
	CLASS_WARLOCK    ClassName = "Warlock"
	CLASS_WIZARD     ClassName = "Wizard"
)

type Alignment struct {
	Morality Morality `json:"morality"`
	Attitude Attitude `json:"attitude"`
}

type ProficientAbility struct {
	Modifier   Modifier `json:"modifier"`
	Proficient bool     `json:"proficient"`
}

type Abilities struct {
	Strength     Ability `json:"strength"`
	Dexterity    Ability `json:"dexterity"`
	Constitution Ability `json:"constitution"`
	Intelligence Ability `json:"intelligence"`
	Wisdom       Ability `json:"wisdom"`
	Charisma     Ability `json:"charisma"`
}

type SavingThrows struct {
	Strength     ProficientAbility `json:"strength"`
	Dexterity    ProficientAbility `json:"dexterity"`
	Constitution ProficientAbility `json:"constitution"`
	Intelligence ProficientAbility `json:"intelligence"`
	Wisdom       ProficientAbility `json:"wisdom"`
	Charisma     ProficientAbility `json:"charisma"`
}

type Skills struct {
	Acrobatics     ProficientAbility `json:"acrobatics"`
	AnimalHandling ProficientAbility `json:"animalHandling"`
	Arcana         ProficientAbility `json:"arcana"`
	Athletics      ProficientAbility `json:"athletics"`
	Deception      ProficientAbility `json:"deception"`
	History        ProficientAbility `json:"history"`
	Insight        ProficientAbility `json:"insight"`
	Intimidation   ProficientAbility `json:"intimidation"`
	Investigation  ProficientAbility `json:"investigation"`
	Medicine       ProficientAbility `json:"medicine"`
	Nature         ProficientAbility `json:"nature"`
	Perception     ProficientAbility `json:"perception"`
	Performance    ProficientAbility `json:"performance"`
	Persuasion     ProficientAbility `json:"persuasion"`
	Religion       ProficientAbility `json:"religion"`
	SleightOfHand  ProficientAbility `json:"sleightOfHand"`
	Stealth        ProficientAbility `json:"stealth"`
	Survival       ProficientAbility `json:"survival"`
}

type DeathSaves struct {
	Success  uint
	Failures uint
}

type Character struct {
	CharacterName      CharacterName `json:"characterName"`
	PlayerName         PlayerName    `json:"clayerName"`
	Class              Class         `json:"class"`
	Level              Level         `json:"level"`
	Background         Background    `json:"background"`
	Race               Race          `json:"race"`
	Alignment          Alignment     `json:"alignment"`
	Experience         Experience    `json:"experience"`
	Inspiration        Inspiration   `json:"inspiration"`
	Proficiency        Proficiency   `json:"proficiency"`
	SavingThrows       SavingThrows  `json:"savingThrows"`
	Abilities          Abilities     `json:"abilities"`
	Skills             Skills        `json:"skills"`
	ArmorClass         ArmorClass    `json:"armorClass"`
	Initiative         Initiative    `json:"initiative"`
	Speed              Speed         `json:"speed"`
	MaxHitPoints       HitPoints     `json:"maxHitPoints"`
	HitPoints          HitPoints     `json:"hitPoints"`
	TemporaryHitPoints HitPoints     `json:"temporaryHitPoints"`
	HitDiceCount       HitDiceCount  `json:"hitDiceCount"`
	HitDie             dice.Die      `json:"hitDie"`
	DeathSaves         DeathSaves    `json:"deathSaves"`
	PassiveWisdom      PassiveWisdom `json:"passiveWisdom"`
}

func (ability Ability) Modifier() int {
	return int(math.Floor((float64(ability) - 10.0) / 2.0))
}

func (abilities Abilities) GenerateSavingThrows() SavingThrows {
	return SavingThrows{
		Strength:     ProficientAbility{Modifier: Modifier(abilities.Strength.Modifier()), Proficient: false},
		Dexterity:    ProficientAbility{Modifier: Modifier(abilities.Dexterity.Modifier()), Proficient: false},
		Constitution: ProficientAbility{Modifier: Modifier(abilities.Constitution.Modifier()), Proficient: false},
		Intelligence: ProficientAbility{Modifier: Modifier(abilities.Intelligence.Modifier()), Proficient: false},
		Wisdom:       ProficientAbility{Modifier: Modifier(abilities.Wisdom.Modifier()), Proficient: false},
		Charisma:     ProficientAbility{Modifier: Modifier(abilities.Charisma.Modifier()), Proficient: false},
	}
}

func (abilities Abilities) GenerateSkills() Skills {
	return Skills{
		Acrobatics:     ProficientAbility{Modifier: Modifier(abilities.Dexterity.Modifier()), Proficient: false},
		AnimalHandling: ProficientAbility{Modifier: Modifier(abilities.Wisdom.Modifier()), Proficient: false},
		Arcana:         ProficientAbility{Modifier: Modifier(abilities.Intelligence.Modifier()), Proficient: false},
		Athletics:      ProficientAbility{Modifier: Modifier(abilities.Strength.Modifier()), Proficient: false},
		Deception:      ProficientAbility{Modifier: Modifier(abilities.Charisma.Modifier()), Proficient: false},
		History:        ProficientAbility{Modifier: Modifier(abilities.Intelligence.Modifier()), Proficient: false},
		Insight:        ProficientAbility{Modifier: Modifier(abilities.Wisdom.Modifier()), Proficient: false},
		Intimidation:   ProficientAbility{Modifier: Modifier(abilities.Charisma.Modifier()), Proficient: false},
		Investigation:  ProficientAbility{Modifier: Modifier(abilities.Intelligence.Modifier()), Proficient: false},
		Medicine:       ProficientAbility{Modifier: Modifier(abilities.Wisdom.Modifier()), Proficient: false},
		Nature:         ProficientAbility{Modifier: Modifier(abilities.Intelligence.Modifier()), Proficient: false},
		Perception:     ProficientAbility{Modifier: Modifier(abilities.Wisdom.Modifier()), Proficient: false},
		Performance:    ProficientAbility{Modifier: Modifier(abilities.Charisma.Modifier()), Proficient: false},
		Persuasion:     ProficientAbility{Modifier: Modifier(abilities.Charisma.Modifier()), Proficient: false},
		Religion:       ProficientAbility{Modifier: Modifier(abilities.Intelligence.Modifier()), Proficient: false},
		SleightOfHand:  ProficientAbility{Modifier: Modifier(abilities.Dexterity.Modifier()), Proficient: false},
		Stealth:        ProficientAbility{Modifier: Modifier(abilities.Dexterity.Modifier()), Proficient: false},
		Survival:       ProficientAbility{Modifier: Modifier(abilities.Wisdom.Modifier()), Proficient: false},
	}
}

func (ability *ProficientAbility) AddProficiency(proficiency Proficiency) {
	ability.Proficient = true
	ability.Modifier += Modifier(proficiency)
}

func (alignment *Alignment) String() string {
	return fmt.Sprintf("%s %s", string(alignment.Attitude), string(alignment.Morality))
}
