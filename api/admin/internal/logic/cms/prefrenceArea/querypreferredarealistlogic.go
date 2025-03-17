package prefrenceArea

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryPreferredAreaListLogic 查询优选专区列表
/*
Author: 刘飞华
Date: 2025/02/04 14:56:41
*/
type QueryPreferredAreaListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPreferredAreaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPreferredAreaListLogic {
	return &QueryPreferredAreaListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryPreferredAreaList 查询优选专区列表
func (l *QueryPreferredAreaListLogic) QueryPreferredAreaList(req *types.QueryPreferredAreaListReq) (resp *types.QueryPreferredAreaListResp, err error) {
	result, err := l.svcCtx.PreferredAreaService.QueryPreferredAreaList(l.ctx, &cmsclient.QueryPreferredAreaListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		Name:       strings.TrimSpace(req.Name),     // 专区名称
		SubTitle:   strings.TrimSpace(req.SubTitle), // 子标题
		ShowStatus: req.ShowStatus,                  // 显示状态：0->不显示；1->显示
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品优选列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品优选失败")
	}

	var list []*types.QueryPreferredAreaListData

	for _, item := range result.List {
		list = append(list, &types.QueryPreferredAreaListData{
			Id:         item.Id,         // 主键ID
			Name:       item.Name,       // 专区名称
			SubTitle:   item.SubTitle,   // 子标题
			Pic:        item.Pic,        // 展示图片
			Sort:       item.Sort,       // 排序
			ShowStatus: item.ShowStatus, // 显示状态：0->不显示；1->显示
			CreateBy:   item.CreateBy,   // 创建者
			CreateTime: item.CreateTime, // 创建时间
			UpdateBy:   item.UpdateBy,   // 更新者
			UpdateTime: item.UpdateTime, // 更新时间
		})
	}

	return &types.QueryPreferredAreaListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询商品优选成功",
	}, nil
}
