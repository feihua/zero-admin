package logic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberProductCollectionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCollectionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCollectionAddLogic {
	return &MemberProductCollectionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCollectionAddLogic) MemberProductCollectionAdd(in *umsclient.MemberProductCollectionAddReq) (*umsclient.MemberProductCollectionAddResp, error) {
	_, err := l.svcCtx.UmsMemberProductCollectionModel.Insert(l.ctx, &umsmodel.UmsMemberProductCollection{
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

	return &umsclient.MemberProductCollectionAddResp{}, nil
}
