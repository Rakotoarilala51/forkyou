package repo

type CloneService struct {
	repo Repository
}

func NewCloneService(r Repository) *CloneService {
	return &CloneService{repo: r}
}

func (s *CloneService) Execute(dest, ref string, create bool) error {
	if err := s.repo.Clone(dest); err != nil {
		return err
	}

	if err := s.repo.Checkout(ref, create); err != nil {
		return err
	}

	return nil
}
