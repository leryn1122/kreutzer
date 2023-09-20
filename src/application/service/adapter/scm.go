package adapter

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/leryn1122/kreutzer/v2/application/service"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

var scmService *SCMServiceImpl

func NewSCMService() service.SCMService {
	if scmService != nil {
		return scmService
	}
	scmService := &SCMServiceImpl{}
	return scmService
}

type SCMServiceImpl struct {
}

func (s *SCMServiceImpl) GetRawFile(repo, tag, path string) ([]byte, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, tag, path)
	resp, err := http.Get(url)
	if resp.StatusCode >= http.StatusMultipleChoices {
		logrus.Errorf("Failed to get raw file at %s", url)
		return nil, errors.Errorf("Failed to get raw file at %s", url)
	}
	if err != nil {
		return nil, err
	}
	body := resp.Body
	defer body.Close()
	return io.ReadAll(body)
}
