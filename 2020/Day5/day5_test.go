package day5

import "testing"

var testData = map[string]int{
	"FBFBBFFRLR": 357,
	"BFFFBBFRRR": 567,
	"FFFBBBFRRR": 119,
	"BBFFBBFRLL": 820,
}

func TestSeatIDCalculation(t *testing.T) {
	for line, result := range testData {
		calculatedResult := calculateID(line)
		if calculatedResult != result {
			t.Errorf("Calculated wrong SeatID %v, want %v", calculatedResult, result)
		}
	}
}
