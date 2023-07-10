package configservicelogic

import (
	"context"
	"fmt"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigListLogic {
	return &ConfigListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigListLogic) ConfigList(in *sysclient.ConfigListReq) (*sysclient.ConfigListResp, error) {
	all, err := l.svcCtx.ConfigModel.FindAll(l.ctx, in.Current, in.PageSize)

	count, _ := l.svcCtx.ConfigModel.Count(l.ctx)
	if err != nil {
		return nil, err
	}

	var list []*sysclient.ConfigListData
	for _, config := range *all {
		fmt.Println(config)
		list = append(list, &sysclient.ConfigListData{
			Id:             config.Id,
			Value:          config.Value,
			Label:          config.Label,
			Type:           config.Type,
			Description:    config.Description,
			Sort:           int64(config.Sort),
			Remarks:        config.Remarks.String,
			DelFlag:        config.DelFlag,
			CreateBy:       config.CreateBy.String,
			CreateTime:     config.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   config.UpdateBy.String,
			LastUpdateTime: config.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return &sysclient.ConfigListResp{
		Total: count,
		List:  list,
	}, nil

}
