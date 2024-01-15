package repository

import (
	"errors"
	"sync"

	"github.com/paulojunior/trafilea/contract"
	"github.com/paulojunior/trafilea/entity"
)

type NumberRepository struct {
	mu      sync.Mutex
	numbers map[int]entity.Number
}

func NewNumberRepository() contract.NumberRepository {
	return &NumberRepository{
		numbers: make(map[int]entity.Number),
	}
}

func (r *NumberRepository) Save(num entity.Number) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.numbers[num.Number]
	if exists {
		return errors.New("number already exists in our database")
	}

	r.numbers[num.Number] = num
	return nil
}

func (r *NumberRepository) FindNumber(number int) (entity.Number, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	num, exists := r.numbers[number]
	if !exists {
		return entity.Number{}, errors.New("number not found")
	}

	return num, nil
}

func (r *NumberRepository) FindAll() ([]entity.Number, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.numbers) == 0 {
		return []entity.Number{}, errors.New("numbers are empty")
	}

	numbers := make([]entity.Number, 0, len(r.numbers))
	for _, number := range r.numbers {
		numbers = append(numbers, number)
	}

	return numbers, nil
}
