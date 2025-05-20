package member_level

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberLevelListLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/13 13:39
*/
type QueryMemberLevelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLevelListLogic {
	return &QueryMemberLevelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberLevelList 查询会员等级列表
func (l *QueryMemberLevelListLogic) QueryMemberLevelList(req *types.QueryMemberLevelListReq) (resp *types.QueryMemberLevelListResp, err error) {
	result, err := l.svcCtx.MemberLevelService.QueryMemberLevelList(l.ctx, &umsclient.QueryMemberLevelListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Name:     strings.TrimSpace(req.Name),
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员等级列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMemberLevelListData

	for _, detail := range result.List {
		list = append(list, &types.QueryMemberLevelListData{
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
		})
	}

	return &types.QueryMemberLevelListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员等级成功",
	}, nil
}
