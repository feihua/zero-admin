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

	if in.Status != 2 {
		q = q.Where(query.SysDictType.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询字典列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询字典列表信息失败")
	}
	var list = make([]*sysclient.DictTypeListData, 0, len(result))
	for _, dict := range result {
		list = append(list, &sysclient.DictTypeListData{
			Id:         dict.ID,                                 // 字典id
			DictName:   dict.DictName,                           // 字典名称
			DictType:   dict.DictType,                           // 字典类型
			Status:     dict.Status,                             // 状态（0：停用，1:正常）
			Remark:     dict.Remark,                             // 备注
			CreateBy:   dict.CreateBy,                           // 创建者
			CreateTime: time_util.TimeToStr(dict.CreateTime),    // 创建时间
			UpdateBy:   dict.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(dict.UpdateTime), // 更新时间
		})
	}

	return &sysclient.QueryDictTypeListResp{
		Total: count,
		List:  list,
	}, nil
}
