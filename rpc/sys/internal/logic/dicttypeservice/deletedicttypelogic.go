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
// 1.查询字典的sys_dict_type表中的dict_type字段
// 2.删除sys_dict_type表中的字典类记录
// 3.删除sys_dict_item表中的字典数据记录
func (l *DeleteDictTypeLogic) DeleteDictType(in *sysclient.DeleteDictTypeReq) (*sysclient.DeleteDictTypeResp, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {

		// 1.查询字典的sys_dict_type中的dict_type字段
		var dictType []string
		q := tx.SysDictType
		err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Select(q.DictType).Scan(&dictType)
		if err != nil {
			return err
		}

		// 2.删除sys_dict_type中的字典类记录
		_, err = q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
		if err != nil {
			return err
		}

		// 3.删除sys_dict_item中的字典数据记录
		dictItem := tx.SysDictItem
		_, err = dictItem.WithContext(l.ctx).Where(dictItem.DictType.In(dictType...)).Delete()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除字典类型信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除字典类型信息失败")
	}

	return &sysclient.DeleteDictTypeResp{}, nil
}
