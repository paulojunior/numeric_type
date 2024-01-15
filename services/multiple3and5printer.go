package services

type Multipe3and5Printer struct{}

func (p Multipe3and5Printer) PrintNumber(num int) string {
	return "Type 3"
}

func (b Multipe3and5Printer) AllowPrinter(num int) bool {
	return num%5 == 0 && num%3 == 0
}
