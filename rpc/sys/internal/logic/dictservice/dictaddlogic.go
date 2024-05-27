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

// DictAddLogic 添加字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type DictAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictAddLogic {
	return &DictAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictAdd 添加字典信息
// 1.根据字典名称查询字典是否已存在
// 2.如果字典已存在,则直接返回
// 3.字典不存在时,则直接添加字典
func (l *DictAddLogic) DictAdd(in *sysclient.DictAddReq) (*sysclient.DictAddResp, error) {

	q := query.SysDict.WithContext(l.ctx)

	// 1.根据字典名称查询字典是否已存在
	label := in.Label
	count, err := q.Where(query.SysDict.Label.Eq(label)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典名称：%s,查询字典信息失败,异常:%s", label, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典信息失败"))
	}

	//2.如果字典已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "字典信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("字典：%s,已存在", label))
	}

	dict := &model.SysDict{
		Value:       in.Value,
		Label:       label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        in.Sort,
		CreateBy:    in.CreateBy,
		Remarks:     in.Remarks,
		DelFlag:     in.DelFlag,
	}
	err = q.Create(dict)

	if err != nil {
		logc.Errorf(l.ctx, "添加字典信息失败,参数:%+v,异常:%s", dict, err.Error())
		return nil, errors.New("添加字典信息失败")
	}

	return &sysclient.DictAddResp{}, nil
}
