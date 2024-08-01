package dice

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

type Die interface {
	String() string
	Roll() int
	Max() int
}

type D4 struct{}
type D6 struct{}
type D8 struct{}
type D10 struct{}
type D12 struct{}
type D20 struct{}
type D100 struct{}

type Dice struct {
	Die   Die
	Count uint
}

func (dice Dice) String() string {
	return fmt.Sprintf("%d%s", dice.Count, dice.Die.String())
}

func RollTimes(d Die, times int) int {
	sum := 0
	for i := 0; i < times; i++ {
		sum += d.Roll()
	}

	return sum
}

func (d D4) Roll() int {
	return rand.Intn(4) + 1
}

func (d D4) String() string {
	return "d4"
}

func (d D4) Max() int {
	return 4
}

func (d D6) Roll() int {
	return rand.Intn(6) + 1
}

func (d D6) String() string {
	return "d6"
}

func (d D6) Max() int {
	return 6
}

func (d D8) Roll() int {
	return rand.Intn(8) + 1
}

func (d D8) String() string {
	return "d8"
}

func (d D8) Max() int {
	return 8
}

func (d D10) Roll() int {
	return rand.Intn(10) + 1
}

func (d D10) String() string {
	return "d10"
}

func (d D10) Max() int {
	return 10
}

func (d D12) Roll() int {
	return rand.Intn(12) + 1
}

func (d D12) String() string {
	return "d12"
}

func (d D12) Max() int {
	return 12
}

func (d D20) Roll() int {
	return rand.Intn(20) + 1
}

func (d D20) String() string {
	return "d20"
}

func (d D20) Max() int {
	return 20
}

func (d D100) Roll() int {
	return rand.Intn(100) + 1
}

func (d D100) String() string {
	return "d100"
}

func (d D100) Max() int {
	return 100
}
