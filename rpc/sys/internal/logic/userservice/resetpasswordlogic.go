package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReSetPasswordLogic 重置用户密码
/*
Author: LiuFeiHua
Date: 2023/12/18 14:12
*/
type ReSetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReSetPassword 重置用户密码
func (l *ReSetPasswordLogic) ReSetPassword(in *sysclient.ReSetPasswordReq) (*sysclient.ReSetPasswordResp, error) {

	// 不能重置重置管理员的密码
	if common.IsAdmin(l.ctx, in.Id, l.svcCtx.DB) {
		logc.Errorf(l.ctx, "不能重置重置管理员的密码,参数:%+v", in)
		return nil, errors.New("不能重置重置管理员的密码")
	}

	_, err := query.SysUser.WithContext(l.ctx).Updates(&model.SysUser{
		ID:       in.Id,
		Password: "123456",
		UpdateBy: in.UpdateBy,
	})

	if err != nil {
		logc.Errorf(l.ctx, "重置用户密码失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("重置用户密码失败")
	}

	key := l.svcCtx.RedisKey + "user"
	filed := strconv.FormatInt(in.Id, 10)
	_, _ = l.svcCtx.Redis.HdelCtx(l.ctx, key, filed)
	return &sysclient.ReSetPasswordResp{}, nil
}
