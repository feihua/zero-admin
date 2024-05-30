package dicttypeservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDictTypeStatusLogic 更新字典类型表状态
/*
Author: LiuFeiHua
Date: 2024/5/30 10:23
*/
type UpdateDictTypeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictTypeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictTypeStatusLogic {
	return &UpdateDictTypeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDictTypeStatus 更新字典类型表状态
func (l *UpdateDictTypeStatusLogic) UpdateDictTypeStatus(in *sysclient.UpdateDictTypeStatusReq) (*sysclient.UpdateDictTypeStatusResp, error) {
	q := query.SysDictType

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.DictStatus, in.DictStatus)

	if err != nil {
		logc.Errorf(l.ctx, "根据字典类型ids：%+v,更新字典类型表状态失败,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新字典类型表状态失败"))
	}

	return &sysclient.UpdateDictTypeStatusResp{}, nil
}
