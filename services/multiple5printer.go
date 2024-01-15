package services

type Multipe5Printer struct{}

func (p Multipe5Printer) PrintNumber(num int) string {
	return "Type 2"
}

func (b Multipe5Printer) AllowPrinter(num int) bool {
	return num%5 == 0 && num%3 != 0
}
