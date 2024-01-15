package services

import "testing"

func TestMultiple3Printer_AllowPrinter(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"ShouldReturnTrueWhenDivisibleBy3", 3, true},            // Example: number divisible by 3
		{"ShouldReturnFalseWhenNotDivisibleBy3Or5", 2, false},    // Example: number not divisible by 3 or 5
		{"ShouldReturnFalseWhenDivisibleBy5", 5, false},          // Example: number divisible by 5
		{"ShouldReturnFalseWhenDivisibleByBoth3And5", 15, false}, // Example: number divisible by both 3 and 5
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Multiple3Printer{}
			if got := b.AllowPrinter(tt.num); got != tt.want {
				t.Errorf("Multiple3Printer.AllowPrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}
