package generator

import "github.com/adamkoehn/dmkit/dice"

type Coins struct {
	CP int `json:"cp"`
	SP int `json:"sp"`
	EP int `json:"ep"`
	GP int `json:"gp"`
	PP int `json:"pp"`
}

type TreasureHoard struct {
	Coins    Coins    `json:"coins"`
	Gems     []string `json:"gems"`
	Artworks []string `json:"art"`
	Items    []string `json:"items"`
}

func AddTableAItems(treasure *TreasureHoard, num int) {
	for i := 0; i < num; i++ {
		treasure.Items = append(treasure.Items, GetMagicItemFromTableA())
	}
}

func AddTableBItems(treasure *TreasureHoard, num int) {
	for i := 0; i < num; i++ {
		treasure.Items = append(treasure.Items, GetMagicItemFromTableB())
	}
}

func AddTableCItems(treasure *TreasureHoard, num int) {
	for i := 0; i < num; i++ {
		treasure.Items = append(treasure.Items, GetMagicItemFromTableC())
	}
}

func AddTableFItems(treasure *TreasureHoard, num int) {
	for i := 0; i < num; i++ {
		treasure.Items = append(treasure.Items, GetMagicItemFromTableF())
	}
}

func AddTableGItems(treasure *TreasureHoard, num int) {
	for i := 0; i < num; i++ {
		treasure.Items = append(treasure.Items, GetMagicItemFromTableG())
	}
}

func Add25GPArtwork(treasure *TreasureHoard, num int) {
	for i := 0; i < num; i++ {
		treasure.Artworks = append(treasure.Artworks, Get25GPArtwork())
	}
}

func Add10GPGems(treasure *TreasureHoard, num int) {
	for i := 0; i < num; i++ {
		treasure.Gems = append(treasure.Gems, Get10GPGemstone())
	}
}

func Add50GPGems(treasure *TreasureHoard, num int) {
	for i := 0; i < num; i++ {
		treasure.Gems = append(treasure.Gems, Get50GPGemstone())
	}
}

func GetTreasureHoard0To4() *TreasureHoard {
	d4 := dice.D4{}
	d6 := dice.D6{}
	d100 := dice.D100{}

	treasure := TreasureHoard{
		Coins: Coins{
			CP: d6.RollTimes(6) * 100,
			SP: d6.RollTimes(3) * 100,
			EP: 0,
			GP: d6.RollTimes(2) * 10,
			PP: 0,
		},
	}

	roll := d100.Roll()

	if roll >= 7 && roll <= 16 {
		Add10GPGems(&treasure, d6.RollTimes(2))
	}

	if roll >= 17 && roll <= 26 {
		Add25GPArtwork(&treasure, d4.RollTimes(2))
	}

	if roll >= 27 && roll <= 36 {
		Add50GPGems(&treasure, d6.RollTimes(2))
	}

	if roll >= 37 && roll <= 44 {
		Add10GPGems(&treasure, d6.RollTimes(2))
		AddTableAItems(&treasure, d6.Roll())
	}

	if roll >= 45 && roll <= 52 {
		Add25GPArtwork(&treasure, d4.RollTimes(2))
		AddTableAItems(&treasure, d6.Roll())
	}

	if roll >= 53 && roll <= 60 {
		Add50GPGems(&treasure, d6.RollTimes(2))
		AddTableAItems(&treasure, d6.Roll())
	}

	if roll >= 61 && roll <= 65 {
		Add10GPGems(&treasure, d6.RollTimes(2))
		AddTableBItems(&treasure, d4.Roll())
	}

	if roll >= 66 && roll <= 70 {
		Add25GPArtwork(&treasure, d4.RollTimes(2))
		AddTableBItems(&treasure, d4.Roll())
	}

	if roll >= 71 && roll <= 75 {
		Add50GPGems(&treasure, d6.RollTimes(2))
		AddTableBItems(&treasure, d4.Roll())
	}

	if roll >= 76 && roll <= 78 {
		Add10GPGems(&treasure, d6.RollTimes(2))
		AddTableCItems(&treasure, d4.Roll())
	}

	if roll >= 79 && roll <= 80 {
		Add25GPArtwork(&treasure, d4.RollTimes(2))
		AddTableCItems(&treasure, d4.Roll())
	}

	if roll >= 81 && roll <= 85 {
		Add50GPGems(&treasure, d6.RollTimes(2))
		AddTableCItems(&treasure, d4.Roll())
	}

	if roll >= 86 && roll <= 92 {
		Add25GPArtwork(&treasure, d4.RollTimes(2))
		AddTableFItems(&treasure, d4.Roll())
	}

	if roll >= 93 && roll <= 97 {
		Add50GPGems(&treasure, d6.RollTimes(2))
		AddTableFItems(&treasure, d4.Roll())
	}

	if roll >= 98 || roll <= 99 {
		Add25GPArtwork(&treasure, d4.RollTimes(2))
		AddTableGItems(&treasure, 1)
	}

	if roll == 100 {
		Add50GPGems(&treasure, d6.RollTimes(2))
		AddTableGItems(&treasure, 1)
	}

	return &treasure
}
