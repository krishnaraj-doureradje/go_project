package user

// Service contains the business logic for Users
type Service struct {
	repo *Repository
}

// UserService defines the interface for server dependency injection
type UserService interface {
	GetAll() ([]User, error)
	Create(u User) (User, error)
}

// NewService creates a new user service
func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

// GetAll fetches all users from the repository
func (s *Service) GetAll() ([]User, error) {
	return s.repo.GetAll()
}

// Create adds a new user to the repository
func (s *Service) Create(u User) (User, error) {
	return s.repo.Create(u)
}
