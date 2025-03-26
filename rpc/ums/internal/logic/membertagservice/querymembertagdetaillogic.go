package membertagservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTagDetailLogic 查询用户标签详情
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
*/
type QueryMemberTagDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTagDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagDetailLogic {
	return &QueryMemberTagDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTagDetail 查询用户标签详情
func (l *QueryMemberTagDetailLogic) QueryMemberTagDetail(in *umsclient.QueryMemberTagDetailReq) (*umsclient.QueryMemberTagDetailResp, error) {
	item, err := query.UmsMemberTag.WithContext(l.ctx).Where(query.UmsMemberTag.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询用户标签详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户标签详情失败")
	}

	data := &umsclient.QueryMemberTagDetailResp{
		Id:                item.ID,                //
		TagName:           item.TagName,           // 标签名称
		FinishOrderCount:  item.FinishOrderCount,  // 自动打标签完成订单数量
		Status:            item.Status,            // 状态：0->禁用；1->启用
		FinishOrderAmount: item.FinishOrderAmount, // 自动打标签完成订单金额
	}

	return data, nil
}
