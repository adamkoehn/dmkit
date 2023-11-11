package generator

type NPC struct {
	Race             string    `json:"race"`
	Appearance       string    `json:"appearance"`
	Abilities        Abilities `json:"abilities"`
	Talent           string    `json:"talent"`
	Mannerism        string    `json:"mannerism"`
	InteractionTrait string    `json:"interaction"`
	Ideals           Ideals    `json:"ideals"`
	Bond             string    `json:"bond"`
	Flaw             string    `json:"flaw"`
}

func GenerateNpc() NPC {
	return NPC{
		Race:             GenerateRace(),
		Appearance:       GenerateAppearance(),
		Abilities:        GenerateAbilities(),
		Talent:           GenerateTalent(),
		Mannerism:        GenerateManerism(),
		InteractionTrait: GenerateInteractionTrait(),
		Ideals:           GenerateIdeals(),
		Bond:             GenerateBond(),
		Flaw:             GenerateFlaw(),
	}
}
