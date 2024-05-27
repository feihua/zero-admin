package dictservicelogic

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

// DictUpdateLogic 更新字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
type DictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictUpdate 更新字典信息
// 1.根据字典id查询字典是否已存在
// 2.如果字典不存在,则直接返回
// 3.字典存在时,则直接更新字典
func (l *DictUpdateLogic) DictUpdate(in *sysclient.DictUpdateReq) (*sysclient.DictUpdateResp, error) {
	q := query.SysDict.WithContext(l.ctx)

	// 1.根据字典名称查询字典是否已存在
	label := in.Label
	count, err := q.Where(query.SysDict.Label.Eq(label)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典名称：%s,查询字典信息失败,异常:%s", label, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典信息失败"))
	}

	//2.如果字典不存在,则直接返回
	if count == 0 {
		logc.Errorf(l.ctx, "字典信息不存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("字典：%s,不存在", label))
	}

	// 3.字典存在时,则直接更新字典
	dict := &model.SysDict{
		ID:          in.Id,
		Value:       in.Value,
		Label:       label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        in.Sort,
		UpdateBy:    in.UpdateBy,
		Remarks:     in.Remarks,
		DelFlag:     in.DelFlag,
	}
	_, err = q.Updates(dict)

	if err != nil {
		logc.Errorf(l.ctx, "更新字典信息失败,参数:%+v,异常:%s", dict, err.Error())
		return nil, errors.New("更新字典信息失败")
	}

	return &sysclient.DictUpdateResp{}, nil
}
