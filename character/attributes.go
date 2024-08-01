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

type Skill struct {
	Skill    SkillName         `json:"skill"`
	Modifier ProficientAbility `json:"modifier"`
}

type Skills struct {
	Acrobatics     Skill `json:"acrobatics"`
	AnimalHandling Skill `json:"animalHandling"`
	Arcana         Skill `json:"arcana"`
	Athletics      Skill `json:"athletics"`
	Deception      Skill `json:"deception"`
	History        Skill `json:"history"`
	Insight        Skill `json:"insight"`
	Intimidation   Skill `json:"intimidation"`
	Investigation  Skill `json:"investigation"`
	Medicine       Skill `json:"medicine"`
	Nature         Skill `json:"nature"`
	Perception     Skill `json:"perception"`
	Performance    Skill `json:"performance"`
	Persuasion     Skill `json:"persuasion"`
	Religion       Skill `json:"religion"`
	SleightOfHand  Skill `json:"sleightOfHand"`
	Stealth        Skill `json:"stealth"`
	Survival       Skill `json:"survival"`
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

func (alignment *Alignment) String() string {
	return fmt.Sprintf("%s %s", string(alignment.Attitude), string(alignment.Morality))
}
