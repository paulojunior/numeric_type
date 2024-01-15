package services

import "github.com/paulojunior/trafilea/contract"

type Printer struct {
	handlers []contract.NumberPrinter
}

func NewPrinter() contract.NumberPrinterHandler {
	return &Printer{
		handlers: []contract.NumberPrinter{
			DefaultPrinter{},
			Multiple3Printer{},
			Multipe5Printer{},
			Multipe3and5Printer{},
		},
	}
}

func (p Printer) GetPrinter(num int) contract.NumberPrinter {
	for _, handler := range p.handlers {
		if handler.AllowPrinter(num) {
			return handler
		}
	}

	return p.handlers[0]
}
