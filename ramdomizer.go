package gonet
import (
	"math/rand"
	"time"
)

var r = rand.New(time.Now().Unix())

func GetFloat64(min float64, max float64) (float64) {
	return (r.Float64() * (max - min)) + min
}

func GetInt(min int, max int) (int) {
	return r.Intn(max - min) + min
}

