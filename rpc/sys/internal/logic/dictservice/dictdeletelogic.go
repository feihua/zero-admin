package dictservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictDeleteLogic 删除字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type DictDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDeleteLogic {
	return &DictDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictDelete 删除字典信息
func (l *DictDeleteLogic) DictDelete(in *sysclient.DictDeleteReq) (*sysclient.DictDeleteResp, error) {
	q := query.SysDict
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除字典信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除字典信息失败")
	}

	return &sysclient.DictDeleteResp{}, nil
}
