package dictitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
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

// QueryDictItemDetail 查询字典数据表详情
func (l *QueryDictItemDetailLogic) QueryDictItemDetail(in *sysclient.QueryDictItemDetailReq) (*sysclient.QueryDictItemDetailResp, error) {
	dictItem, err := query.SysDictItem.WithContext(l.ctx).Where(query.SysDictItem.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询字典数据表详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询字典数据表详情失败")
	}

	data := &sysclient.QueryDictItemDetailResp{
		CreateBy:   dictItem.CreateBy,
		CreateTime: dictItem.CreateTime.Format("2006-01-02 15:04:05"),
		DictLabel:  dictItem.DictLabel,
		DictSort:   dictItem.DictSort,
		DictStatus: dictItem.DictStatus,
		DictType:   dictItem.DictType,
		DictValue:  dictItem.DictValue,
		Id:         dictItem.ID,
		IsDefault:  dictItem.IsDefault,
		Remark:     dictItem.Remark,
		UpdateBy:   dictItem.UpdateBy,
		UpdateTime: common.TimeToString(dictItem.UpdateTime),
	}

	logc.Infof(l.ctx, "查询字典数据表详情,参数：%+v,响应：%+v", in, data)

	return data, nil

}
