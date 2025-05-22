package memberpointslogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberPointsLogLogic 添加会员积分记录
/*
Author: LiuFeiHua
Date: 2025/05/22 15:25:56
*/
type AddMemberPointsLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberPointsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberPointsLogLogic {
	return &AddMemberPointsLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberPointsLog 添加会员积分记录
func (l *AddMemberPointsLogLogic) AddMemberPointsLog(in *umsclient.AddMemberPointsLogReq) (*umsclient.AddMemberPointsLogResp, error) {
	q := query.UmsMemberPointsLog

	item := &model.UmsMemberPointsLog{
		MemberID:     in.MemberId,     // 会员ID
		ChangeType:   in.ChangeType,   // 变更类型：1-添加积分，2-减少积分
		ChangePoints: in.ChangePoints, // 变更积分
		SourceType:   in.SourceType,   // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
		Description:  in.Description,  // 描述
		OperateMan:   in.OperateMan,   // 操作人员
		OperateNote:  in.OperateNote,  // 操作备注
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员积分记录失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员积分记录失败")
	}

	return &umsclient.AddMemberPointsLogResp{}, nil
}
