package character

import "github.com/adamkoehn/dmkit/dice"

type Class interface {
	ClassName() ClassName
	String() string
	GetHitDie() dice.Die
	GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows
}

type Barbarian struct{}

func (barbarian Barbarian) ClassName() ClassName {
	return CLASS_BARBARIAN
}

func (barbarian Barbarian) String() string {
	return string(barbarian.ClassName())
}

func (Barbarian Barbarian) GetHitDie() dice.Die {
	return dice.D12{}
}

func (barbarian Barbarian) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Strength.AddProficiency(proficiency)
	saves.Constitution.AddProficiency(proficiency)

	return saves
}

type Bard struct{}

func (bard Bard) ClassName() ClassName {
	return CLASS_BARD
}

func (bard Bard) String() string {
	return string(bard.ClassName())
}

func (bard Bard) GetHitDie() dice.Die {
	return dice.D8{}
}

func (bard Bard) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Dexterity.AddProficiency(proficiency)
	saves.Charisma.AddProficiency(proficiency)

	return saves
}

type Cleric struct{}

func (cleric Cleric) ClassName() ClassName {
	return CLASS_CLERIC
}

func (cleric Cleric) String() string {
	return string(cleric.ClassName())
}

func (cleric Cleric) GetHitDie() dice.Die {
	return dice.D8{}
}

func (cleric Cleric) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Wisdom.AddProficiency(proficiency)
	saves.Charisma.AddProficiency(proficiency)

	return saves
}

type Druid struct{}

func (druid Druid) ClassName() ClassName {
	return CLASS_DRUID
}

func (druid Druid) String() string {
	return string(druid.ClassName())
}

func (druid Druid) GetHitDie() dice.Die {
	return dice.D8{}
}

func (druid Druid) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Intelligence.AddProficiency(proficiency)
	saves.Wisdom.AddProficiency(proficiency)

	return saves
}

type Fighter struct{}

func (fighter Fighter) ClassName() ClassName {
	return CLASS_FIGHTER
}

func (fighter Fighter) String() string {
	return string(fighter.ClassName())
}

func (fighter Fighter) GetHitDie() dice.Die {
	return dice.D10{}
}

func (fighter Fighter) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Strength.AddProficiency(proficiency)
	saves.Constitution.AddProficiency(proficiency)

	return saves
}

type Monk struct{}

func (monk Monk) ClassName() ClassName {
	return CLASS_MONK
}

func (monk Monk) String() string {
	return string(monk.ClassName())
}

func (monk Monk) GetHitDie() dice.Die {
	return dice.D8{}
}

func (monk Monk) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Strength.AddProficiency(proficiency)
	saves.Dexterity.AddProficiency(proficiency)

	return saves
}

type Paladin struct{}

func (paladin Paladin) ClassName() ClassName {
	return CLASS_PALADIN
}

func (paladin Paladin) String() string {
	return string(paladin.ClassName())
}

func (paladin Paladin) GetHitDie() dice.Die {
	return dice.D10{}
}

func (paladin Paladin) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Wisdom.AddProficiency(proficiency)
	saves.Charisma.AddProficiency(proficiency)

	return saves
}

type Ranger struct{}

func (ranger Ranger) ClassName() ClassName {
	return CLASS_RANGER
}

func (ranger Ranger) String() string {
	return string(ranger.ClassName())
}

func (ranger Ranger) GetHitDie() dice.Die {
	return dice.D10{}
}

func (ranger Ranger) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Strength.AddProficiency(proficiency)
	saves.Dexterity.AddProficiency(proficiency)

	return saves
}

type Rogue struct{}

func (rogue Rogue) ClassName() ClassName {
	return CLASS_ROGUE
}

func (rogue Rogue) String() string {
	return string(rogue.ClassName())
}

func (rogue Rogue) GetHitDie() dice.Die {
	return dice.D8{}
}

func (rogue Rogue) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Dexterity.AddProficiency(proficiency)
	saves.Intelligence.AddProficiency(proficiency)

	return saves
}

type Sorcerer struct{}

func (sorcerer Sorcerer) ClassName() ClassName {
	return CLASS_SORCERER
}

func (sorcerer Sorcerer) String() string {
	return string(sorcerer.ClassName())
}

func (sorcerer Sorcerer) GetHitDie() dice.Die {
	return dice.D6{}
}

func (sorcerer Sorcerer) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Constitution.AddProficiency(proficiency)
	saves.Charisma.AddProficiency(proficiency)

	return saves
}

type Warlock struct{}

func (warlock Warlock) ClassName() ClassName {
	return CLASS_WARLOCK
}

func (warlock Warlock) String() string {
	return string(warlock.ClassName())
}

func (warlock Warlock) GetHitDie() dice.Die {
	return dice.D8{}
}

func (warlock Warlock) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Wisdom.AddProficiency(proficiency)
	saves.Charisma.AddProficiency(proficiency)

	return saves
}

type Wizard struct{}

func (wizard Wizard) ClassName() ClassName {
	return CLASS_WIZARD
}

func (wizard Wizard) String() string {
	return string(wizard.ClassName())
}

func (wizard Wizard) GetHitDie() dice.Die {
	return dice.D6{}
}

func (wizard Wizard) GenerateSavingThrows(abilities Abilities, proficiency Proficiency) SavingThrows {
	saves := abilities.GenerateSavingThrows()

	saves.Intelligence.AddProficiency(proficiency)
	saves.Wisdom.AddProficiency(proficiency)

	return saves
}
