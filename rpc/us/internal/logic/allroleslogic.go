package logic

import (
	"context"
	"go-zero-admin/common"
	"go-zero-admin/rpc/model/usmodel"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllRolesLogic {
	return &AllRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllRolesLogic) AllRoles(in *us.AllRoleReq) (*us.AllRoleResp, error) {
	var lastAll *[]usmodel.UsRoles
	all, err := RedisListAllRoles(l.svcCtx)
	count := 0

	if err == nil {
		lastAll = all
		count = len(*lastAll)
	} else {
		all, err = l.svcCtx.UsRolesModel.FindAll(0, 20)
		if err != nil{
			logx.Errorf("UsRolesModel find all error:" + err.Error())
			return &us.AllRoleResp{
				Total: int64(count),
				List:  nil,
			}, errorRoleUnRegister
		}
		logx.Info("find all size:"+ strconv.Itoa(len(*all)))
		lastAll = all
		tempCount, _ := l.svcCtx.UsRolesModel.Count()
		count = int(tempCount)

		for _, item := range *lastAll {
			//将 UsRoles list的所有id缓存在 redis中
			AddRole(l.svcCtx.RedisConn, &item)
		}
		//设置role多行id 缓存时间
		SetRoleBizCacheExpire(l.svcCtx.RedisConn, int(common.GetCommonRedisExpireSeconds()))
	}

	var list []*us.RoleData
	for _, item := range *lastAll {
		list = append(list, &us.RoleData{
			Id:       item.Id,
			RoleName: item.RoleName.String,
			Remark:   item.Remark.String,
			CreateTime: item.CreateTime.Time.String(),
			UpdateTime: item.UpdateTime.Time.String(),
		})
	}

	return &us.AllRoleResp{
		Total: int64(count),
		List:  list,
	}, nil
}
