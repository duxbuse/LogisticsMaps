package logisticsmaps

import "math/rand"
import "sort"

/*
ChanceOfSuccess is a function for determining the chance of beating a certain value when rolling two dice. The second paramter is to determine if the user whishes to get under or over the threshold value. It also handles the concept of minimizing or maximising. Maximising is the act of rolling an additional dice and then discarding the lowest value for the total result. Minimising is the reverse. The last option is to allow rerolls in the evnt of a failed case.
*/
func ChanceOfSuccess(threshold int, over bool, reroll bool, min int, max int) float64 {
	iterations := 1000000
	successes := 0

	for i := 0; i < iterations; i++ {
		resultsCount := 2 + min + max
		sum := generateresults(resultsCount, max, min)

		if over { //get higher than the value
			if sum >= threshold {
				successes++
			} else if reroll {
				// reroll as it failed
				sum = generateresults(resultsCount, max, min)
				if sum >= threshold {
					successes++
				}
			}
		} else { //get under the value
			if sum <= threshold {
				successes++
			} else if reroll {
				// reroll as it failed
				sum = generateresults(resultsCount, max, min)
				if sum <= threshold {
					successes++
				}
			}
		}
	}
	chanceOfSuccess := float64(successes) / float64(iterations)
	return chanceOfSuccess
}

func generateresults(resultsCount int, max int, min int) int {
	results := make([]int, resultsCount)

	//Roll all the dice
	for index := range results {
		results[index] = RollDice()
	}

	//Sort the dice to make it easy to remove min and max values
	sort.Ints(results)

	//sum results of dice you dont discard
	sum := results[max] + results[len(results)-1-min]

	return sum
}

/*
RollDice is a function representing rolling a single dice with 6 sides. returning a random interger between 1 and 6.
*/
func RollDice() int {
	return rand.Intn(6) + 1
}
