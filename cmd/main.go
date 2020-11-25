package main

import (
	"fmt"
	"strconv"

	"github.com/lwood54/go_experiment/math"
)

func main() {
	addedNum := strconv.Itoa(math.Add(5, 62))
	subNum := strconv.Itoa(math.Subtract(34, 8))
	fmt.Println("addedNum: ", addedNum)
	fmt.Println("subNum: ", subNum)
}
