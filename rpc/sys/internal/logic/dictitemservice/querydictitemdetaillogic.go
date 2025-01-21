package dictitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

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
func (l *QueryDictItemDetailLogic) QueryDictItemDetail(in *sysclient.QueryDictItemDetailReq) (*sysclient.QueryDictItemDetailResp, error) {
	dictItem, err := query.SysDictItem.WithContext(l.ctx).Where(query.SysDictItem.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询字典数据详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询字典数据详情失败")
	}

	if dictItem == nil {
		logc.Errorf(l.ctx, "查询字典数据详情失败,参数：%+v,字典数据不存在", in)
		return nil, errors.New("查询字典数据详情失败,字典数据不存在")
	}

	createTime := dictItem.CreateTime.Format("2006-01-02 15:04:05")
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
		CreateTime: createTime,                                  // 创建时间
		UpdateBy:   dictItem.UpdateBy,                           // 更新者
		UpdateTime: time_util.TimeToString(dictItem.UpdateTime), // 更新时间
	}

	logc.Infof(l.ctx, "查询字典数据详情,参数：%+v,响应：%+v", in, data)

	return data, nil

}
