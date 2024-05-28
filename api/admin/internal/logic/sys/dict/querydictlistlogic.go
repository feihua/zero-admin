package dict

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryDictListLogic 查询字典列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type QueryDictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryDictListLogic {
	return QueryDictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDictList 查询字典列表
func (l *QueryDictListLogic) QueryDictList(req *types.ListDictReq) (*types.ListDictResp, error) {
	var resp, err = l.svcCtx.DictService.QueryDictList(l.ctx, &sysclient.DictListReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		IsSystem:   req.IsSystem,
		DictName:   req.DictName,
		DictStatus: req.DictStatus,
		DictType:   req.DictType,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListDictData

	for _, dict := range resp.List {
		list = append(list, &types.ListDictData{
			CreateBy:   dict.CreateBy,
			CreateTime: dict.CreateTime,
			DelFlag:    dict.DelFlag,
			DictName:   dict.DictName,
			DictStatus: dict.DictStatus,
			DictType:   dict.DictType,
			Id:         dict.Id,
			IsSystem:   dict.IsSystem,
			Remark:     dict.Remark,
			UpdateBy:   dict.UpdateBy,
			UpdateTime: dict.UpdateTime,
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
