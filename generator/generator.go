package generator

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}
