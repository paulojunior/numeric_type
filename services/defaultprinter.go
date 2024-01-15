package services

import "fmt"

type DefaultPrinter struct{}

func (p DefaultPrinter) PrintNumber(num int) string {
	return fmt.Sprintf("%d", num)
}

func (b DefaultPrinter) AllowPrinter(num int) bool {
	return num%3 != 0 && num%5 != 0
}
