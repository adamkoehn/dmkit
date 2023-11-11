package generator

import (
	"math/rand"
)

type NotFoundError struct {
}

func (e *NotFoundError) Error() string {
	return "Item not found"
}

func FindItem(roll int, items map[int]string) (string, error) {
	for k, v := range items {
		if roll < k {
			return v, nil
		}
	}

	return "", &NotFoundError{}
}

func Must(item string, e error) string {
	if e != nil {
		panic(e)
	}

	return item
}

func GetMagicItemsFromTable(items map[int]string) string {
	return Must(FindItem(rand.Intn(100), items))
}

func GetMagicItemFromTableA() string {
	items := map[int]string{
		50:  "Potion of healing",
		60:  "Spell scroll (cantrip)",
		70:  "Potion of climbing",
		90:  "Spell scroll (1st level)",
		94:  "Spell scroll (2nd level)",
		98:  "Potion of greater healing",
		99:  "Bag of holding",
		100: "Driftglobe",
	}

	return GetMagicItemsFromTable(items)
}

func GetMagicItemFromTableB() string {
	items := map[int]string{
		15:  "Potion of greater healing",
		22:  "Potion of fire breath",
		29:  "Potion of resistance",
		34:  "Ammunition, +1",
		39:  "Potion of animal friendship",
		44:  "Potion of hill giant strength",
		49:  "Potion of growth",
		54:  "Potion of water breathing",
		59:  "Spell scroll (2nd level)",
		64:  "Spell scroll (3rd level)",
		67:  "Bag of holding",
		70:  "Keoghtom's ointment",
		73:  "Oil of slipperiness",
		75:  "Dust of disappearance",
		77:  "Dust of dryness",
		79:  "Dust of sneezing and choking",
		81:  "Elemental gem",
		83:  "Philter of love",
		84:  "Alchemy jub",
		85:  "Cap of water breathing",
		86:  "Cloak of the manta ray",
		87:  "Driftglobe",
		88:  "Goggles of might",
		89:  "Helm of comprehending languages",
		90:  "Immovable rod",
		91:  "Lantern of revealing",
		92:  "Mariner's armor",
		93:  "Mithral armor",
		94:  "Potion of poison",
		95:  "Ring of swimming",
		96:  "Robe of useful items",
		97:  "Rope of climbing",
		98:  "Saddle of the cavalier",
		99:  "Wand of magic detection",
		100: "Wand of secrets",
	}

	return GetMagicItemsFromTable(items)
}

func GetMagicItemFromTableC() string {
	items := map[int]string{
		15:  "Potion of superior healing",
		22:  "Spell scroll (4th level)",
		27:  "Ammunition, +2",
		32:  "Potion of clairvoyance",
		37:  "Potion of diminution",
		42:  "Potion of gaseous form",
		47:  "Potion of frost giant strength",
		52:  "Portion of stone giant strength",
		57:  "Potion of heroism",
		62:  "Potion of invulnerability",
		67:  "Potion of mind reading",
		72:  "Spell scroll (5th level)",
		75:  "Elixir of health",
		78:  "Oil of etherealness",
		81:  "Potion of fire giant strength",
		84:  "Quaal's feather token",
		87:  "Scroll of protection",
		89:  "Bag of beans",
		91:  "Bead of force",
		92:  "Chime of opening",
		93:  "Decanter of endless water",
		94:  "Eyes of minute seeing",
		95:  "Folding boat",
		96:  "Heward's handy haversack",
		97:  "Horseshoes of speed",
		98:  "Necklace of fireballs",
		99:  "Periapt of health",
		100: "Sending stones",
	}

	return GetMagicItemsFromTable(items)
}

func GetMagicItemFromTableF() string {
	items := map[int]string{
		15:  "Weapon, +1",
		18:  "Shield, +1",
		21:  "Sentinel shield",
		23:  "Amulet of proof against detection and location",
		25:  "Boots of elvenkind",
		27:  "Boots of striding and springing",
		29:  "Bracers of archery",
		31:  "Brooch of shielding",
		33:  "Broom of flying",
		35:  "Cloak of elvenkind",
		37:  "Cloak of protection",
		39:  "Gauntlets of ogre power",
		41:  "Hat of disguise",
		43:  "Javelin of lightning",
		45:  "Pearl of power",
		47:  "Rod of the pact keeper, +1",
		49:  "Slippers of spider climbing",
		51:  "Staff of adder",
		53:  "Staff of the python",
		55:  "Sword of vengeance",
		57:  "Trident of fish command",
		59:  "Wand of magic missiles",
		61:  "Wand of the war mage, +1",
		63:  "Wand of web",
		65:  "Weapon of warning",
		66:  "Adamantine armer (chain mail)",
		67:  "Adamantine armor (chain shirt)",
		68:  "Adamantine armor (scale mail)",
		70:  "Bag of tricks (rust)",
		71:  "Bag of tricks (tan)",
		72:  "Boots of the winterlands",
		73:  "Circlet of blasting",
		74:  "Deck of illusions",
		75:  "Eversmoking bottle",
		76:  "Eyes of charming",
		77:  "Eyes of the eagle",
		78:  "Figurine of wondrous power (silver raven)",
		79:  "Gem of brightness",
		80:  "Gloves of missile snaring",
		81:  "Gloves of swimming and climbing",
		82:  "Gloves of tthievery",
		83:  "headband of intellect",
		84:  "Helm of telepathy",
		85:  "Instrument of the bards (Doss lute)",
		86:  "Instrument of the bards (Fochlucan bandore)",
		87:  "Instrument of the bards (Mac-Fuimidh cittern)",
		88:  "Medallion of thoughts",
		89:  "Necklace of adaptation",
		90:  "Periapt of wound closure",
		91:  "Pipes of haunting",
		92:  "Pipes of the sewers",
		93:  "Ring of jumping",
		94:  "Ring of mind shielding",
		95:  "Ring of warmth",
		96:  "Ring of water walking",
		97:  "Quiver of Ehlonna",
		98:  "Stone of good luck",
		99:  "Wind fan",
		100: "Winged boots",
	}

	return GetMagicItemsFromTable(items)
}

