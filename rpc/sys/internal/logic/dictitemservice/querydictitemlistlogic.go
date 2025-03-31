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

	if len(in.DictType) > 0 {
		q = q.Where(query.SysDictItem.DictType.Eq(in.DictType))
	}
	if len(in.DictLabel) > 0 {
		q = q.Where(query.SysDictItem.DictLabel.Like("%" + in.DictLabel + "%"))
	}

	if in.Status != 2 {
		q = q.Where(query.SysDictItem.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询字典数据列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询字典数据列表失败")
	}

	var list = make([]*sysclient.DictItemListData, 0, len(result))

	for _, item := range result {
		list = append(list, &sysclient.DictItemListData{
			Id:         item.ID,                                 // 字典数据id
			DictSort:   item.DictSort,                           // 字典排序
			DictLabel:  item.DictLabel,                          // 字典标签
			DictValue:  item.DictValue,                          // 字典键值
			DictType:   item.DictType,                           // 字典类型
			CssClass:   item.CSSClass,                           // 样式属性（其他样式扩展）
			ListClass:  item.ListClass,                          // 表格回显样式
			IsDefault:  item.IsDefault,                          // 是否默认（Y是 N否）
			Status:     item.Status,                             // 状态（0：停用，1:正常）
			Remark:     item.Remark,                             // 备注
			CreateBy:   item.CreateBy,                           // 创建者
			CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	return &sysclient.QueryDictItemListResp{
		Total: count,
		List:  list,
	}, nil
}
