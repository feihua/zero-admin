package seckillsessionservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateSeckillSessionLogic 更新秒杀场次
/*
Author: LiuFeiHua
Date: 2025/06/11 10:29:58
*/
type UpdateSeckillSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSeckillSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillSessionLogic {
	return &UpdateSeckillSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSeckillSession 更新秒杀场次
func (l *UpdateSeckillSessionLogic) UpdateSeckillSession(in *smsclient.UpdateSeckillSessionReq) (*smsclient.UpdateSeckillSessionResp, error) {
	session := query.SmsSeckillSession
	q := session.WithContext(l.ctx)

	// 1.根据秒杀场次id查询秒杀场次是否已存在
	detail, err := q.Where(session.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "秒杀场次不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("秒杀场次不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询秒杀场次异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询秒杀场次异常")
	}

	count, err := session.WithContext(l.ctx).Where(session.Name.Eq(in.Name), session.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀场次失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新秒杀场次失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("秒杀场次：%s,已存在", in.Name))
	}

	now := time.Now()
	item := &model.SmsSeckillSession{
		ID:         in.Id,             // 秒杀场次ID
		Name:       in.Name,           // 场次名称
		StartTime:  in.StartTime,      // 开始时间
		EndTime:    in.EndTime,        // 结束时间
		Sort:       in.Sort,           // 排序
		CreateBy:   detail.CreateBy,   // 创建人ID
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   &in.UpdateBy,      // 更新人ID
		UpdateTime: &now,              // 更新时间
	}

	// 2.秒杀场次存在时,则直接更新秒杀场次
	err = l.svcCtx.DB.Model(&model.SmsSeckillSession{}).WithContext(l.ctx).Where(session.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀场次失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新秒杀场次失败")
	}

	return &smsclient.UpdateSeckillSessionResp{}, nil
}
