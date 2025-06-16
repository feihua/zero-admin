package productspecvalueservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductSpecValueListLogic 查询商品规格值列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSpecValueListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductSpecValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpecValueListLogic {
	return &QueryProductSpecValueListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductSpecValueList 查询商品规格值列表
func (l *QueryProductSpecValueListLogic) QueryProductSpecValueList(in *pmsclient.QueryProductSpecValueListReq) (*pmsclient.QueryProductSpecValueListResp, error) {
	productSpecValue := query.PmsProductSpecValue
	q := productSpecValue.WithContext(l.ctx)
	if in.SpecId != 0 {
		q = q.Where(productSpecValue.SpecID.Eq(in.SpecId))
	}
	if in.Status != 2 {
		q = q.Where(productSpecValue.Status.Eq(in.Status))
	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品规格值列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品规格值列表失败")
	}

	var list []*pmsclient.ProductSpecValueListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductSpecValueListData{
			Id:         item.ID,                                          //
			SpecId:     item.SpecID,                                      // 规格ID
			Value:      item.Value,                                       // 规格值
			Sort:       item.Sort,                                        // 排序
			Status:     item.Status,                                      // 状态：0->禁用；1->启用
			CreateBy:   item.CreateBy,                                    // 创建人ID
			CreateTime: time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:   pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime: time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &pmsclient.QueryProductSpecValueListResp{
		Total: count,
		List:  list,
	}, nil
}
