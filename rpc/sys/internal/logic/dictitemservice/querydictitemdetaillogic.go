package dictitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryDictItemDetailLogic QueryDictItemDetail
/*
Author: LiuFeiHua
Date: 2024/5/30 10:27
*/
type QueryDictItemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemDetailLogic {
	return &QueryDictItemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDictItemDetail 查询字典数据详情
// 1.判断字典数据是否存在
func (l *QueryDictItemDetailLogic) QueryDictItemDetail(in *sysclient.QueryDictItemDetailReq) (*sysclient.QueryDictItemDetailResp, error) {
	dictItem, err := query.SysDictItem.WithContext(l.ctx).Where(query.SysDictItem.ID.Eq(in.Id)).First()

	// 1.判断字典数据是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "字典数据不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("字典数据不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询字典数据异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询字典数据异常")
	}

	data := &sysclient.QueryDictItemDetailResp{
		Id:         dictItem.ID,                                 // 编号
		DictType:   dictItem.DictType,                           // 字典类型
		DictLabel:  dictItem.DictLabel,                          // 字典标签
		DictValue:  dictItem.DictValue,                          // 字典键值
		DictStatus: dictItem.DictStatus,                         // 字典状态
		DictSort:   dictItem.DictSort,                           // 排序
		Remark:     dictItem.Remark,                             // 备注信息
		IsDefault:  dictItem.IsDefault,                          // 是否默认  0：否  1：是
		IsDeleted:  dictItem.IsDeleted,                          // 是否删除  0：否  1：是
		CreateBy:   dictItem.CreateBy,                           // 创建者
		CreateTime: time_util.TimeToStr(dictItem.CreateTime),    // 创建时间
		UpdateBy:   dictItem.UpdateBy,                           // 更新者
		UpdateTime: time_util.TimeToString(dictItem.UpdateTime), // 更新时间
	}

	return data, nil

}
