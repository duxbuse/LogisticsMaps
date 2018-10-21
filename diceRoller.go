package main
import "math/rand"
import "sort"

func chanceOfSuccess(threshold int, forward bool, min int, max int) float64 {
	iterations := 1000000
	successes := 0

	for i := 0; i < iterations; i++{
		resultsCount := 2 + min + max
		results := make([]int, resultsCount)

		//Roll all the dice
		for index, _ := range results{
			results[index] = rollDice()
		}

		//Sort the list to make it easy to remove min and max values
		sort.Ints(results)

		//sum results
		sum := results[max] + results[len(results) - 1 - min]

		if (forward){
			if (sum >= threshold){
				successes++
			}
		}else{
			if (sum <= threshold){
				successes++
			}
		}
	}
	chanceOfSuccess := float64(successes)/float64(iterations)
	return chanceOfSuccess
}

func rollDice() int {
	return rand.Intn(6) + 1
}