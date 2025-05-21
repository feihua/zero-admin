package memberinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberInfoLogic 删除会员信息
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type DeleteMemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberInfoLogic {
	return &DeleteMemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberInfo 删除会员信息
func (l *DeleteMemberInfoLogic) DeleteMemberInfo(in *umsclient.DeleteMemberInfoReq) (*umsclient.DeleteMemberInfoResp, error) {
	q := query.UmsMemberInfo

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除会员信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除会员信息失败")
	}

	return &umsclient.DeleteMemberInfoResp{}, nil
}
