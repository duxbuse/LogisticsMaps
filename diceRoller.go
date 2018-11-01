package logisticsmaps

import "math/rand"
import "sort"

/*
ChanceOfSuccess is a function for determining the chance of beating a certain value when rolling two dice. The second paramter is to determine if the user whishes to get under or over the threshold value. It also handles the concept of minimizing or maximising. Maximising is the act of rolling an additional dice and then discarding the lowest value fo rthe total result. Minimising is the reverse.
*/
func ChanceOfSuccess(threshold int, forward bool, min int, max int) float64 {
	iterations := 1000000
	successes := 0

	for i := 0; i < iterations; i++ {
		resultsCount := 2 + min + max
		results := make([]int, resultsCount)

		//Roll all the dice
		for index := range results {
			results[index] = RollDice()
		}

		//Sort the list to make it easy to remove min and max values
		sort.Ints(results)

		//sum results
		sum := results[max] + results[len(results)-1-min]

		if forward {
			if sum >= threshold {
				successes++
			}
		} else {
			if sum <= threshold {
				successes++
			}
		}
	}
	chanceOfSuccess := float64(successes) / float64(iterations)
	return chanceOfSuccess
}

/*
RollDice is a function representing rolling a single dice with 6 sides. returning a random interger between 1 and 6.
*/
func RollDice() int {
	return rand.Intn(6) + 1
}
