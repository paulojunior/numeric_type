package contract

import "github.com/paulojunior/trafilea/entity"

type NumberPrinterHandler interface {
	GetPrinter(num int) NumberPrinter
}

type NumberPrinter interface {
	PrintNumber(num int) string
	AllowPrinter(num int) bool
}

type NumberService interface {
	FindAll() ([]entity.Number, error)
	FindNumber(num int) (entity.Number, error)
	Save(num int) error
}

type NumberRepository interface {
	FindAll() ([]entity.Number, error)
	FindNumber(num int) (entity.Number, error)
	Save(entity.Number) error
}