func GetMagicItemFromTableG() string {
	items := map[int]string{
		11:  "Weapon, +2",
		14:  "Figurine of wondrous power",
		15:  "Adamantine armor (breastplate)",
		16:  "Adamantine armor (splint)",
		17:  "Amulet of health",
		18:  "Armor of vulnerability",
		19:  "Arrow-catching shield",
		20:  "Belt of dwarvenkind",
		21:  "Belt of hill giant strength",
		22:  "Berserker axe",
		23:  "Boots of levitation",
		24:  "Boots of speed",
		25:  "Bowl of commanding water elementals",
		26:  "Bracers of defense",
		27:  "Brazier of commanding fire elementals",
		28:  "Cape of the mountebank",
		29:  "Censer of controlling air elementals",
		30:  "Armor, +1 chain mail",
		31:  "Armor of resistence (chain mail)",
		32:  "Armor, +1 chain shirt",
		33:  "Armor of resistance (chain shirt)",
		34:  "Cloak of displacement",
		35:  "Cloak of the bat",
		36:  "Cube of force",
		37:  "Daern's instant fortress",
		38:  "Dagger of venom",
		39:  "Dimensional shackles",
		40:  "Dragon slayer",
		41:  "Elven chain",
		42:  "Flame tongue",
		43:  "Gem of seeing",
		44:  "Giant slayer",
		45:  "Glamoured studded leather",
		46:  "Helm of teleportation",
		47:  "Horn of blasting",
		48:  "Horn of Valhalla (silver or brass)",
		49:  "Instrument of the bards (Canaith mandolin)",
		50:  "Instrument of the bards (Cli lyre)",
		51:  "Ioun stone (awareness)",
		52:  "Ioun stone (protection)",
		53:  "Ioun stone (reserve)",
		54:  "Ioun stone (sustenance)",
		55:  "Iron bands of Bilarro",
		56:  "Armor, +1 leather",
		57:  "Armor of resistance (leather)",
		58:  "Mace of disruption",
		59:  "Mace of smiting",
		60:  "Mace of terror",
		61:  "Mantle of spell resistance",
		62:  "Necklace of prayer beads",
		63:  "Periapt of proof against poison",
		64:  "Ring of animal influence",
		65:  "Ring of evasion",
		66:  "Ring of feather falling",
		67:  "Ring of free action",
		68:  "Ring of protection",
		69:  "Ring of resistance",
		70:  "Ring of spell storing",
		71:  "Ring of the ram",
		72:  "Ring of X-ray vision",
		73:  "Robe of eyes",
		74:  "Rod of rulership",
		75:  "Rod of the pact keeper, +2",
		76:  "Rope of entanglement",
		77:  "Armor, +1 scale mail",
		78:  "Armor of resistance (scale mail)",
		79:  "Shield, +2",
		80:  "Shield of missile attraction",
		81:  "Staff of charming",
		82:  "Staff of healing",
		83:  "Staff of swarming insects",
		84:  "Staff of the woodlands",
		85:  "Staff of withering",
		86:  "Stone of controlling earth elementals",
		87:  "Sun blade",
		88:  "Sword of life stealing",
		89:  "Sword of wounding",
		90:  "Tentacle rod",
		91:  "Vicious weapon",
		92:  "Wand of binding",
		93:  "Wand of enemy detection",
		94:  "Wand of fear",
		95:  "Wand of fireballs",
		96:  "Wand of ligtning bolts",
		97:  "Wand of paralysis",
		98:  "Wand of the war mage, +2",
		99:  "Wand of wonder",
		100: "Wings of flying",
	}

	roll := rand.Intn(100)
	item := Must(FindItem(roll, items))

	if roll > 11 && roll < 14 {
		figurineRoll := rand.Intn(8)
		figurineTypes := map[int]string{
			1: "Bronze griffon",
			2: "Ebony fly",
			3: "Golden lions",
			4: "Ivory goats",
			5: "Marble elephant",
			7: "Onyx dog",
			8: "Serpintine owl",
		}

		figurineType := Must(FindItem(figurineRoll, figurineTypes))
		item += " (" + figurineType + ")"
	}

	return item
}
