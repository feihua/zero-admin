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

// QueryDictItemListLogic 查询字典数据列表
/*
Author: LiuFeiHua
Date: 2024/5/28 17:03
*/
type QueryDictItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemListLogic {
	return &QueryDictItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDictItemList 查询字典数据列表
func (l *QueryDictItemListLogic) QueryDictItemList(in *sysclient.QueryDictItemListReq) (*sysclient.QueryDictItemListResp, error) {
	q := query.SysDictItem.WithContext(l.ctx)

	q = q.Where(query.SysDictItem.DictType.Eq(in.DictType))
	if len(in.DictLabel) > 0 {
		q = q.Where(query.SysDictItem.DictLabel.Like("%" + in.DictLabel + "%"))
	}

	if in.DictStatus != 2 {
		q = q.Where(query.SysDictItem.DictStatus.Eq(in.DictStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询字典数据列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询字典数据列表信息失败")
	}

	var list []*sysclient.DictItemListData

	for _, dict := range result {
		list = append(list, &sysclient.DictItemListData{
			Id:         dict.ID,                                       // 编号
			DictType:   dict.DictType,                                 // 字典类型
			DictLabel:  dict.DictLabel,                                // 字典标签
			DictValue:  dict.DictValue,                                // 字典键值
			DictStatus: dict.DictStatus,                               // 字典状态
			DictSort:   dict.DictSort,                                 // 排序
			Remark:     dict.Remark,                                   // 备注信息
			IsDefault:  dict.IsDefault,                                // 是否默认  0：否  1：是
			IsDeleted:  dict.IsDeleted,                                // 是否删除  0：否  1：是
			CreateBy:   dict.CreateBy,                                 // 创建者
			CreateTime: dict.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间
			UpdateBy:   dict.UpdateBy,                                 // 更新者
			UpdateTime: time_util.TimeToString(dict.UpdateTime),       // 更新时间
		})
	}

	return &sysclient.QueryDictItemListResp{
		Total: count,
		List:  list,
	}, nil
}
