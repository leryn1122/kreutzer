package adapter

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/application/entity"
	"github.com/leryn1122/kreutzer/v2/application/repository"
	"github.com/leryn1122/kreutzer/v2/application/repository/adapter"
	service2 "github.com/leryn1122/kreutzer/v2/application/service"
	"github.com/leryn1122/kreutzer/v2/infra/actions"
	wh "github.com/leryn1122/kreutzer/v2/infra/webhook"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/yaml"
	"net/http"
)

type webhookService struct {
	WebhookRepository       repository.WebhookRepository
	WebhookStrategyFactory  wh.WebhookStrategyFactory
	SCMService              service2.SCMService
	PipelineExecutorService service2.PipelineExecutorService
}

func NewWebhookService() service2.WebhookService {
	return webhookService{
		WebhookRepository:       adapter.NewWebhookRepository(),
		WebhookStrategyFactory:  wh.NewWebhookStrategyFactory(),
		SCMService:              NewSCMService(),
		PipelineExecutorService: NewPipelineExecutorService(),
	}
}

func (s webhookService) Receive(request *http.Request) error {
	webhook, err := s.resolveWebhook(request)
	if err != nil {
		return err
	}

	err = s.VerifyWebhook(webhook)
	if err != nil {
		return err
	}

	strategy := s.WebhookStrategyFactory.ChooseByType(webhook.ChannelID)
	logrus.Infof("%+v", strategy)

	file, err := s.SCMService.GetRawFile("leryn1122/vsp", "unstable", "src/ci/github-actions/ci.yaml")
	if err != nil {
		return err
	}
	workflow := &actions.Workflow{}
	err = yaml.Unmarshal(file, workflow)

	if err != nil {
	}

	return nil
}

// VerifyWebhook verifies whether the webhook payload is signed by the authority.
func (s webhookService) VerifyWebhook(webhook *entity.Webhook) error {
	webhook, err := s.WebhookRepository.Get(context.Background(), webhook)
	if err != nil || webhook == nil {
		logrus.Errorf("Webhook [%s] is not registered.", webhook.HookID)
		return err
	}
	err = webhook.Verify()
	if err != nil {
		logrus.Errorf("Webhook [%s] is not verified by %s.", webhook.HookID, webhook.Channel)
		return err
	}
	return nil
}

func (s webhookService) resolveWebhook(request *http.Request) (*entity.Webhook, error) {
	webhook, err := wh.ResolveWebhookFromRequest(request)
	if err != nil {
		logrus.Errorf("Failed to resolve the payload of webhook: %s", request.Host)
		return nil, err
	}

	return &entity.Webhook{
		HookID:    webhook.GetHookID(),
		Channel:   "github",
		ChannelID: wh.GithubType,
		Enabled:   true,
		Payload:   webhook,
	}, nil
}
