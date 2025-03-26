package growthchangehistoryservicelogic

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

// AddGrowthChangeHistoryLogic 添加成长值变化历史记录
/*
Author: LiuFeiHua
Date: 2024/6/11 14:24
*/
type AddGrowthChangeHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGrowthChangeHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGrowthChangeHistoryLogic {
	return &AddGrowthChangeHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddGrowthChangeHistory 添加成长值变化历史记录
func (l *AddGrowthChangeHistoryLogic) AddGrowthChangeHistory(in *umsclient.AddGrowthChangeHistoryReq) (*umsclient.AddGrowthChangeHistoryResp, error) {
	err := query.UmsGrowthChangeHistory.WithContext(l.ctx).Create(&model.UmsGrowthChangeHistory{
		MemberID:    in.MemberId,    // 会员id
		ChangeType:  in.ChangeType,  // 改变类型：0->增加；1->减少
		ChangeCount: in.ChangeCount, // 积分改变数量
		OperateMan:  in.OperateMan,  // 操作人员
		OperateNote: in.OperateNote, // 操作备注
		SourceType:  in.SourceType,  // 积分来源：0->购物；1->管理员修改
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加成长值变化历史记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加成长值变化历史记录失败")
	}

	return &umsclient.AddGrowthChangeHistoryResp{}, nil
}
