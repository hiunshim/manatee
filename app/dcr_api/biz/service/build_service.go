package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tiktok-privacy-innovation/PrivacyGo-DataCleanRoom/app/dcr_api/biz/dal/db"
	"github.com/tiktok-privacy-innovation/PrivacyGo-DataCleanRoom/pkg/utils"
)

func BuildImage(c context.Context, j db.Job, token string) error {
	if utils.RunningInsideKubernetes() {
		// use kaniko
		kanikoService := NewKanikoService(c)
		err := kanikoService.BuildImage(&j, token)
		if err != nil {
			hlog.Errorf("failed to run task %+v", err)
			return err
		}
	} else {
		hlog.Error("[BuildService] Not on kubernetes, can't build image")
	}
	return nil
}
