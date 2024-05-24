package userservicelogic

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
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

// Login 根据用户名和密码登录
// 1.判断用户是否存在
// 2.判断密码是否正确
// 3.生成token
// 4.查询权限(判断是不是超级管理员：角色id:1是预留超级管理员的角色)
// 5.保存登录日志
func (l *LoginLogic) Login(in *sysclient.LoginReq) (*sysclient.LoginResp, error) {
	q := query.SysUser
	user, err := q.WithContext(l.ctx).Where(q.UserName.Eq(in.Account)).Or(q.Mobile.Eq(in.Account)).First()

	// 1.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		l.savaLoginLog(in, "error")
		logc.Errorf(l.ctx, "用户不存在,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		l.savaLoginLog(in, "error")
		logc.Errorf(l.ctx, "查询用户信息,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户信息异常")
	}

	// 2.判断密码是否正确
	if user.Password != in.Password {
		l.savaLoginLog(in, "error")
		logc.Errorf(l.ctx, "用户密码不正确,参数:%s", in.Password)
		return nil, errors.New("用户密码不正确")
	}

	// 3.生成token
	jwtToken, err := l.getJwtToken(user.ID, user.UserName)

	if err != nil {
		l.savaLoginLog(in, "error")
		logc.Errorf(l.ctx, "生成token失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("生成token失败")
	}

	//4.查询权限
	apiUrls := l.queryApiUrls(user)

	// 5.保存登录日志
	l.savaLoginLog(in, "success")
	return &sysclient.LoginResp{
		Id:          user.ID,
		UserName:    user.UserName,
		AccessToken: jwtToken,
		ApiUrls:     apiUrls,
	}, nil
}

// 4.查询权限(判断是不是超级管理员：角色id:1是预留超级管理员的角色)
func (l *LoginLogic) queryApiUrls(user *model.SysUser) []string {
	role := query.SysUserRole
	count, _ := role.WithContext(l.ctx).Where(role.RoleID.Eq(1), role.UserID.Eq(user.ID)).Count()

	var apiUrls []string
	// 4.1查询出不为0，则是超级管理员，拿所有权限
	backgroundUrl := query.SysMenu.BackgroundURL
	if count != 0 {
		_ = query.SysMenu.WithContext(l.ctx).Select(backgroundUrl).Scan(&apiUrls)
	} else {
		sql := `select sm.background_url
			from sys_user_role sur
					 left join sys_role sr on sur.role_id = sr.id
					 left join sys_role_menu srm on sr.id = srm.role_id
					 left join sys_menu sm on srm.menu_id = sm.id
			where sur.user_id = ?
			order by sm.id`
		db := l.svcCtx.DB
		_ = db.WithContext(l.ctx).Raw(sql, user.ID).Select("background_url").Scan(&apiUrls).Error
	}
	return apiUrls
}

// 保存登录日志
func (l *LoginLogic) savaLoginLog(in *sysclient.LoginReq, loginStatus string) {
	_ = query.SysLoginLog.WithContext(l.ctx).Create(&model.SysLoginLog{
		LoginName:   in.Account,
		LoginStatus: loginStatus,
		LoginIP:     in.LoginIp,
		LoginTime:   time.Now(),
	})
}

// 生成jwt的token
func (l *LoginLogic) getJwtToken(userId int64, userName string) (string, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["userId"] = userId
	claims["userName"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(accessSecret))
}
