package services

import (
	"github.com/paulojunior/trafilea/contract"
	"github.com/paulojunior/trafilea/entity"
)

type NumberService struct {
	numberRespository contract.NumberRepository
	printerHandler    contract.NumberPrinterHandler
}

func NewNumberService(numberRepository contract.NumberRepository, printerHandler contract.NumberPrinterHandler) contract.NumberService {
	return &NumberService{
		numberRespository: numberRepository,
		printerHandler:    printerHandler,
	}
}

func (s *NumberService) FindAll() ([]entity.Number, error) {
	numbers, err := s.numberRespository.FindAll()
	if err != nil {
		return []entity.Number{}, err
	}

	return numbers, nil
}

func (s *NumberService) FindNumber(num int) (entity.Number, error) {
	numbers, err := s.numberRespository.FindNumber(num)
	if err != nil {
		return entity.Number{}, err
	}

	return numbers, nil
}

func (s *NumberService) Save(num int) error {
	printer := s.printerHandler.GetPrinter(num)
	numberType := printer.PrintNumber(num)

	number := entity.Number{
		Number: num,
		Type:   numberType,
	}

	err := s.numberRespository.Save(number)
	if err != nil {
		return err
	}

	return nil
}
