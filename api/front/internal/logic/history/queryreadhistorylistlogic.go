package history

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryReadHistoryListLogic
/*
Author: LiuFeiHua
Date: 2024/5/16 10:45
*/
type QueryReadHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryReadHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryReadHistoryListLogic {
	return &QueryReadHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryReadHistoryList 查询会员浏览商品的记录
func (l *QueryReadHistoryListLogic) QueryReadHistoryList() (resp *types.ReadHistoryListResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	historyList, _ := l.svcCtx.MemberReadHistoryService.MemberReadHistoryList(l.ctx, &umsclient.MemberReadHistoryListReq{
		Current:  1,
		PageSize: 100,
		MemberId: memberId,
	})

	var list []types.ReadHistoryList

	for _, member := range historyList.List {
		list = append(list, types.ReadHistoryList{
			Id:              member.Id,
			MemberId:        member.MemberId,
			MemberNickName:  member.MemberNickName,
			MemberIcon:      member.MemberIcon,
			ProductId:       member.ProductId,
			ProductName:     member.ProductName,
			ProductPic:      member.ProductPic,
			ProductSubTitle: member.ProductSubTitle,
			ProductPrice:    member.ProductPrice,
			CreateTime:      member.CreateTime,
		})
	}

	return &types.ReadHistoryListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
