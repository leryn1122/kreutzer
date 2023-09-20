package service

import (
	"github.com/leryn1122/kreutzer/v2/application/vo"
	"time"
)

type AppStoreService interface {
	ListAppstores() ([]*vo.AppStore, error)
	RefreshAppStore(name string) (*time.Time, error)
	RefreshAllAppStores() (*time.Time, error)
}
