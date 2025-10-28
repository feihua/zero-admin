package userservicelogic

import (
	"context"
	"errors"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserDetailLogic 查询用户详情
/*
Author: LiuFeiHua
Date: 2024/5/30 14:33
*/
type QueryUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryUserDetail 查询用户详情
func (l *QueryUserDetailLogic) QueryUserDetail(in *sysclient.QueryUserDetailReq) (*sysclient.QueryUserDetailResp, error) {
	idStr := strconv.FormatInt(in.Id, 10)
	key := l.svcCtx.RedisKey + "user"
	cachedData, _ := l.svcCtx.Redis.HgetCtx(l.ctx, key, idStr)

	var cached sysclient.QueryUserDetailResp
	if sonic.Unmarshal([]byte(cachedData), &cached) == nil {
		return &cached, nil
	}

	item, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.ID.Eq(in.Id)).First()

	// 1.判断用户是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("用户不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询用户异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询用户异常")
	}

	var postIds = make([]int64, 0)
	post := query.SysUserPost
	err = post.WithContext(l.ctx).Select(post.PostID).Where(post.UserID.Eq(in.Id)).Scan(&postIds)
	if err != nil {
		logc.Errorf(l.ctx, "查询用户岗位失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户详情失败")
	}

	var roleIds []int64
	role := query.SysUserRole
	err = role.WithContext(l.ctx).Select(role.RoleID).Where(role.UserID.Eq(in.Id)).Scan(&roleIds)
	if err != nil {
		logc.Errorf(l.ctx, "查询用户角色失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户详情失败")
	}

	data := &sysclient.QueryUserDetailResp{
		Id:            item.ID,                                    // 用户id
		Mobile:        item.Mobile,                                // 手机号码
		UserName:      item.UserName,                              // 用户账号
		NickName:      item.NickName,                              // 用户昵称
		UserType:      item.UserType,                              // 用户类型（00系统用户）
		Avatar:        item.Avatar,                                // 头像路径
		Email:         item.Email,                                 // 用户邮箱
		Status:        item.Status,                                // 状态(1:正常，0:禁用)
		DeptId:        item.DeptID,                                // 部门ID
		LoginIp:       item.LoginIP,                               // 最后登录IP
		LoginDate:     time_util.TimeToString(item.LoginDate),     // 最后登录时间
		LoginBrowser:  item.LoginBrowser,                          // 浏览器类型
		LoginOs:       item.LoginOs,                               // 操作系统
		PwdUpdateDate: time_util.TimeToString(item.PwdUpdateDate), // 密码最后更新时间
		Remark:        item.Remark,                                // 备注
		DelFlag:       item.DelFlag,                               // 删除标志（0代表删除 1代表存在）
		CreateBy:      item.CreateBy,                              // 创建者
		CreateTime:    time_util.TimeToStr(item.CreateTime),       // 创建时间
		UpdateBy:      item.UpdateBy,                              // 更新者
		UpdateTime:    time_util.TimeToString(item.UpdateTime),    // 更新时间
		PostIds:       postIds,                                    // 岗位id
		RoleIds:       roleIds,                                    // 角色id
	}

	value, _ := sonic.Marshal(data)
	filed := strconv.FormatInt(item.ID, 10)
	_ = l.svcCtx.Redis.HsetCtx(l.ctx, key, filed, string(value))
	return data, nil
}
