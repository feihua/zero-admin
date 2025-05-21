package memberlevelservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLevelLogic 更新会员等级
/*
Author: LiuFeiHua
Date: 2024/6/11 14:17
*/
type UpdateMemberLevelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLevelLogic {
	return &UpdateMemberLevelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberLevel 更新会员等级
func (l *UpdateMemberLevelLogic) UpdateMemberLevel(in *umsclient.UpdateMemberLevelReq) (*umsclient.UpdateMemberLevelResp, error) {
	q := query.UmsMemberLevel

	// 1.根据会员等级id查询会员等级是否已存在
	item, err := q.WithContext(l.ctx).Where(query.UmsMemberLevel.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员等级不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员等级不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员等级异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员等级异常")
	}

	count, err := q.WithContext(l.ctx).Where(q.ID.Neq(in.Id), q.Name.Eq(in.Name)).Count()
	if count > 0 {
		return nil, errors.New(fmt.Sprintf("会员等级名称：%s,已存在", in.Name))
	}
	count, err = q.WithContext(l.ctx).Where(q.ID.Neq(in.Id), q.Level.Eq(in.Level)).Count()
	if count > 0 {
		return nil, errors.New(fmt.Sprintf("会员等级：%d,已存在", in.Level))
	}

	now := time.Now()
	level := &model.UmsMemberLevel{
		ID:           in.Id,                    // 主键ID
		Name:         in.Name,                  // 等级名称
		Level:        in.Level,                 // 等级
		GrowthPoint:  in.GrowthPoint,           // 升级所需成长值
		DiscountRate: float64(in.DiscountRate), // 折扣率(0-100)
		FreeFreight:  in.FreeFreight,           // 是否免运费
		CommentExtra: in.CommentExtra,          // 是否可评论获取奖励
		Privileges:   in.Privileges,            // 会员特权JSON
		Remark:       in.Remark,                // 备注
		IsEnabled:    in.IsEnabled,             // 是否启用
		CreateBy:     item.CreateBy,            // 创建者
		CreateTime:   item.CreateTime,          // 创建时间
		UpdateBy:     &in.UpdateBy,             // 更新者
		UpdateTime:   &now,                     // 更新时间
		IsDeleted:    item.IsDeleted,
	}

	// 2.会员等级存在时,则直接更新会员等级
	err = l.svcCtx.DB.Model(&model.UmsMemberLevel{}).WithContext(l.ctx).Where(query.UmsMemberLevel.ID.Eq(in.Id)).Save(level).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新会员等级失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新会员等级失败")
	}

	return &umsclient.UpdateMemberLevelResp{}, nil
}
