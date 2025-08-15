package store

import (
	"errors"
	"sync"

	"todo/model"
)

type MemoryStore struct {
	mu     sync.Mutex
	tasks  []model.Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks:  []model.Task{},
		nextID: 1,
	}
}

func (s *MemoryStore) GetAll() []model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]model.Task(nil), s.tasks...)
}

func (s *MemoryStore) GetByID(id int) (*model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, t := range s.tasks {
		if t.ID == id {
			copy := t
			return &copy, nil
		}
	}
	return nil, errors.New("not found")
}

func (s *MemoryStore) Create(t model.Task) model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	t.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, t)
	return t
}

func (s *MemoryStore) Update(id int, t model.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, item := range s.tasks {
		if item.ID == id {
			t.ID = id
			s.tasks[i] = t
			return nil
		}
	}
	return errors.New("not found")
}

func (s *MemoryStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, item := range s.tasks {
		if item.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
