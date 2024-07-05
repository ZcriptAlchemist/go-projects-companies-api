package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateCID() string {
	// Seed the random number generator to ensure different results each time
    // rand.Seed()

    // Generate a random integer
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
    randomInt := r.Intn(9000) + 1000
	fmt.Println(randomInt)
	return strconv.Itoa(randomInt)
}