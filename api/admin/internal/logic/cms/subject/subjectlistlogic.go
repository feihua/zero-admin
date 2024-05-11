package subject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectListLogic {
	return &SubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectListLogic) SubjectList(req *types.ListSubjectReq) (resp *types.ListSubjectResp, err error) {
	subjectList, err := l.svcCtx.SubjectService.SubjectList(l.ctx, &cmsclient.SubjectListReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		Title:           req.Title,
		RecommendStatus: req.RecommendStatus,
		ShowStatus:      req.ShowStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询专题列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询专题失败")
	}

	var list []*types.ListtSubjectData

	for _, item := range subjectList.List {
		list = append(list, &types.ListtSubjectData{
			Id:              item.Id,
			CategoryId:      item.CategoryId,
			Title:           item.Title,
			Pic:             item.Pic,
			ProductCount:    item.ProductCount,
			RecommendStatus: item.RecommendStatus,
			CreateTime:      item.CreateTime,
			CollectCount:    item.CollectCount,
			ReadCount:       item.ReadCount,
			CommentCount:    item.CommentCount,
			AlbumPics:       item.AlbumPics,
			Description:     item.Description,
			ShowStatus:      item.ShowStatus,
			Content:         item.Content,
			ForwardCount:    item.ForwardCount,
			CategoryName:    item.CategoryName,
		})
	}

	return &types.ListSubjectResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    subjectList.Total,
		Code:     "000000",
		Message:  "查询专题成功",
	}, nil
}
