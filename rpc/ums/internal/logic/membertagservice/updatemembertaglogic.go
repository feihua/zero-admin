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
	"gorm.io/gorm"
	"time"
)

// UpdateMemberTagLogic 更新用户标签
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
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
	tag, err := q.Where(query.UmsMemberTag.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "用户标签不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("用户标签不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询用户标签异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询用户标签异常")
	}

	now := time.Now()
	item := &model.UmsMemberTag{
		ID:                in.Id,                         // 主键ID
		TagName:           in.TagName,                    // 标签名称
		Description:       in.Description,                // 标签描述
		FinishOrderCount:  in.FinishOrderCount,           // 自动打标签完成订单数量
		FinishOrderAmount: float64(in.FinishOrderAmount), // 自动打标签完成订单金额
		Status:            in.Status,                     // 状态：0-禁用，1-启用
		CreateBy:          tag.CreateBy,                  // 创建人ID
		CreateTime:        tag.CreateTime,                // 创建时间
		UpdateBy:          &in.UpdateBy,                  // 更新人ID
		UpdateTime:        &now,                          // 更新时间

	}

	// 2.用户标签存在时,则直接更新用户标签
	err = l.svcCtx.DB.Model(&model.UmsMemberTag{}).WithContext(l.ctx).Where(query.UmsMemberTag.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新用户标签失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新用户标签失败")
	}

	return &umsclient.UpdateMemberTagResp{}, nil
}
