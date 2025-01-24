package membertagservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberTagLogic 删除用户标签
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
*/
type DeleteMemberTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberTagLogic {
	return &DeleteMemberTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberTag 删除用户标签
func (l *DeleteMemberTagLogic) DeleteMemberTag(in *umsclient.DeleteMemberTagReq) (*umsclient.DeleteMemberTagResp, error) {
	q := query.UmsMemberTag

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除用户标签失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除用户标签失败")
	}

	logc.Infof(l.ctx, "删除用户标签成功,参数：%+v", in)
	return &umsclient.DeleteMemberTagResp{}, nil
}
