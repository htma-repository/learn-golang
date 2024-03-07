package dependency_injection

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	Repository *SimpleRepository
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{
		Error: false,
	}
}

// Provider func (Function Constructor) with parameter
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	// Provider error
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		// Provider return
		return &SimpleService{
			Repository: repository,
		}, nil
	}
}
