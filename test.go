package main

import (
	"fmt"

	"github.com/adamkoehn/dmkit/character"
)

func main() {
	character := character.CreateRandomCharacter()

	fmt.Printf("%s - %s\n", character.Race.String(), character.Class.String())
	fmt.Printf("Level: %d\n", character.Level)
	fmt.Printf("Hit Dice: %s\n", character.HitDie.String())
	fmt.Printf("Hit Points: %d\n", character.HitPoints)
	fmt.Println(character.Alignment.String())
	fmt.Println("_______________Abilities___________________")
	fmt.Printf("  Strength: %d (%d)\n", character.Abilities.Strength, character.Abilities.Strength.Modifier())
	fmt.Printf("  Dexterity: %d (%d)\n", character.Abilities.Dexterity, character.Abilities.Dexterity.Modifier())
	fmt.Printf("  Constitution: %d (%d)\n", character.Abilities.Constitution, character.Abilities.Constitution.Modifier())
	fmt.Printf("  Intelligence: %d (%d)\n", character.Abilities.Intelligence, character.Abilities.Intelligence.Modifier())
	fmt.Printf("  Wisdom: %d (%d)\n", character.Abilities.Wisdom, character.Abilities.Wisdom.Modifier())
	fmt.Printf("  Charisma: %d (%d)\n", character.Abilities.Charisma, character.Abilities.Charisma.Modifier())
	fmt.Println("_________________Saves_____________________")
	fmt.Printf("  Strength: %d\n", character.SavingThrows.Strength.Modifier)
	fmt.Printf("  Dexterity: %d\n", character.SavingThrows.Dexterity.Modifier)
	fmt.Printf("  Constitution: %d\n", character.SavingThrows.Constitution.Modifier)
	fmt.Printf("  Intelligence: %d\n", character.SavingThrows.Intelligence.Modifier)
	fmt.Printf("  Wisdom: %d\n", character.SavingThrows.Wisdom.Modifier)
	fmt.Printf("  Charisma: %d\n", character.SavingThrows.Charisma.Modifier)
}
