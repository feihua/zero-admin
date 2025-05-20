package memberlevelservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberLevelLogic 添加会员等级
/*
Author: LiuFeiHua
Date: 2024/6/11 14:16
*/
type AddMemberLevelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberLevelLogic {
	return &AddMemberLevelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberLevel 添加会员等级
func (l *AddMemberLevelLogic) AddMemberLevel(in *umsclient.AddMemberLevelReq) (*umsclient.AddMemberLevelResp, error) {
	q := query.UmsMemberLevel
	count, err := q.WithContext(l.ctx).Where(q.Name.Eq(in.Name)).Count()
	if count > 0 {
		return nil, errors.New(fmt.Sprintf("会员等级名称：%s,已存在", in.Name))
	}
	count, err = q.WithContext(l.ctx).Where(q.Level.Eq(in.Level)).Count()
	if count > 0 {
		return nil, errors.New(fmt.Sprintf("会员等级：%d,已存在", in.Level))
	}

	err = q.WithContext(l.ctx).Create(&model.UmsMemberLevel{
		Name:         in.Name,                  // 等级名称
		Level:        in.Level,                 // 等级
		GrowthPoint:  in.GrowthPoint,           // 升级所需成长值
		DiscountRate: float64(in.DiscountRate), // 折扣率(0-100)
		FreeFreight:  in.FreeFreight,           // 是否免运费
		CommentExtra: in.CommentExtra,          // 是否可评论获取奖励
		Privileges:   in.Privileges,            // 会员特权JSON
		Remark:       in.Remark,                // 备注
		IsEnabled:    in.IsEnabled,             // 是否启用
		CreateBy:     in.CreateBy,              // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员等级失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加会员等级失败")
	}

	return &umsclient.AddMemberLevelResp{}, nil
}
