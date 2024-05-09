package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type DictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictListLogic {
	return DictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DictList 字典列表
func (l *DictListLogic) DictList(req types.ListDictReq) (*types.ListDictResp, error) {
	resp, err := l.svcCtx.DictService.DictList(l.ctx, &sysclient.DictListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Value:    req.Value,
		Label:    req.Label,
		Type:     req.Type,
		DelFlag:  req.DelFlag,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询字典列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询字典失败")
	}

	var list []*types.ListDictData

	for _, dict := range resp.List {
		list = append(list, &types.ListDictData{
			dict.Id,
			dict.Value,
			dict.Label,
			dict.Type,
			dict.Description,
			0,
			dict.CreateBy,
			dict.CreateTime,
			dict.UpdateBy,
			dict.UpdateTime,
			dict.Remarks,
			dict.DelFlag,
		})
	}

	return &types.ListDictResp{
		Code:     "000000",
		Message:  "查询字典成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
