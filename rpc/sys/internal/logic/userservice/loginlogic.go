package userservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogic 用户登录
/*
Author: LiuFeiHua
Date: 2023/12/18 14:08
*/
type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 根据用户名(用名称或者手机号)和密码登录
// 1.判断用户是否存在
// 2.判断密码是否正确
// 3.查询用户权限
// 4.获取部门信息
// 5.生成token
// 6.保存登录日志
// 7.更新登录时间
func (l *LoginLogic) Login(in *sysclient.LoginReq) (*sysclient.LoginResp, error) {
	q := query.SysUser
	user, err := q.WithContext(l.ctx).Where(q.UserName.Eq(in.Account)).Or(q.Mobile.Eq(in.Account)).First()

	// 1.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		l.savaLoginLog(in, "error", fmt.Sprintf("用户不存在: %+v", in))
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		l.savaLoginLog(in, "error", fmt.Sprintf("查询用户信息异常: %+v", in))
		return nil, errors.New("查询用户信息异常")
	}

	// 2.判断密码是否正确
	if user.Password != in.Password {
		l.savaLoginLog(in, "error", fmt.Sprintf("用户密码不正确: %+v", in))
		return nil, errors.New("用户密码不正确")
	}

	// 3.查询用户权限
	apiUrls := l.queryApiUrls(user.ID)

	if len(apiUrls) == 0 {
		l.savaLoginLog(in, "error", fmt.Sprintf("用户还没有分配角色或者还没有分配角色权限: %+v", in))
		return nil, errors.New("用户还没有分配角色或者还没有分配角色权限")
	}

	// 4.获取部门信息
	sysDept := query.SysDept
	dept, err := sysDept.WithContext(l.ctx).Select(sysDept.DeptName).Where(sysDept.ID.Eq(user.DeptID)).First()
	if err != nil {
		l.savaLoginLog(in, "error", fmt.Sprintf("查询用户部门信息异常: %+v", in))
		return nil, errors.New("查询用户部门信息异常")
	}
	// 5.生成token
	jwtToken, err := l.createToken(user.ID, user.DeptID, user.UserName, dept.DeptName)

	if err != nil {
		l.savaLoginLog(in, "error", fmt.Sprintf("生成token失败: %+v", in))
		return nil, errors.New("生成token失败")
	}

	// 6.保存登录日志
	l.savaLoginLog(in, "success", "登录成功")

	// 7.更新登录时间
	now := time.Now()
	_, _ = q.WithContext(l.ctx).Where(q.ID.Eq(user.ID)).Updates(&model.SysUser{
		LoginTime:    &now,         // 登录时间
		LoginIP:      in.IpAddress, // 登录ip
		LoginOs:      in.Os,        // 登录os
		LoginBrowser: in.Browser,   // 登录浏览器
	})

	return &sysclient.LoginResp{
		Id:          user.ID,       // 用户id
		UserName:    user.UserName, // 用户名
		AccessToken: jwtToken,      // token
		ApiUrls:     apiUrls,       // 权限
	}, nil
}

// 3.查询权限(判断是不是超级管理员：角色is_admin:1是超级管理员的角色)
func (l *LoginLogic) queryApiUrls(userId int64) []string {
	db := l.svcCtx.DB

	var apiUrls []string
	// 4.1判断是不是超级管理员，则是超级管理员，拿所有权限
	if common.IsAdmin(l.ctx, userId, l.svcCtx.DB) {
		sql := `select background_url from sys_menu where background_url != ''`
		db.WithContext(l.ctx).Raw(sql).Select("background_url").Scan(&apiUrls)
		return apiUrls
	}

	sql := `select sm.background_url
			from sys_user_role sur
					 left join sys_role sr on sur.role_id = sr.id
					 left join sys_role_menu srm on sr.id = srm.role_id
					 left join sys_menu sm on srm.menu_id = sm.id
			where sur.user_id = ?`
	db.WithContext(l.ctx).Raw(sql, userId).Select("background_url").Scan(&apiUrls)

	return apiUrls
}

// 保存登录日志
func (l *LoginLogic) savaLoginLog(in *sysclient.LoginReq, status, errorMsg string) {
	_ = query.SysLoginLog.WithContext(l.ctx).Create(&model.SysLoginLog{
		UserName:    in.Account,   // 用户名
		LoginStatus: status,       // 登录状态
		IPAddress:   in.IpAddress, // IP地址
		Browser:     in.Browser,   // 浏览器
		Os:          in.Os,        // 操作信息
		ErrorMsg:    errorMsg,     // 登录失败信息
		LoginTime:   time.Now(),   // 登录时间
	})
}

// 生成jwt的token
func (l *LoginLogic) createToken(userId, deptID int64, userName, deptName string) (string, error) {
	now := time.Now().Unix()                         // 当前时间
	accessExpire := l.svcCtx.Config.JWT.AccessExpire // token过期时间
	accessSecret := l.svcCtx.Config.JWT.AccessSecret // token密钥

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire       // 过期时间
	claims["iat"] = now                      // 签发时间
	claims["userId"] = userId                // 用户id
	claims["userName"] = userName            // 用户名
	claims["deptID"] = deptID                // 部门id
	claims["deptName"] = deptName            // 部门名称
	token := jwt.New(jwt.SigningMethodHS256) // 创建token
	token.Claims = claims                    // 设置claims
	return token.SignedString([]byte(accessSecret))
}
