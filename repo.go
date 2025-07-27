package forkyou

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

type GHRepo struct{
	RepoDir string
	owner string
	project string
	repo *git.Repository
}

func NewGHRepo(repository string) (*GHRepo, error){
	values := strings.Split(repository, "/")
	if len(values)!=2{
		return nil, fmt.Errorf("Repository must be in format owner/project")
	}
	return &GHRepo{
		owner: values[0],
		project: values[1],
	},nil
}
func (g *GHRepo) RepositoryURL() string{
	return fmt.Sprintf("https://github.com/%s/%s", g.owner, g.project)
}
func (g *GHRepo) Clone(dest string) error{
	fullPath := filepath.Join(dest, fmt.Sprintf("%s-%s", g.owner, g.project))
	repo, err := git.PlainClone(fullPath, &git.CloneOptions{
		URL: g.RepositoryURL(),
	})
	if err != nil {
		return err
	}
	g.repo := repo
	g.RepoDir := fullPath
	return nil
}