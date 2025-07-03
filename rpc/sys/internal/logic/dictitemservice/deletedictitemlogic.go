package dictitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteDictItemLogic 删除字典数据
/*
Author: LiuFeiHua
Date: 2024/5/28 17:03
*/
type DeleteDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictItemLogic {
	return &DeleteDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteDictItem 删除字典数据
func (l *DeleteDictItemLogic) DeleteDictItem(in *sysclient.DeleteDictItemReq) (*sysclient.DeleteDictItemResp, error) {
	q := query.SysDictItem

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除字典数据信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除字典数据信息失败")
	}

	ids := make([]string, 0, len(in.Ids))
	for _, id := range in.Ids {
		ids = append(ids, strconv.FormatInt(id, 10))
	}
	key := l.svcCtx.RedisKey + "dict:item"
	_, _ = l.svcCtx.Redis.HdelCtx(l.ctx, key, ids...)
	return &sysclient.DeleteDictItemResp{}, nil
}
