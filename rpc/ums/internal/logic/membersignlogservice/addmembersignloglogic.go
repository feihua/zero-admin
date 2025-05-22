package membersignlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// AddMemberSignLogLogic 添加会员签到记录
/*
Author: LiuFeiHua
Date: 2025/05/22 14:12:14
*/
type AddMemberSignLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberSignLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberSignLogLogic {
	return &AddMemberSignLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberSignLog 添加会员签到记录
func (l *AddMemberSignLogLogic) AddMemberSignLog(in *umsclient.AddMemberSignLogReq) (*umsclient.AddMemberSignLogResp, error) {
	q := query.UmsMemberSignLog
	signDate, _ := time.Parse("2006-01-02", in.SignDate)
	item := &model.UmsMemberSignLog{
		MemberID:     in.MemberId,     // 会员ID
		SignDate:     signDate,        // 签到日期
		ContinueDays: in.ContinueDays, // 连续签到天数
		Points:       in.Points,       // 获得积分
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员签到记录失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员签到记录失败")
	}

	return &umsclient.AddMemberSignLogResp{}, nil
}
