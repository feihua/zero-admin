package dicttypeservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDictTypeLogic 更新字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
type UpdateDictTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictTypeLogic {
	return &UpdateDictTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDictType 更新字典表
// 1.根据字典id查询字典是否已存在
// 2.字典存在时,则直接更新字典
func (l *UpdateDictTypeLogic) UpdateDictType(in *sysclient.UpdateDictTypeReq) (*sysclient.UpdateDictTypeResp, error) {
	q := query.SysDictType.WithContext(l.ctx)

	// 1.根据字典id查询字典是否已存在
	_, err := q.Where(query.SysDictType.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典id：%d,查询字典信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典信息失败"))
	}

	// 2.字典存在时,则直接更新字典
	dict := &model.SysDictType{
		ID:         in.Id,
		DictName:   in.DictName,
		DictType:   in.DictType,
		DictStatus: in.DictStatus,
		Remark:     in.Remark,
		IsSystem:   in.IsSystem,
		UpdateBy:   in.UpdateBy,
	}
	_, err = q.Updates(dict)

	if err != nil {
		logc.Errorf(l.ctx, "更新字典信息失败,参数:%+v,异常:%s", dict, err.Error())
		return nil, errors.New("更新字典信息失败")
	}

	return &sysclient.UpdateDictTypeResp{}, nil
}
