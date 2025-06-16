package productspecservicelogic

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

// QueryProductSpecListLogic 查询商品规格列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSpecListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductSpecListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpecListLogic {
	return &QueryProductSpecListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductSpecList 查询商品规格列表
func (l *QueryProductSpecListLogic) QueryProductSpecList(in *pmsclient.QueryProductSpecListReq) (*pmsclient.QueryProductSpecListResp, error) {
	productSpec := query.PmsProductSpec
	q := productSpec.WithContext(l.ctx)
	if in.CategoryId != 0 {
		q = q.Where(productSpec.CategoryID.Eq(in.CategoryId))
	}
	if len(in.Name) > 0 {
		q = q.Where(productSpec.Name.Like("%" + in.Name + "%"))
	}
	if in.Status != 2 {
		q = q.Where(productSpec.Status.Eq(in.Status))
	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品规格列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品规格列表失败")
	}

	var list []*pmsclient.ProductSpecListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductSpecListData{
			Id:         item.ID,                                          //
			CategoryId: item.CategoryID,                                  // 分类ID
			Name:       item.Name,                                        // 规格名称
			Sort:       item.Sort,                                        // 排序
			Status:     item.Status,                                      // 状态：0->禁用；1->启用
			CreateBy:   item.CreateBy,                                    // 创建人ID
			CreateTime: time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:   pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime: time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &pmsclient.QueryProductSpecListResp{
		Total: count,
		List:  list,
	}, nil
}
