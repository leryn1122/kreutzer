package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/kreutzer/vo"
	"github.com/leryn1122/kreutzer/v2/lib/web"
	"github.com/leryn1122/kreutzer/v2/pkg/helm"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/cli"
	"os"
)

type HelmController struct{}

func (c HelmController) GetEnvs(ctx *gin.Context) {
	settings := cli.New()
	web.OnSuccess(ctx, settings.EnvVars())
}

func (c HelmController) ListRepo(ctx *gin.Context) {
	file, err := helm.LoadCachedRepos()
	if err != nil {
		logrus.Errorf("%+v", err)
		web.OnError(ctx, err)
		return
	}

	var repoVOs []*vo.Repo
	for _, repo_ := range file.Repositories {
		repoVO := &vo.Repo{
			Name: repo_.Name,
			URL:  repo_.URL,
		}
		repoVOs = append(repoVOs, repoVO)
	}
	web.OnSuccess(ctx, repoVOs)
}

func (c HelmController) ListCharts(ctx *gin.Context) {
	file, err := helm.LoadCachedRepoIndex(ctx.Param("repo"))
	if err != nil {
		logrus.Errorf("%+v", err)
		web.OnError(ctx, err)
		return
	}

	var chartVOs []*vo.Chart
	for _, chart := range file.Entries {
		latestChart := chart[chart.Len()-1]
		chartVO := &vo.Chart{
			Name:        latestChart.Name,
			Version:     latestChart.Version,
			AppVersion:  latestChart.AppVersion,
			Description: latestChart.Description,
		}
		chartVOs = append(chartVOs, chartVO)
	}
	web.OnSuccess(ctx, chartVOs)
}

func (c HelmController) UpdateRepo(ctx *gin.Context) {
	file, err := helm.LoadCachedRepos()
	if err != nil {
		logrus.Errorf("%+v", err)
		web.OnError(ctx, err)
		return
	}

	repoErrs := helm.UpdateRepo(file.Get(ctx.Param("repo")))
	if repoErrs != nil {
		web.OnError(ctx, fmt.Errorf("fail: %v", repoErrs))
		return
	}
	web.OnSuccess(ctx, nil)
}

func (c HelmController) ShowChartInfo(ctx *gin.Context) {
}

func (c HelmController) ListReleases(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	kubeInfo := helm.InitKubeInfo(namespace, "default", os.ExpandEnv("$HOME/.kube/config"))
	client := helm.NewHelmClient()
	releases, _ := client.ListAppsByNamespace(kubeInfo, namespace)

	var releaseVOs []*vo.Release
	for _, release := range releases {
		releaseVO := &vo.Release{
			Name:       release.Name,
			Namespace:  release.Namespace,
			Revision:   release.Version,
			Updated:    release.Info.LastDeployed,
			Status:     release.Info.Status.String(),
			Chart:      release.Chart.Name(),
			AppVersion: release.Chart.AppVersion(),
		}
		releaseVOs = append(releaseVOs, releaseVO)
	}

	web.OnSuccess(ctx, releaseVOs)
}

func (c HelmController) ShowReleaseInfo(ctx *gin.Context) {

}

func (c HelmController) GetReleaseStatus(ctx *gin.Context) {

}

func (c HelmController) ListReleaseHistories(ctx *gin.Context) {

}
