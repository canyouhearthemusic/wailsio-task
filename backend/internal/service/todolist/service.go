package todolist

import "wailsproject/backend/internal/domain/todo"

type Service struct {
	repo todo.Repository
}

type Configuration func(s *Service) error

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func WithTodoRepository(repo todo.Repository) Configuration {
	return func(s *Service) error {
		s.repo = repo

		return nil
	}
}
