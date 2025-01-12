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

	for _, dept := range result {
		list = append(list, &sysclient.DictItemListData{
			CreateBy:   dept.CreateBy,
			CreateTime: dept.CreateTime.Format("2006-01-02 15:04:05"),
			DictLabel:  dept.DictLabel,
			DictSort:   dept.DictSort,
			DictStatus: dept.DictStatus,
			DictType:   dept.DictType,
			DictValue:  dept.DictValue,
			Id:         dept.ID,
			IsDefault:  dept.IsDefault,
			Remark:     dept.Remark,
			UpdateBy:   dept.UpdateBy,
			UpdateTime: time_util.TimeToString(dept.UpdateTime),
		})
	}

	logc.Infof(l.ctx, "查询字典数据列表信息,参数：%+v,响应：%+v", in, list)

	return &sysclient.QueryDictItemListResp{
		Total: count,
		List:  list,
	}, nil
}
