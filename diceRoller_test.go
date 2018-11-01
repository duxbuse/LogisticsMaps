package LogisticsMaps

import "testing"
import "math"

func TestRollDice(t *testing.T) {
	for i := 0; i < 10000; i++ {
		value := RollDice()
		if value < 1 || value > 6 {
			t.Errorf("Rolled a dice with a value not [1-6]")
		}
	}
}

func Test7s(t *testing.T) {
	expected := 58.00
	recieved := math.Floor(ChanceOfSuccess(7, true, 0, 0) * 100)
	if recieved != expected {
		t.Errorf("Calculated wrong Value, got: %.2f, want: %.2f.", recieved, expected)
	}
}
