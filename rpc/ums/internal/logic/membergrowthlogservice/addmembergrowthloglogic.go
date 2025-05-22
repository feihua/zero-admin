package membergrowthlogservicelogic

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

// AddMemberGrowthLogLogic 添加会员成长值记录
/*
Author: LiuFeiHua
Date: 2025/05/22 15:25:56
*/
type AddMemberGrowthLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberGrowthLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberGrowthLogLogic {
	return &AddMemberGrowthLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberGrowthLog 添加会员成长值记录
func (l *AddMemberGrowthLogLogic) AddMemberGrowthLog(in *umsclient.AddMemberGrowthLogReq) (*umsclient.AddMemberGrowthLogResp, error) {
	q := query.UmsMemberGrowthLog

	item := &model.UmsMemberGrowthLog{
		MemberID:     in.MemberId,     // 会员ID
		ChangeType:   in.ChangeType,   // 变更类型：1-添加成长值，2-减少成长值
		ChangeGrowth: in.ChangeGrowth, // 变更成长值
		SourceType:   in.SourceType,   // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
		Description:  in.Description,  // 描述
		OperateMan:   in.OperateMan,   // 操作人员
		OperateNote:  in.OperateNote,  // 操作备注
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员成长值记录失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员成长值记录失败")
	}

	return &umsclient.AddMemberGrowthLogResp{}, nil
}
