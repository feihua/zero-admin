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
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDictItemLogic 更新字典数据
/*
Author: LiuFeiHua
Date: 2024/5/28 17:03
*/
type UpdateDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemLogic {
	return &UpdateDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDictItem 更新字典数据
// 1.判断字典数据是否存在
// 2.根据字典类型查询字典是否已存在
// 3.查询字典标签是否已存在,如果字典标签已存在,则直接返回
// 4.查询字典键值是否已存在,如果字典键值已存在,则直接返回
// 5.如果更新字典数据是默认,则修改其他选项为非默认状态
// 6.字典数据存在时,则直接更新字典数据
func (l *UpdateDictItemLogic) UpdateDictItem(in *sysclient.UpdateDictItemReq) (*sysclient.UpdateDictItemResp, error) {
	q := query.SysDictItem

	dictType := in.DictType

	// 1.判断字典数据是否存在
	dictItem, err := q.WithContext(l.ctx).Where(query.SysDictItem.ID.Eq(in.Id)).First()

	// 1.判断字典数据是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "字典数据不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("字典数据不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询字典数据异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询字典数据异常")
	}

	// 2.判断字典类型是否存在
	count, err := query.SysDictType.WithContext(l.ctx).Where(query.SysDictType.DictType.Eq(dictType)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "根据字典类型：%s,查询字典信息失败,异常:%s", dictType, err.Error())
		return nil, errors.New(fmt.Sprintf("更新字典数据失败"))
	}

	if count == 0 {
		return nil, errors.New(fmt.Sprintf("更新字典数据失败,字典类型：%s,不存在", dictType))
	}

	// 3.查询字典标签是否已存在,如果字典标签已存在,则直接返回
	count, err = q.WithContext(l.ctx).Where(q.ID.Neq(in.Id), q.DictLabel.Eq(in.DictLabel), q.DictType.Eq(dictType)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典标签：%s,查询字典数据失败,异常:%s", in.DictLabel, err.Error())
		return nil, errors.New(fmt.Sprintf("更新字典数据失败"))
	}

	if count > 0 {
		logc.Errorf(l.ctx, "更新字典数据失败,字典标签已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("更新字典数据失败,字典标签：%s,已存在", in.DictLabel))
	}

	// 4.查询字典键值是否已存在,如果字典键值已存在,则直接返回
	count, err = q.WithContext(l.ctx).Where(q.ID.Neq(in.Id), q.DictValue.Eq(in.DictValue), q.DictType.Eq(dictType)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典键值：%s,查询字典数据失败,异常:%s", in.DictValue, err.Error())
		return nil, errors.New(fmt.Sprintf("更新字典数据失败"))
	}

	if count > 0 {
		logc.Errorf(l.ctx, "更新字典数据失败,字典键值已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("更新字典数据失败,字典键值：%s,已存在", in.DictValue))
	}

	// 5.如果更新字典数据是默认,则修改其他选项为非默认状态
	if in.IsDefault == 1 {
		_, err = q.WithContext(l.ctx).Where(q.ID.Neq(in.Id)).Where(q.DictType.Eq(dictType)).Where(q.IsDefault.Eq(1)).Update(q.IsDefault, 0)
		if err != nil {
			logc.Errorf(l.ctx, "修改字典数据默认状态失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("更新字典数据失败")
		}
	}

	now := time.Now()
	data := &model.SysDictItem{
		ID:         in.Id,               // 编号
		DictType:   dictType,            // 字典类型
		DictLabel:  in.DictLabel,        // 字典标签
		DictValue:  in.DictValue,        // 字典键值
		DictStatus: in.DictStatus,       // 字典状态
		DictSort:   in.DictSort,         // 排序
		Remark:     in.Remark,           // 备注信息
		IsDefault:  in.IsDefault,        // 是否默认  0：否  1：是
		CreateBy:   dictItem.CreateBy,   // 创建者
		CreateTime: dictItem.CreateTime, // 创建时间
		UpdateBy:   in.UpdateBy,         // 更新者
		UpdateTime: &now,                // 更新时间
	}

	// 6.字典数据存在时,则直接更新字典数据
	err = l.svcCtx.DB.Model(&model.SysDictItem{}).WithContext(l.ctx).Where(query.SysDictItem.ID.Eq(in.Id)).Save(data).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新字典数据失败,参数:%+v,异常:%s", dictItem, err.Error())
		return nil, errors.New(fmt.Sprintf("更新字典数据失败"))
	}

	return &sysclient.UpdateDictItemResp{}, nil
}
