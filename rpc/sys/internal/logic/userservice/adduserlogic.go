package userservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddUserLogic 新增用户
/*
Author: LiuFeiHua
Date: 2023/12/18 14:13
*/
type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddUser 新增用户
// 1.查询用户名称是否存在
// 2.查询手机号是否存在
// 3.查询邮箱是否存在
// 4.用户不存在时,则直接添加用户
// 5.清空用户与岗位关联
// 6.添加用户与岗位关联
// 7.清空用户与角色关联(防止脏数据)
func (l *AddUserLogic) AddUser(in *sysclient.AddUserReq) (*sysclient.AddUserResp, error) {
	q := query.SysUser

	// 1.查询用户名称是否存在
	name := in.UserName
	count, err := q.WithContext(l.ctx).Where(q.UserName.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据用户名称：%s,查询用户失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("新增用户失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("用户：%s,已存在", name))
	}

	// 2.查询手机号是否存在
	mobile := in.Mobile
	count, err = q.WithContext(l.ctx).Where(q.Mobile.Eq(mobile)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据手机号：%s,查询用户失败,异常:%s", mobile, err.Error())
		return nil, errors.New(fmt.Sprintf("新增用户失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("手机号：%s,已存在", mobile))
	}

	// 3.查询邮箱是否存在
	email := in.Email
	count, err = q.WithContext(l.ctx).Where(q.Email.Eq(email)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据邮箱：%s,查询用户失败,异常:%s", email, err.Error())
		return nil, errors.New(fmt.Sprintf("新增用户失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("邮箱：%s,已存在", email))
	}

	avatar := in.Avatar
	// 默认用户头像
	if len(avatar) == 0 {
		avatar = "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	}
	user := &model.SysUser{
		Mobile:   in.Mobile,   // 手机号码
		UserName: in.UserName, // 用户账号
		NickName: in.NickName, // 用户昵称
		UserType: in.UserType, // 用户类型（00系统用户）
		Avatar:   in.Avatar,   // 头像路径
		Email:    in.Email,    // 用户邮箱
		Password: in.Password, // 密码
		Status:   in.Status,   // 状态(1:正常，0:禁用)
		DeptID:   in.DeptId,   // 部门ID
		Remark:   in.Remark,   // 备注
		CreateBy: in.CreateBy, // 创建者
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		// 4.用户不存在时,则直接添加用户
		err = tx.SysUser.WithContext(l.ctx).Create(user)

		if err != nil {
			logc.Errorf(l.ctx, "新增用户异常,参数:%+v,异常:%s", user, err.Error())
			return err
		}

		var userPosts []*model.SysUserPost
		for _, postId := range in.PostIds {
			userPosts = append(userPosts, &model.SysUserPost{
				UserID: user.ID,
				PostID: postId,
			})
		}

		postDo := tx.SysUserPost.WithContext(l.ctx)
		// 5.清空用户与岗位关联
		_, err = postDo.Where(tx.SysUserPost.UserID.Eq(user.ID)).Delete()
		if err != nil {
			logc.Errorf(l.ctx, "删除用户与岗位关联异常,参数:%+v,异常:%s", user, err.Error())
			return err
		}

		// 6.添加用户与岗位关联
		err = postDo.CreateInBatches(userPosts, len(userPosts))
		if err != nil {
			logc.Errorf(l.ctx, "新增用户与岗位关联异常,参数:%+v,异常:%s", user, err.Error())
			return err
		}

		roleDo := tx.SysUserRole.WithContext(l.ctx)
		// 7.清空用户与角色关联(防止脏数据)
		_, err = roleDo.Where(tx.SysUserRole.UserID.Eq(user.ID)).Delete()
		if err != nil {
			logc.Errorf(l.ctx, "删除用户与角色关联异常,参数:%+v,异常:%s", user, err.Error())
			return err
		}
		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "新增用户异常,参数:%+v,异常:%s", user, err.Error())
		return nil, errors.New("新增用户异常")
	}
	return &sysclient.AddUserResp{}, nil
}
