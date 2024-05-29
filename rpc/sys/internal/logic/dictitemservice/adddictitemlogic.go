package dictitemservicelogic

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

// AddDictItemLogic 添加字典项表
/*
Author: LiuFeiHua
Date: 2024/5/28 17:03
*/
type AddDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictItemLogic {
	return &AddDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddDictItem 添加字典项表
// 1.根据字典类型查询字典是否已存在
// 2.根据字典项名称查询字典项是否已存在
// 3.如果字典项已存在,则直接返回
// 4.字典项不存在时,则直接添加字典项
func (l *AddDictItemLogic) AddDictItem(in *sysclient.DictItemAddReq) (*sysclient.DictItemAddResp, error) {
	q := query.SysDictItem

	//1.根据字典类型查询字典是否已存在
	_, err := query.SysDict.WithContext(l.ctx).Where(query.SysDict.DictType.Eq(in.DictType)).First()
	if err != nil {
		logc.Errorf(l.ctx, "根据字典类型：%s,查询字典信息失败,异常:%s", in.DictType, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典信息失败"))
	}

	//2.根据字典项名称查询字典项是否已存在
	count, err := q.WithContext(l.ctx).Where(q.DictLabel.Eq(in.DictLabel), q.DictType.Eq(in.DictType)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典项名称：%s,查询字典项信息失败,异常:%s", in.DictLabel, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典项信息失败"))
	}

	//3.如果字典项已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "字典项信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("字典项：%s,已存在", in.DictLabel))
	}

	//4.字典项不存在时,则直接添加字典项
	dictItem := &model.SysDictItem{
		DictType:   in.DictType,
		DictLabel:  in.DictLabel,
		DictValue:  in.DictValue,
		DictStatus: in.DictStatus,
		DictSort:   in.DictSort,
		Remark:     in.Remark,
		IsDefault:  in.IsDefault,
		DelFlag:    1, //0：已删除  1：正常'
		CreateBy:   in.CreateBy,
	}

	err = q.WithContext(l.ctx).Create(dictItem)

	if err != nil {
		logc.Errorf(l.ctx, "添加字典项信息失败,参数:%+v,异常:%s", dictItem, err.Error())
		return nil, errors.New("添加字典项信息失败")
	}

	return &sysclient.DictItemAddResp{}, nil
}
