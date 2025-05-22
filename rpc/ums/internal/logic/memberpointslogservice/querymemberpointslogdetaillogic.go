package memberpointslogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryMemberPointsLogDetailLogic 查询会员积分记录详情
/*
Author: LiuFeiHua
Date: 2025/05/22 15:25:56
*/
type QueryMemberPointsLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberPointsLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberPointsLogDetailLogic {
	return &QueryMemberPointsLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberPointsLogDetail 查询会员积分记录详情
func (l *QueryMemberPointsLogDetailLogic) QueryMemberPointsLogDetail(in *umsclient.QueryMemberPointsLogDetailReq) (*umsclient.QueryMemberPointsLogDetailResp, error) {
	item, err := query.UmsMemberPointsLog.WithContext(l.ctx).Where(query.UmsMemberPointsLog.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员积分记录不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员积分记录不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员积分记录异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员积分记录异常")
	}

	data := &umsclient.QueryMemberPointsLogDetailResp{
		Id:           item.ID,                                       //
		MemberId:     item.MemberID,                                 // 会员ID
		ChangeType:   item.ChangeType,                               // 变更类型：1-添加积分，2-减少积分
		ChangePoints: item.ChangePoints,                             // 变更积分
		SourceType:   item.SourceType,                               // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
		Description:  item.Description,                              // 描述
		OperateMan:   item.OperateMan,                               // 操作人员
		OperateNote:  item.OperateNote,                              // 操作备注
		CreateTime:   item.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间
	}

	return data, nil
}
