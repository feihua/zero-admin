package memberreadhistoryservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberReadHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReadHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReadHistoryAddLogic {
	return &MemberReadHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReadHistoryAddLogic) MemberReadHistoryAdd(in *umsclient.MemberReadHistoryAddReq) (*umsclient.MemberReadHistoryAddResp, error) {
	_, err := l.svcCtx.UmsMemberReadHistoryModel.Insert(l.ctx, &umsmodel.UmsMemberReadHistory{
		MemberId:        in.MemberId,
		MemberNickName:  in.MemberNickName,
		MemberIcon:      in.MemberIcon,
		ProductId:       in.ProductId,
		ProductName:     in.ProductName,
		ProductPic:      in.ProductPic,
		ProductSubTitle: sql.NullString{String: in.ProductSubTitle, Valid: true},
		ProductPrice:    in.ProductPrice,
		CreateTime:      time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReadHistoryAddResp{}, nil
}
