package roleservicelogic

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"strconv"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleDetailLogic 查询角色详情
/*
Author: LiuFeiHua
Date: 2024/5/30 14:18
*/
type QueryRoleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleDetailLogic {
	return &QueryRoleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryRoleDetail 查询角色详情
// 1.判断角色是否存在
func (l *QueryRoleDetailLogic) QueryRoleDetail(in *sysclient.QueryRoleDetailReq) (*sysclient.QueryRoleDetailResp, error) {
	idStr := strconv.FormatInt(in.Id, 10)
	key := l.svcCtx.RedisKey + "role"
	cachedData, _ := l.svcCtx.Redis.HgetCtx(l.ctx, key, idStr)

	var cached sysclient.QueryRoleDetailResp
	if sonic.Unmarshal([]byte(cachedData), &cached) == nil {
		return &cached, nil
	}
	item, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.ID.Eq(in.Id)).First()

	// 1.判断角色是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("角色不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询角色异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询角色异常")
	}

	data := &sysclient.QueryRoleDetailResp{
		Id:         item.ID,                                 // 角色id
		RoleName:   item.RoleName,                           // 名称
		RoleKey:    item.RoleKey,                            // 角色权限字符串
		DataScope:  item.DataScope,                          // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:     item.Status,                             // 状态(1:正常，0:禁用)
		Remark:     item.Remark,                             // 备注
		DelFlag:    item.DelFlag,                            // 删除标志（0代表删除 1代表存在）
		CreateBy:   item.CreateBy,                           // 创建者
		CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                           // 更新者
		UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	value, _ := sonic.Marshal(data)
	filed := strconv.FormatInt(item.ID, 10)
	_ = l.svcCtx.Redis.HsetCtx(l.ctx, key, filed, string(value))
	return data, nil
}
