package memberproductcollectionservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberProductCollectionLogic 删除商品收藏/清空当前用户商品收藏列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:08
*/
type DeleteMemberProductCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberProductCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberProductCollectionLogic {
	return &DeleteMemberProductCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberProductCollection 删除商品收藏/清空当前用户商品收藏列表
func (l *DeleteMemberProductCollectionLogic) DeleteMemberProductCollection(in *umsclient.DeleteMemberProductCollectionReq) (*umsclient.DeleteMemberProductCollectionResp, error) {
	for _, id := range in.Ids {
		_, err := l.svcCtx.MemberProductCollectionModel.Deletes(l.ctx, id, in.MemberId)
		if err != nil {
			logc.Errorf(l.ctx, "删除商品收藏失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("删除商品收藏失败")
		}
	}

	return &umsclient.DeleteMemberProductCollectionResp{}, nil
}
