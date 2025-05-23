package membertagservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberTagLogic 添加用户标签
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type AddMemberTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTagLogic {
	return &AddMemberTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberTag 添加用户标签
func (l *AddMemberTagLogic) AddMemberTag(in *umsclient.AddMemberTagReq) (*umsclient.AddMemberTagResp, error) {
	q := query.UmsMemberTag

	item := &model.UmsMemberTag{
		TagName:           in.TagName,                    // 标签名称
		Description:       in.Description,                // 标签描述
		FinishOrderCount:  in.FinishOrderCount,           // 自动打标签完成订单数量
		FinishOrderAmount: float64(in.FinishOrderAmount), // 自动打标签完成订单金额
		Status:            in.Status,                     // 状态：0-禁用，1-启用
		CreateBy:          in.CreateBy,                   // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加用户标签失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加用户标签失败")
	}

	return &umsclient.AddMemberTagResp{}, nil
}
