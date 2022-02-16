package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictUpdateLogic {
	return DictUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictUpdateLogic) DictUpdate(req types.UpdateDictReq) (*types.UpdateDictResp, error) {
	_, err := l.svcCtx.Sys.DictUpdate(l.ctx, &sysclient.DictUpdateReq{
		Id:          req.Id,
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Sort:        int64(req.Sort),
		Remarks:     req.Remarks,
		//todo 从token里面拿
		LastUpdateBy: "admin",
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新字典信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新字典失败")
	}

	return &types.UpdateDictResp{
		Code:    "000000",
		Message: "更新字典成功",
	}, nil
}
