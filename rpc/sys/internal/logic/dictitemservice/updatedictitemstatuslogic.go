package dictitemservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDictItemStatusLogic 更新字典数据状态
/*
Author: LiuFeiHua
Date: 2024/5/30 10:23
*/
type UpdateDictItemStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictItemStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemStatusLogic {
	return &UpdateDictItemStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDictItemStatus 更新字典数据状态
func (l *UpdateDictItemStatusLogic) UpdateDictItemStatus(in *sysclient.UpdateDictItemStatusReq) (*sysclient.UpdateDictItemStatusResp, error) {
	q := query.SysDictItem

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "根据字典数据ids：%+v,更新字典数据状态失败,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新字典数据状态失败"))
	}

	ids := make([]string, 0, len(in.Ids))
	for _, id := range in.Ids {
		ids = append(ids, strconv.FormatInt(id, 10))
	}
	key := l.svcCtx.RedisKey + "dict:item"
	_, _ = l.svcCtx.Redis.HdelCtx(l.ctx, key, ids...)
	return &sysclient.UpdateDictItemStatusResp{}, nil
}
