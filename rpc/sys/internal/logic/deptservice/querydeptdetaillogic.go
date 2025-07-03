package deptservicelogic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryDeptDetailLogic 查询部门信息详情
/*
Author: LiuFeiHua
Date: 2024/5/30 10:13
*/
type QueryDeptDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDeptDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptDetailLogic {
	return &QueryDeptDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDeptDetail 查询部门信息详情
func (l *QueryDeptDetailLogic) QueryDeptDetail(in *sysclient.QueryDeptDetailReq) (*sysclient.QueryDeptDetailResp, error) {
	idStr := strconv.FormatInt(in.Id, 10)
	key := l.svcCtx.RedisKey + "dept"
	cachedData, _ := l.svcCtx.Redis.HgetCtx(l.ctx, key, idStr)

	var cached sysclient.QueryDeptDetailResp
	if json.Unmarshal([]byte(cachedData), &cached) == nil {
		return &cached, nil
	}
	dept, err := query.SysDept.WithContext(l.ctx).Where(query.SysDept.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "部门不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("部门不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询部门异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询部门异常")
	}

	data := &sysclient.QueryDeptDetailResp{
		Id:         dept.ID,                                 // 部门id
		ParentId:   dept.ParentID,                           // 上级部门id
		Ancestors:  dept.Ancestors,                          // 祖级列表
		DeptName:   dept.DeptName,                           // 部门名称
		Sort:       dept.Sort,                               // 显示顺序
		Leader:     dept.Leader,                             // 负责人
		Phone:      dept.Phone,                              // 联系电话
		Email:      dept.Email,                              // 邮箱
		Status:     dept.Status,                             // 部门状态（0：停用，1:正常）
		DelFlag:    dept.DelFlag,                            // 删除标志（0代表删除 1代表存在）
		Remark:     dept.Remark,                             // 备注信息
		CreateBy:   dept.CreateBy,                           // 创建者
		CreateTime: time_util.TimeToStr(dept.CreateTime),    // 创建时间
		UpdateBy:   dept.UpdateBy,                           // 更新者
		UpdateTime: time_util.TimeToString(dept.UpdateTime), // 更新时间
	}

	value, _ := json.Marshal(data)
	filed := strconv.FormatInt(dept.ID, 10)
	_ = l.svcCtx.Redis.HsetCtx(l.ctx, key, filed, string(value))
	return data, nil

}
