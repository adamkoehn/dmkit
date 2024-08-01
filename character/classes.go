package character

import "github.com/adamkoehn/dmkit/dice"

type Class interface {
	ClassName() ClassName
	String() string
	GetHitDie() dice.Die
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
