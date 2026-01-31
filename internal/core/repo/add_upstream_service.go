package repo

type AddUpstreamService struct {
	repo Repository
}

func NewAddUpstreamService(r Repository) *AddUpstreamService {
	return &AddUpstreamService{repo: r}
}

func (s *AddUpstreamService) Execute(upstream *GHRepo) error {
	return s.repo.AddUpstream(upstream)
}
