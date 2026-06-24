package users

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) FindByUserOrEmail(value string) (*User, error) {
	return s.repository.FindByUserOrEmail(value)
}

func (s *Service) InsertUser(user *User) error {
	return s.repository.InsertUser(user)
}
