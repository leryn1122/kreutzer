package webhook

import (
	"net/http"
	"regexp"
)

const (
	GithubType = iota
	GitlabType
	GiteaType
)

const UserAgent = "User-Agent"

type Webhook interface {
	Verify(secret []byte) error
	GetHookID() string
	GetChannel() string
}

//goland:noinspection GoNameStartsWithPackageName
type WebhookStrategyFactory struct {
}

func NewWebhookStrategyFactory() WebhookStrategyFactory {
	return WebhookStrategyFactory{}
}

func (f *WebhookStrategyFactory) ChooseByType(scmType int) WebhookStrategy {
	switch scmType {
	case GithubType:
		return GithubWebhookStrategy{}
	case GitlabType:
		return GitlabWebhookStrategy{}
	case GiteaType:
		return GiteaWebhookStrategy{}
	default:
		panic("unreachable")
	}
}

//goland:noinspection GoNameStartsWithPackageName
type WebhookStrategy interface {
}

func ResolveWebhookFromRequest(request *http.Request) (Webhook, error) {
	userAgent := request.Header[UserAgent][0]
	isGithub, _ := regexp.MatchString("GitHub-Hookshot/[0-9]+", userAgent)
	if isGithub {
		webhook, err := NewGithubWebhook(request)
		return webhook, err
	}
	return nil, nil
}

type EventHandler = func(interface{})
