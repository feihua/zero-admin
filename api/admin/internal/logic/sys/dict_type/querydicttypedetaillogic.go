package dict_type

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryDictTypeDetailLogic 查询字典类型详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:01
*/
type QueryDictTypeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictTypeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictTypeDetailLogic {
	return &QueryDictTypeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDictTypeDetail 查询字典类型详情
func (l *QueryDictTypeDetailLogic) QueryDictTypeDetail(req *types.QueryDictTypeDetailReq) (resp *types.QueryDictTypeDetailResp, err error) {
	detail, err := l.svcCtx.DictTypeService.QueryDictTypeDetail(l.ctx, &sysclient.QueryDictTypeDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典类型列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	dictType := types.QueryDictTypeDetailData{
		CreateBy:   detail.CreateBy,
		CreateTime: detail.CreateTime,
		DictName:   detail.DictType,
		DictStatus: detail.DictStatus,
		DictType:   detail.DictType,
		Id:         detail.Id,
		IsSystem:   detail.IsSystem,
		Remark:     detail.Remark,
		UpdateBy:   detail.UpdateBy,
		UpdateTime: detail.UpdateTime,
	}

	return &types.QueryDictTypeDetailResp{
		Code:    "000000",
		Message: "查询字典类型成功",
		Data:    dictType,
	}, nil
}
