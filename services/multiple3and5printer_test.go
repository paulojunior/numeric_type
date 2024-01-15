package services

import "testing"

func TestMultipe3and5Printer_AllowPrinter(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"ShouldReturnTrueWhenDivisibleByBoth3And5", 15, true}, // Example: number divisible by both 3 and 5
		{"ShouldReturnFalseWhenNotDivisibleBy3Or5", 2, false},  // Example: number not divisible by 3 or 5
		{"ShouldReturnFalseWhenDivisibleBy3", 3, false},        // Example: number divisible by 3
		{"ShouldReturnFalseWhenDivisibleBy5", 5, false},        // Example: number divisible by 5
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Multipe3and5Printer{}
			if got := b.AllowPrinter(tt.num); got != tt.want {
				t.Errorf("Multipe3and5Printer.AllowPrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}
