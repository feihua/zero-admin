package history

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReadHistoryListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 16:34
*/
type ReadHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadHistoryListLogic {
	return &ReadHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReadHistoryList 查询会员浏览商品的记录
func (l *ReadHistoryListLogic) ReadHistoryList() (resp *types.ReadHistoryListResp, err error) {
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
