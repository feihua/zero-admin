package membertaskrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryMemberTaskRelationDetailLogic 查询会员任务关联详情
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTaskRelationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTaskRelationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTaskRelationDetailLogic {
	return &QueryMemberTaskRelationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTaskRelationDetail 查询会员任务关联详情
func (l *QueryMemberTaskRelationDetailLogic) QueryMemberTaskRelationDetail(in *umsclient.QueryMemberTaskRelationDetailReq) (*umsclient.QueryMemberTaskRelationDetailResp, error) {
	item, err := query.UmsMemberTaskRelation.WithContext(l.ctx).Where(query.UmsMemberTaskRelation.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员任务关联不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员任务关联不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员任务关联异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员任务关联异常")
	}

	data := &umsclient.QueryMemberTaskRelationDetailResp{
		Id:         item.ID,                              // 主键ID
		MemberId:   item.MemberID,                        // 会员ID
		TaskId:     item.TaskID,                          // 任务ID
		CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
	}

	return data, nil
}
