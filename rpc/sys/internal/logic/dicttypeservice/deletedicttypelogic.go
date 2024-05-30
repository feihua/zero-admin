package dicttypeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteDictTypeLogic 删除字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type DeleteDictTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictTypeLogic {
	return &DeleteDictTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteDictType 删除字典类型
func (l *DeleteDictTypeLogic) DeleteDictType(in *sysclient.DeleteDictTypeReq) (*sysclient.DeleteDictTypeResp, error) {
	q := query.SysDictType
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除字典类型信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除字典类型信息失败")
	}

	return &sysclient.DeleteDictTypeResp{}, nil
}
