package logisticsmaps

import "testing"
import "math"

func TestRollDice(t *testing.T) {
	// Dice modeled are d6's so can only roll number 1-6
	for i := 0; i < 10000; i++ {
		value := RollDice()
		if value < 1 || value > 6 {
			t.Errorf("Rolled a dice with a value not [1-6]")
		}
	}
}

func Test7s(t *testing.T) {
	// Rolling 2 dice with no modifiers will sum to 7, 58% of the time
	expected := 58.00
	recieved := math.Floor(ChanceOfSuccess(7, true, false, 0, 0) * 100)
	if recieved != expected {
		t.Errorf("Calculated wrong Value, got: %.2f, want: %.2f.", recieved, expected)
	}
}
