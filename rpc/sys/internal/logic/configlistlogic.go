package logic

import (
	"context"
	"fmt"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

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

func (l *ConfigListLogic) ConfigList(in *sys.ConfigListReq) (*sys.ConfigListResp, error) {
	all, err := l.svcCtx.ConfigModel.FindAll(in.Current, in.PageSize)

	count, _ := l.svcCtx.ConfigModel.Count()
	if err != nil {
		return nil, err
	}

	var list []*sys.ConfigListData
	for _, config := range *all {
		fmt.Println(config)
		list = append(list, &sys.ConfigListData{
			Id:             config.Id,
			Value:          config.Value,
			Label:          config.Label,
			Type:           config.Type,
			Description:    config.Description,
			Sort:           int64(config.Sort),
			Remarks:        config.Remarks,
			DelFlag:        config.DelFlag,
			CreateBy:       config.CreateBy,
			CreateTime:     config.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   config.LastUpdateBy,
			LastUpdateTime: config.LastUpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &sys.ConfigListResp{
		Total: count,
		List:  list,
	}, nil

}
