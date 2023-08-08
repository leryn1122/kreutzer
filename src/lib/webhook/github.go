package webhook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/go-errors/errors"
	gogithub "github.com/google/go-github/v53/github"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

const (
	ErrMsgNoSignature      = "no signature"
	ErrMsgUnknownMethod    = "unknown method"
	ErrMsgInvalidSignature = "invalid signature"

	GithubDeliveryHeader  = "X-Github-Delivery"
	GithubHookIdHeader    = "X-GitHub-Hook-ID"
	GitHubEventHeader     = "X-GitHub-Event"
	GitHubSignatureHeader = "X-Hub-Signature"

	signaturePrefix = "sha1="
	signatureLength = 45
)

// GitHubWebhook
// Early resolved GitHub Webhook from `http.Request` discarding unnecessary fields.
type GitHubWebhook struct {
	// `X-GitHub-Hook-ID`
	HookId string
	// `X-Github-Delivery` header
	DeliveryId string
	// `X-GitHub-Event` header.
	Event string
	// `X-Hub-Signature` header.
	Signature string
	// Request body in JSON format
	Payload []byte
}

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

// SignedBy Determine whether the payload match the signature or not.
// If not, the webhook might be fabricated by a third party.
func (hook *GitHubWebhook) SignedBy(secret []byte) bool {
	if len(hook.Signature) != signatureLength || !strings.HasPrefix(hook.Signature, signaturePrefix) {
		return false
	}
	actual := make([]byte, 20)
	hex.Decode(actual, []byte(hook.Signature[5:]))
	return hmac.Equal(signBody(secret, hook.Payload), actual)
}

// NewGithubWebhook Construct a new GitHubWebhook from `http.Request`, need to be verified later.
func NewGithubWebhook(request *http.Request) (*GitHubWebhook, error) {
	if !strings.EqualFold(request.Method, http.MethodPost) {
		return nil, errors.New(ErrMsgUnknownMethod)
	}
	if signature := request.Header.Get(GitHubSignatureHeader); len(signature) == 0 {
		return nil, errors.New(ErrMsgNoSignature)
	}
	payload, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	hook := &GitHubWebhook{
		DeliveryId: request.Header.Get(GithubDeliveryHeader),
		HookId:     request.Header.Get(GithubHookIdHeader),
		Event:      request.Header.Get(GitHubEventHeader),
		Signature:  request.Header.Get(GitHubSignatureHeader),
		Payload:    payload,
	}
	return hook, nil
}

func (hook *GitHubWebhook) Verify(secret []byte) (*GitHubWebhook, error) {
	if !hook.SignedBy(secret) {
		logrus.Errorf("failed to verify signature of webhook [hook_id=%s] with signature: `%s`",
			hook.HookId, hook.Signature)
		return hook, errors.New(ErrMsgInvalidSignature)
	}
	return hook, nil
}

// UnmarshalEvent Unmarshal payload in JSON format into the event of specified type.
func (hook *GitHubWebhook) UnmarshalEvent(event string, body []byte) (interface{}, error) {
	var event_ interface{}
	switch event {
	case "commit_comment":
		event_ = gogithub.CommitCommentEvent{}
	case "create":
		event_ = gogithub.CreateEvent{}
	case "delete":
		event_ = gogithub.DeleteEvent{}
	case "deployment":
		event_ = gogithub.DeploymentEvent{}
	case "deployment_status":
		event_ = gogithub.DeploymentStatusEvent{}
	case "fork":
		event_ = gogithub.ForkEvent{}
	case "gollum":
		event_ = gogithub.GollumEvent{}
	case "issue_comment":
		event_ = gogithub.IssueCommentEvent{}
	case "issues":
		event_ = gogithub.IssuesEvent{}
	case "member":
		event_ = gogithub.MemberEvent{}
	case "membership":
		event_ = gogithub.MembershipEvent{}
	case "page_build":
		event_ = gogithub.PageBuildEvent{}
	case "public":
		event_ = gogithub.PublicEvent{}
	case "pull_request_review_comment":
		event_ = gogithub.PullRequestReviewCommentEvent{}
	case "pull_request_review":
		event_ = gogithub.PullRequestReviewEvent{}
	case "pull_request":
		event_ = gogithub.PullRequestEvent{}
	case "push":
		event_ = gogithub.PushEvent{}
	case "repository":
		event_ = gogithub.RepositoryEvent{}
	case "release":
		event_ = gogithub.ReleaseEvent{}
	case "status":
		event_ = gogithub.StatusEvent{}
	case "team_add":
		event_ = gogithub.TeamAddEvent{}
	case "watch":
		event_ = gogithub.WatchEvent{}
	default:
		return nil, errors.Errorf("unknown event type: %s", event)
	}

	err := json.Unmarshal(body, &event_)
	if err != nil {
		return nil, err
	}
	return event_, nil
}
