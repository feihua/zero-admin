package dicttypeservicelogic

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

// QueryDictTypeDetailLogic 查询字典类型表详情
/*
Author: LiuFeiHua
Date: 2024/5/30 10:50
*/
type QueryDictTypeDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictTypeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictTypeDetailLogic {
	return &QueryDictTypeDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDictTypeDetail 查询字典类型表详情
func (l *QueryDictTypeDetailLogic) QueryDictTypeDetail(in *sysclient.QueryDictTypeDetailReq) (*sysclient.QueryDictTypeDetailResp, error) {
	dict, err := query.SysDictType.WithContext(l.ctx).Where(query.SysDictType.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询字典类型表详情失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询字典类型表详情失败")
	}

	data := &sysclient.QueryDictTypeDetailResp{
		CreateBy:   dict.CreateBy,
		CreateTime: dict.CreateTime.Format("2006-01-02 15:04:05"),
		DictName:   dict.DictName,
		DictStatus: dict.DictStatus,
		DictType:   dict.DictType,
		Id:         dict.ID,
		IsSystem:   dict.IsSystem,
		Remark:     dict.Remark,
		UpdateBy:   dict.UpdateBy,
		UpdateTime: common.TimeToString(dict.UpdateTime),
	}

	logc.Infof(l.ctx, "查询字典类型表详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
