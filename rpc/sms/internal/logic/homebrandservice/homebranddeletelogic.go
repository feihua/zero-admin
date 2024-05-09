package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeBrandDeleteLogic 首页品牌
/*
Author: LiuFeiHua
Date: 2024/5/8 10:18
*/
type HomeBrandDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandDeleteLogic {
	return &HomeBrandDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeBrandDelete 删除首页品牌
func (l *HomeBrandDeleteLogic) HomeBrandDelete(in *smsclient.HomeBrandDeleteReq) (*smsclient.HomeBrandDeleteResp, error) {
	q := query.SmsHomeBrand
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.HomeBrandDeleteResp{}, nil
}
