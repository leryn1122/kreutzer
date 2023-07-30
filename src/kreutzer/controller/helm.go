package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/kreutzer/vo"
	"github.com/leryn1122/kreutzer/v2/lib/helm"
	"github.com/leryn1122/kreutzer/v2/lib/web"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/cli"
	"os"
)

type HelmController struct{}

var client = helm.NewHelmClient()

func (c HelmController) GetEnvs(ctx *gin.Context) {
	settings := cli.New()
	web.OnSuccess(ctx, settings.EnvVars())
}

func (c HelmController) ListRepos(ctx *gin.Context) {
	repos, err := client.ListRepos()
	if err != nil {
		web.OnError(ctx, err)
		return
	}

	var repoVOs []*vo.Repo
	for _, repo_ := range repos {
		repoVO := &vo.Repo{
			Name:         repo_.Name,
			URL:          repo_.URL,
			LastSyncTime: repo_.LastSyncTime,
		}
		repoVOs = append(repoVOs, repoVO)
	}
	web.OnSuccess(ctx, repoVOs)
}

func (c HelmController) AddRepo(ctx *gin.Context) {
	err := client.AddRepo()
	if err != nil {
		logrus.Errorf("%+v", err)
		web.OnError(ctx, err)
		return
	}

	web.OnSuccessMessage(ctx, "Successful update")
}

func (c HelmController) UpdateRepo(ctx *gin.Context) {
	repo := ctx.Param("repo")

	err := client.UpdateRepo(repo)
	if err != nil {
		logrus.Errorf("%+v", err)
		web.OnError(ctx, err)
		return
	}

	web.OnSuccess(ctx, repo)
}

func (c HelmController) ListCharts(ctx *gin.Context) {
	repoName := ctx.Param("repo")

	charts, err := client.ListCharts(repoName)
	if err != nil {
		logrus.Errorf("%+v", err)
		web.OnError(ctx, err)
		return
	}
	web.OnSuccess(ctx, charts)
}

func (c HelmController) ShowChartInfo(ctx *gin.Context) {
}

func (c HelmController) ListReleases(ctx *gin.Context) {
	namespace := ctx.Param("namespace")

	kubeInfo := helm.InitKubeInfo(namespace, "default", os.ExpandEnv("$HOME/.kube/config"))
	releases, err := client.ListAppsByNamespace(kubeInfo, namespace)

	if err != nil {
		logrus.Errorf("%+v", err)
		web.OnError(ctx, err)
		return
	}

	web.OnSuccess(ctx, releases)
}

func (c HelmController) ShowReleaseInfo(ctx *gin.Context) {

}

func (c HelmController) GetReleaseStatus(ctx *gin.Context) {

}

func (c HelmController) ListReleaseHistories(ctx *gin.Context) {

}
