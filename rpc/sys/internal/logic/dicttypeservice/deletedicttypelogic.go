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
// 1.查询字典是否存在
// 2.字典类型下是否有字典数据
func (l *DeleteDictTypeLogic) DeleteDictType(in *sysclient.DeleteDictTypeReq) (*sysclient.DeleteDictTypeResp, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {
		q := tx.SysDictType
		q1 := tx.SysDictItem

		ids := in.Ids
		for _, id := range ids {
			// 1.查询字典是否存在
			dictType, err := q.WithContext(l.ctx).Where(q.ID.Eq(id)).First()
			if err != nil {
				return err
			}

			// 2.字典类型下是否有字典数据
			count, err := q1.WithContext(l.ctx).Where(q1.DictType.Eq(dictType.DictType)).Count()
			if err != nil {
				return err
			}
			if count > 0 {
				return errors.New("字典类型下有字典数据，不允许删除")
			}
		}

		_, err := q.WithContext(l.ctx).Where(q.ID.In(ids...)).Delete()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除字典类型失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除字典类型失败")
	}

	return &sysclient.DeleteDictTypeResp{}, nil
}
