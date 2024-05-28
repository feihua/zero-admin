package dictservicelogic

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

// QueryDictListLogic 查询字典列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
type QueryDictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictListLogic {
	return &QueryDictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDictList 查询字典表列表
func (l *QueryDictListLogic) QueryDictList(in *sysclient.DictListReq) (*sysclient.DictListResp, error) {
	q := query.SysDict.WithContext(l.ctx)
	if len(in.DictName) > 0 {
		q = q.Where(query.SysDict.DictName.Like("%" + in.DictName + "%"))
	}
	if len(in.DictType) > 0 {
		q = q.Where(query.SysDict.DictType.Like("%" + in.DictType + "%"))
	}

	if in.DictStatus != 2 {
		q = q.Where(query.SysDict.DictStatus.Eq(in.DictStatus))
	}
	if in.IsSystem != 2 {
		q = q.Where(query.SysDict.IsSystem.Eq(in.IsSystem))
	}

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询字典列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询字典列表信息失败")
	}
	var list []*sysclient.DictListData
	for _, dict := range result {
		list = append(list, &sysclient.DictListData{
			CreateBy:   dict.CreateBy,
			CreateTime: dict.CreateTime.Format("2006-01-02 15:04:05"),
			DelFlag:    dict.DelFlag,
			DictName:   dict.DictName,
			DictStatus: dict.DictStatus,
			DictType:   dict.DictType,
			Id:         dict.ID,
			IsSystem:   dict.IsSystem,
			Remark:     dict.Remark,
			UpdateBy:   dict.UpdateBy,
			UpdateTime: common.TimeToString(dict.UpdateTime),
		})
	}

	logc.Infof(l.ctx, "查询字典列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.DictListResp{
		Total: count,
		List:  list,
	}, nil
}
