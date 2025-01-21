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
		Id:         detail.Id,         // 编号
		DictType:   detail.DictType,   // 字典类型
		DictLabel:  detail.DictLabel,  // 字典标签
		DictValue:  detail.DictValue,  // 字典键值
		DictStatus: detail.DictStatus, // 字典状态
		DictSort:   detail.DictSort,   // 排序
		Remark:     detail.Remark,     // 备注信息
		IsDefault:  detail.IsDefault,  // 是否默认  0：否  1：是
		IsDeleted:  detail.IsDeleted,  // 是否删除  0：否  1：是
		CreateBy:   detail.CreateBy,   // 创建者
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新者
		UpdateTime: detail.UpdateTime, // 更新时间
	}

	return &types.QueryDictItemDetailResp{
		Code:    "000000",
		Message: "查询字典数据详情成功",
		Data:    dictItem,
	}, nil
}
