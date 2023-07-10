package history

import (
	"context"
	"time"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReadHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReadHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReadHistoryAddLogic {
	return &AddReadHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReadHistoryAddLogic) AddReadHistoryAdd(req *types.AddReadHistoryReq) (resp *types.AddReadHistoryResp, err error) {
	member, _ := l.svcCtx.MemberService.QueryMemberById(l.ctx, &umsclient.MemberByIdReq{Id: l.ctx.Value("memberId").(int64)})

	_, err = l.svcCtx.MemberReadHistoryService.MemberReadHistoryAdd(l.ctx, &umsclient.MemberReadHistoryAddReq{
		MemberId:        member.Id,
		MemberNickName:  member.Nickname,
		MemberIcon:      member.Icon,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		ProductPic:      req.ProductPic,
		ProductSubTitle: req.ProductSubTitle,
		ProductPrice:    req.ProductPrice,
		CreateTime:      time.Now().Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		return nil, err
	}
	return &types.AddReadHistoryResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
