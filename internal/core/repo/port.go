package repo

type Repository interface {
	Clone(dest string) error
	Checkout(ref string, create bool) error
	AddUpstream(repository *GHRepo) error
}
