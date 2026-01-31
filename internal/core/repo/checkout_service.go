package repo

type CheckoutService struct {
	repo Repository
}

func NewCheckoutService(r Repository) *CheckoutService {
	return &CheckoutService{repo: r}
}

func (s *CheckoutService) Execute(ref string, create bool) error {
	return s.repo.Checkout(ref, create)
}
