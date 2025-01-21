package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserDetailLogic 查询用户详情信息
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

// QueryUserDetail 查询用户详情信息
func (l *QueryUserDetailLogic) QueryUserDetail(in *sysclient.QueryUserDetailReq) (*sysclient.QueryUserDetailResp, error) {
	item, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询用户详情信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户详情信息失败")
	}

	var postIds []int64

	sql := `select t1.post_id
			from sys_user_post t1
			where  t1.user_id = ?;`
	l.svcCtx.DB.Raw(sql, in.Id).Select("post_id").Scan(&postIds)
	loginTime := item.LoginTime.Format("2006-01-02 15:04:05")
	createTime := item.CreateTime.Format("2006-01-02 15:04:05")
	data := &sysclient.QueryUserDetailResp{
		Id:           item.ID,                                 // 编号
		UserName:     item.UserName,                           // 用户名
		NickName:     item.NickName,                           // 昵称
		Avatar:       item.Avatar,                             // 头像
		Email:        item.Email,                              // 邮箱
		Mobile:       item.Mobile,                             // 手机号
		UserStatus:   item.UserStatus,                         // 帐号状态（0正常 1停用）
		DeptId:       item.DeptID,                             // 部门id
		Remark:       item.Remark,                             // 备注信息
		IsDeleted:    item.IsDeleted,                          // 是否删除  0：否  1：是
		LoginTime:    loginTime,                               // 登录时间
		LoginIp:      item.LoginIP,                            // 登录ip
		LoginOs:      item.LoginOs,                            // 登录os
		LoginBrowser: item.LoginBrowser,                       // 登录浏览器
		CreateBy:     item.CreateBy,                           // 创建者
		CreateTime:   createTime,                              // 创建时间
		UpdateBy:     item.UpdateBy,                           // 更新者
		UpdateTime:   time_util.TimeToString(item.UpdateTime), // 更新时间
		PostIds:      postIds,
	}

	logc.Infof(l.ctx, "查询用户详情信息,参数：%+v,响应：%+v", in, data)
	return data, nil
}
