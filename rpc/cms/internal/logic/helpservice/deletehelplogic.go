package helpservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHelpLogic 删除帮助
/*
Author: LiuFeiHua
Date: 2025/01/23 15:23:59
*/
type DeleteHelpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpLogic {
	return &DeleteHelpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHelp 删除帮助
func (l *DeleteHelpLogic) DeleteHelp(in *cmsclient.DeleteHelpReq) (*cmsclient.DeleteHelpResp, error) {
	q := query.CmsHelp

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除帮助失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除帮助失败")
	}

	logc.Infof(l.ctx, "删除帮助成功,参数：%+v", in)
	return &cmsclient.DeleteHelpResp{}, nil
}
