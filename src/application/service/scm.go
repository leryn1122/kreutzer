package service

type SCMService interface {
	GetRawFile(repo, tag, path string) ([]byte, error)
}
