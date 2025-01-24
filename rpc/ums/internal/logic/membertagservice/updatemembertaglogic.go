package membertagservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTagLogic 更新用户标签
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
*/
type UpdateMemberTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagLogic {
	return &UpdateMemberTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberTag 更新用户标签
func (l *UpdateMemberTagLogic) UpdateMemberTag(in *umsclient.UpdateMemberTagReq) (*umsclient.UpdateMemberTagResp, error) {
	q := query.UmsMemberTag.WithContext(l.ctx)

	// 1.根据用户标签id查询用户标签是否已存在
	_, err := q.Where(query.UmsMemberTag.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据用户标签id：%d,查询用户标签失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询用户标签失败"))
	}

	item := &model.UmsMemberTag{
		ID:                in.Id,                //
		TagName:           in.TagName,           // 标签名称
		FinishOrderCount:  in.FinishOrderCount,  // 自动打标签完成订单数量
		Status:            in.Status,            // 状态：0->禁用；1->启用
		FinishOrderAmount: in.FinishOrderAmount, // 自动打标签完成订单金额
	}

	// 2.用户标签存在时,则直接更新用户标签
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户标签失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新用户标签失败")
	}

	logc.Infof(l.ctx, "更新用户标签成功,参数：%+v", in)
	return &umsclient.UpdateMemberTagResp{}, nil
}
