package history

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ReadHistoryListLogic) ReadHistoryList(req *types.ReadHistoryListReq) (resp *types.ReadHistoryListResp, err error) {
	historyList, _ := l.svcCtx.MemberReadHistoryService.MemberReadHistoryList(l.ctx, &umsclient.MemberReadHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		MemberId: l.ctx.Value("memberId").(int64),
	})

	var list []*types.ReadHistoryList

	for _, member := range historyList.List {
		list = append(list, &types.ReadHistoryList{
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
		Data: types.ReadHistoryListData{
			Total: historyList.Total,
			Pages: historyList.Total,
			Limit: req.PageSize,
			Page:  req.Current,
			List:  nil,
		},
	}, nil
}
