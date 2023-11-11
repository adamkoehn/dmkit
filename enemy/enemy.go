package enemy

import (
	"encoding/json"
	"os"

	"github.com/adamkoehn/dmkit/attributes"
	"github.com/google/uuid"
)

type Enemy struct {
	Name            attributes.Name              `json:"Name"`
	Description     attributes.Description       `json:"Description"`
	ArmorClass      attributes.ArmorClass        `json:"AC"`
	HitPoints       attributes.HitPoints         `json:"HP"`
	Speed           attributes.Speed             `json:"Speed"`
	Strength        attributes.Attribute         `json:"STR"`
	Dexterity       attributes.Attribute         `json:"DEX"`
	Constitution    attributes.Attribute         `json:"CON"`
	Intelligence    attributes.Attribute         `json:"INT"`
	Wisdom          attributes.Attribute         `json:"WIS"`
	Charisma        attributes.Attribute         `json:"CHA"`
	Senses          attributes.PassivePerception `json:"Senses"`
	ChallengeRating attributes.ChallengeRating   `json:"CR"`
	Skills          map[string]attributes.Skill  `json:"Skills"`
	Traits          []attributes.Trait           `json:"Traits"`
	Actions         []attributes.Action          `json:"Actions"`
	Reactions       []attributes.Reaction        `json:"Reactions"`
}

type Encounter struct {
	Enemies []Enemy `json:"enemies"`
}

func (encounter Encounter) Write(name uuid.UUID) {
	file := "encounters/" + name.String() + ".json"
	contents, err := json.Marshal(encounter)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(file, contents, 0644)
	if err != nil {
		panic(err)
	}
}

func (encounter Encounter) Attack(enemy uint, hit uint, damage uint) {
	if len(encounter.Enemies) > int(enemy) {
		if encounter.Enemies[enemy].ArmorClass.CanHit(hit) {
			hp := int(encounter.Enemies[enemy].HitPoints)
			hp -= int(damage)
		}
	}
}

func GetEnemy(name string) (*Enemy, error) {
	file := "stats/" + name + ".json"
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var enemy Enemy
	err = json.Unmarshal(data, &enemy)
	if err != nil {
		return nil, err
	}

	return &enemy, nil
}

func GetEncounter(id uuid.UUID) (*Encounter, error) {
	file := "encounters/" + id.String() + ".json"
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var encounter Encounter
	err = json.Unmarshal(data, &encounter)
	if err != nil {
		return nil, err
	}

	return &encounter, nil
}
