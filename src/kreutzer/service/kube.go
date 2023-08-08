package service

import (
	"github.com/leryn1122/kreutzer/v2/kreutzer/dao"
	"github.com/leryn1122/kreutzer/v2/kreutzer/vo"
	"github.com/leryn1122/kreutzer/v2/lib/db"
	"github.com/leryn1122/kreutzer/v2/lib/kube"
	kubeApi "github.com/leryn1122/kreutzer/v2/lib/kube/resources"
	"time"
)

type KubeService struct{}

func (s *KubeService) fetchConfig(clusterID string) (*kube.ManagedKubeConfig, error) {
	var cluster dao.ManagedCluster
	tx := db.DBClient.Find(&cluster).
		Where("name = {}", clusterID).
		Limit(1)

	if err := tx.Error; err != nil {
		return nil, err
	}
	return kube.NewManagedKubeConfig(cluster.Name, cluster.URL, cluster.Cert, cluster.Username, cluster.Token), nil
}

func (s KubeService) ListNamespaces(clusterID string) (*[]vo.Namespace, error) {
	config, err := s.fetchConfig(clusterID)
	if err != nil {
		return nil, err
	}

	namespaces, err := kubeApi.ListNamespaces(config)
	if err != nil {
		return nil, err
	}

	var res []vo.Namespace
	for _, ns := range namespaces.Items {
		res = append(res, vo.Namespace{
			Name:   ns.Name,
			Status: string(ns.Status.Phase),
			// TODO User-friendly time format
			Age: time.Now().Sub(ns.GetCreationTimestamp().Time).String(),
		})
	}
	return &res, err
}
