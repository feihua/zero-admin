package member_level

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberLevelDetailLogic 查询会员等级表详情
/*
Author: 刘飞华
Date: 2025/02/05 10:34:53
*/
type QueryMemberLevelDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberLevelDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLevelDetailLogic {
	return &QueryMemberLevelDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberLevelDetail 查询会员等级表详情
func (l *QueryMemberLevelDetailLogic) QueryMemberLevelDetail(req *types.QueryMemberLevelDetailReq) (resp *types.QueryMemberLevelDetailResp, err error) {

	detail, err := l.svcCtx.MemberLevelService.QueryMemberLevelDetail(l.ctx, &umsclient.QueryMemberLevelDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员等级表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberLevelDetailData{
		Id:           detail.Id,           // 主键ID
		Name:         detail.Name,         // 等级名称
		Level:        detail.Level,        // 等级
		GrowthPoint:  detail.GrowthPoint,  // 升级所需成长值
		DiscountRate: detail.DiscountRate, // 折扣率(0-100)
		FreeFreight:  detail.FreeFreight,  // 是否免运费
		CommentExtra: detail.CommentExtra, // 是否可评论获取奖励
		Privileges:   detail.Privileges,   // 会员特权JSON
		Remark:       detail.Remark,       // 备注
		IsEnabled:    detail.IsEnabled,    // 是否启用
		CreateBy:     detail.CreateBy,     // 创建人ID
		CreateTime:   detail.CreateTime,   // 创建时间
		UpdateBy:     detail.UpdateBy,     // 更新人ID
		UpdateTime:   detail.UpdateTime,   // 更新时间
		IsDeleted:    detail.IsDeleted,    // 是否删除
	}
	return &types.QueryMemberLevelDetailResp{
		Code:    "000000",
		Message: "查询会员等级表成功",
		Data:    data,
	}, nil
}
