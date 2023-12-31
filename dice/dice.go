package dice

import (
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixMicro())
}

type Die interface {
	Roll() int
	RollTimes(times int) int
}

type D4 struct{}
type D6 struct{}
type D8 struct{}
type D10 struct{}
type D12 struct{}
type D20 struct{}
type D100 struct{}

func RollTimes(d Die, times int) int {
	sum := 0
	for i := 0; i < times; i++ {
		sum += d.Roll()
	}

	return sum
}

func (d *D4) Roll() int {
	return rand.Intn(4) + 1
}

func (d *D4) RollTimes(times int) int {
	return RollTimes(d, times)
}

func (d *D6) Roll() int {
	return rand.Intn(6) + 1
}

func (d *D6) RollTimes(times int) int {
	return RollTimes(d, times)
}

func (d *D8) Roll() int {
	return rand.Intn(8) + 1
}

func (d *D8) RollTimes(times int) int {
	return RollTimes(d, times)
}

func (d *D10) Roll() int {
	return rand.Intn(10) + 1
}

func (d *D10) RollTimes(times int) int {
	return RollTimes(d, times)
}

func (d *D12) Roll() int {
	return rand.Intn(12) + 1
}

func (d *D12) RollTimes(times int) int {
	return RollTimes(d, times)
}

func (d *D20) Roll() int {
	return rand.Intn(20) + 1
}

func (d *D20) RollTimes(times int) int {
	return RollTimes(d, times)
}

func (d *D100) Roll() int {
	return rand.Intn(100) + 1
}

func (d *D100) RollTimes(times int) int {
	return RollTimes(d, times)
}
