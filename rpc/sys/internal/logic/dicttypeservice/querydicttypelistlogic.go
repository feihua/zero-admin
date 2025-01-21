package dicttypeservicelogic

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

// QueryDictTypeListLogic 查询字典列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
type QueryDictTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictTypeListLogic {
	return &QueryDictTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDictTypeList 查询字典表列表
func (l *QueryDictTypeListLogic) QueryDictTypeList(in *sysclient.QueryDictTypeListReq) (*sysclient.QueryDictTypeListResp, error) {
	q := query.SysDictType.WithContext(l.ctx)
	if len(in.DictName) > 0 {
		q = q.Where(query.SysDictType.DictName.Like("%" + in.DictName + "%"))
	}
	if len(in.DictType) > 0 {
		q = q.Where(query.SysDictType.DictType.Like("%" + in.DictType + "%"))
	}

	if in.DictStatus != 2 {
		q = q.Where(query.SysDictType.DictStatus.Eq(in.DictStatus))
	}
	if in.IsSystem != 2 {
		q = q.Where(query.SysDictType.IsSystem.Eq(in.IsSystem))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询字典列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询字典列表信息失败")
	}
	var list []*sysclient.DictTypeListData
	for _, dict := range result {
		createTime := dict.CreateTime.Format("2006-01-02 15:04:05")
		list = append(list, &sysclient.DictTypeListData{
			Id:         dict.ID,                                 // 编号
			DictName:   dict.DictName,                           // 字典名称
			DictType:   dict.DictType,                           // 字典类型
			DictStatus: dict.DictStatus,                         // 字典状态
			Remark:     dict.Remark,                             // 备注信息
			IsSystem:   dict.IsSystem,                           // 是否系统预留  0：否  1：是
			IsDeleted:  dict.IsDeleted,                          // 是否删除  0：否  1：是
			CreateBy:   dict.CreateBy,                           // 创建者
			CreateTime: createTime,                              // 创建时间
			UpdateBy:   dict.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(dict.UpdateTime), // 更新时间
		})
	}

	logc.Infof(l.ctx, "查询字典列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.QueryDictTypeListResp{
		Total: count,
		List:  list,
	}, nil
}
