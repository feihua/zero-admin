package dicttypeservicelogic

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

// QueryDictTypeDetailLogic 查询字典类型详情
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

// QueryDictTypeDetail 查询字典类型详情
// 1.判断字典类型是否存在
func (l *QueryDictTypeDetailLogic) QueryDictTypeDetail(in *sysclient.QueryDictTypeDetailReq) (*sysclient.QueryDictTypeDetailResp, error) {
	dict, err := query.SysDictType.WithContext(l.ctx).Where(query.SysDictType.ID.Eq(in.Id)).First()

	// 1.判断字典类型是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "字典类型不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("字典类型不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询字典类型异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询字典类型异常")
	}

	data := &sysclient.QueryDictTypeDetailResp{
		Id:         dict.ID,                                 // 编号
		DictName:   dict.DictName,                           // 字典名称
		DictType:   dict.DictType,                           // 字典类型
		DictStatus: dict.DictStatus,                         // 字典状态
		Remark:     dict.Remark,                             // 备注信息
		IsSystem:   dict.IsSystem,                           // 是否系统预留  0：否  1：是
		IsDeleted:  dict.IsDeleted,                          // 是否删除  0：否  1：是
		CreateBy:   dict.CreateBy,                           // 创建者
		CreateTime: time_util.TimeToStr(dict.CreateTime),    // 创建时间
		UpdateBy:   dict.UpdateBy,                           // 更新者
		UpdateTime: time_util.TimeToString(dict.UpdateTime), // 更新时间
	}

	return data, nil
}
