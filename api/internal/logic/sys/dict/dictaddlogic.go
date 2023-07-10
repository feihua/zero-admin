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

type DictAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictAddLogic {
	return DictAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictAddLogic) DictAdd(req types.AddDictReq) (*types.AddDictResp, error) {
	_, err := l.svcCtx.DictService.DictAdd(l.ctx, &sysclient.DictAddReq{
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Sort:        req.Sort,
		Remarks:     req.Remarks,
		CreateBy:    l.ctx.Value("userName").(string),
		DelFlag:     req.DelFlag,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加字典信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加字典失败")
	}

	return &types.AddDictResp{
		Code:    "000000",
		Message: "添加字典成功",
	}, nil
}
