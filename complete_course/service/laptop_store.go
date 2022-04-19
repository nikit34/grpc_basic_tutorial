package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"

	"github.com/nikit34/grpc_basic_tutorial/complete_course/pb"
)


var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	data map[string]*pb.Laptop
	mutex sync.RWMutex
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other
	return nil
}