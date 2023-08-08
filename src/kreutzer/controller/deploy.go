package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/kreutzer/dao"
	"github.com/leryn1122/kreutzer/v2/lib/db"
	"github.com/leryn1122/kreutzer/v2/lib/web"
	"time"
)

type DeployController struct{}

func (c DeployController) CreateEnvInfo(ctx *gin.Context) {
	request := new(CreateEnvRequest)
	body, err := ctx.GetRawData()
	if err != nil {
		web.OnError(ctx, err)
	}
	if err = json.Unmarshal(body, &request); err != nil {
		web.OnError(ctx, err)
	}
	env := &dao.Environment{
		Name:        request.Name,
		Key:         request.Key,
		Description: request.Description,
		BaseDao: dao.BaseDao{
			CreationTime: time.Now(),
			ModifiedTime: time.Now(),
		},
	}
	tx := db.DBClient.Save(&env)
	if err := tx.Error; tx.Error != nil {
		web.OnError(ctx, err)
	}
	web.OnSuccess(ctx, env.Name)
}

func (c DeployController) ListEnvInfo(ctx *gin.Context) {
	var envs []dao.Environment
	tx := db.DBClient.Find(&envs)
	if err := tx.Error; tx.Error != nil {
		web.OnError(ctx, err)
	}
	web.OnSuccess(ctx, envs)
}

type CreateEnvRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
}
