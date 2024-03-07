package dependency_injection

type BarRepository struct {
}

type BarService struct {
	*BarRepository
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

func NewBarService(repository *BarRepository) *BarService {
	return &BarService{
		BarRepository: repository,
	}
}
