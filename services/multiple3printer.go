package services

type Multiple3Printer struct{}

func (p Multiple3Printer) PrintNumber(num int) string {
	return "Type 1"
}

func (b Multiple3Printer) AllowPrinter(num int) bool {
	return num%3 == 0 && num%5 != 0
}
