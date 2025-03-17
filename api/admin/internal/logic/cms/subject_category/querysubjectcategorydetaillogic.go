package subject_category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySubjectCategoryDetailLogic 查询专题分类表详情
/*
Author: 刘飞华
Date: 2025/02/07 10:21:09
*/
type QuerySubjectCategoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubjectCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCategoryDetailLogic {
	return &QuerySubjectCategoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySubjectCategoryDetail 查询专题分类表详情
func (l *QuerySubjectCategoryDetailLogic) QuerySubjectCategoryDetail(req *types.QuerySubjectCategoryDetailReq) (resp *types.QuerySubjectCategoryDetailResp, err error) {

	detail, err := l.svcCtx.SubjectCategoryService.QuerySubjectCategoryDetail(l.ctx, &cmsclient.QuerySubjectCategoryDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询专题分类表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QuerySubjectCategoryDetailData{
		Id:           detail.Id,           // 主键ID
		Name:         detail.Name,         // 专题分类名称
		Icon:         detail.Icon,         // 分类图标
		SubjectCount: detail.SubjectCount, // 专题数量
		ShowStatus:   detail.ShowStatus,   // 显示状态：0->不显示；1->显示
		Sort:         detail.Sort,         // 排序
		CreateBy:     detail.CreateBy,     // 创建者
		CreateTime:   detail.CreateTime,   // 创建时间
		UpdateBy:     detail.UpdateBy,     // 更新者
		UpdateTime:   detail.UpdateTime,   // 更新时间
	}
	return &types.QuerySubjectCategoryDetailResp{
		Code:    "000000",
		Message: "查询专题分类表成功",
		Data:    data,
	}, nil
}
