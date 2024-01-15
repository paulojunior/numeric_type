package services

import "testing"

func TestDefaultPrinter_AllowPrinter(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"ShouldReturnTrueWhenNotDivisibleBy3Or5", 2, true},      // Example: number not divisible by 3 or 5
		{"ShouldReturnFalseWhenDivisibleBy3", 3, false},          // Example: number divisible by 3
		{"ShouldReturnFalseWhenDivisibleBy5", 5, false},          // Example: number divisible by 5
		{"ShouldReturnFalseWhenDivisibleByBoth3And5", 15, false}, // Example: number divisible by both 3 and 5
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := DefaultPrinter{}
			if got := b.AllowPrinter(tt.num); got != tt.want {
				t.Errorf("DefaultPrinter.AllowPrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}
