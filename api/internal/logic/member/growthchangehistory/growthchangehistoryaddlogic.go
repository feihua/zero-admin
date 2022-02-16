package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryAddLogic {
	return GrowthChangeHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryAddLogic) GrowthChangeHistoryAdd(req types.AddGrowthChangeHistoryReq) (*types.AddGrowthChangeHistoryResp, error) {
	_, err := l.svcCtx.Ums.GrowthChangeHistoryAdd(l.ctx, &umsclient.GrowthChangeHistoryAddReq{
		MemberId:    req.MemberId,
		CreateTime:  req.CreateTime,
		ChangeType:  req.ChangeType,
		ChangeCount: req.ChangeCount,
		OperateMan:  req.OperateMan,
		OperateNote: req.OperateNote,
		SourceType:  req.SourceType,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员成长值变化历史记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加成长值变化历史记录失败")
	}

	return &types.AddGrowthChangeHistoryResp{
		Code:    "000000",
		Message: "添加成长值变化历史记录成功",
	}, nil
}
