package shrink

import (
	"math/rand"
	"time"
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomize(percentage float64) bool {
	num := generator.Float64()
	return num < percentage
}
