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
	user, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询用户详情信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户详情信息失败")
	}

	var postIds []int64

	sql := `select t1.post_id
			from sys_user_post t1
			where  t1.user_id = ?;`
	l.svcCtx.DB.Raw(sql, in.Id).Select("post_id").Scan(&postIds)
	data := &sysclient.QueryUserDetailResp{
		Avatar:     user.Avatar,
		CreateBy:   user.CreateBy,
		CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
		DeptId:     user.DeptID,
		Email:      user.Email,
		Id:         user.ID,
		UserStatus: user.UserStatus,
		LoginIp:    user.LoginIP,
		LoginTime:  time_util.TimeToString(user.LoginTime),
		Mobile:     user.Mobile,
		NickName:   user.NickName,
		Remark:     user.Remark,
		UpdateBy:   user.UpdateBy,
		UpdateTime: time_util.TimeToString(user.UpdateTime),
		UserName:   user.UserName,
		PostIds:    postIds,
	}

	logc.Infof(l.ctx, "查询用户详情信息,参数：%+v,响应：%+v", in, data)
	return data, nil
}
