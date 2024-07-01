package simple

type BarRepository struct{}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	BarRepository *BarRepository
}

func NewBarService(barRepo *BarRepository) *BarService {
	return &BarService{
		BarRepository: barRepo,
	}
}
