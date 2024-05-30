package dict_item

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

// QueryDictItemDetailLogic 查询字典数据详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:13
*/
type QueryDictItemDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemDetailLogic {
	return &QueryDictItemDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDictItemDetail 查询字典数据详情
func (l *QueryDictItemDetailLogic) QueryDictItemDetail(req *types.QueryDictItemDetailReq) (resp *types.QueryDictItemDetailResp, err error) {
	detail, err := l.svcCtx.DictItemService.QueryDictItemDetail(l.ctx, &sysclient.QueryDictItemDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典数据详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	dictItem := types.QueryDictItemDetailData{
		CreateBy:   detail.CreateBy,
		CreateTime: detail.CreateTime,
		DictLabel:  detail.DictLabel,
		DictSort:   detail.DictStatus,
		DictStatus: detail.DictStatus,
		DictType:   detail.DictType,
		DictValue:  detail.DictValue,
		Id:         detail.Id,
		IsDefault:  detail.IsDefault,
		Remark:     detail.Remark,
		UpdateBy:   detail.UpdateBy,
		UpdateTime: detail.UpdateTime,
	}

	return &types.QueryDictItemDetailResp{
		Code:    "000000",
		Message: "查询字典数据详情成功",
		Data:    dictItem,
	}, nil
}
