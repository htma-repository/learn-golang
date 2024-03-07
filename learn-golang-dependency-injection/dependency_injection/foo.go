package dependency_injection

type FooRepository struct {
}

type FooService struct {
	*FooRepository
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

func NewFooService(repository *FooRepository) *FooService {
	return &FooService{
		FooRepository: repository,
	}
}
