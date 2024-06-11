package subject

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

type QuerySubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectListLogic {
	return &QuerySubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySubjectListLogic) QuerySubjectList(req *types.QuerySubjectListReq) (resp *types.QuerySubjectListResp, err error) {
	subjectList, err := l.svcCtx.SubjectService.QuerySubjectList(l.ctx, &cmsclient.QuerySubjectListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		Title:           strings.TrimSpace(req.Title),
		RecommendStatus: req.RecommendStatus,
		ShowStatus:      req.ShowStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询专题列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询专题失败")
	}

	var list []*types.QuerySubjectListData

	for _, item := range subjectList.List {
		list = append(list, &types.QuerySubjectListData{
			Id:              item.Id,
			CategoryId:      item.CategoryId,
			Title:           item.Title,
			Pic:             item.Pic,
			ProductCount:    item.ProductCount,
			RecommendStatus: item.RecommendStatus,
			CollectCount:    item.CollectCount,
			ReadCount:       item.ReadCount,
			CommentCount:    item.CommentCount,
			AlbumPics:       item.AlbumPics,
			Description:     item.Description,
			ShowStatus:      item.ShowStatus,
			Content:         item.Content,
			ForwardCount:    item.ForwardCount,
			CategoryName:    item.CategoryName,
			CreateBy:        item.CreateBy,
			CreateTime:      item.CreateTime,
			UpdateBy:        item.UpdateBy,
			UpdateTime:      item.UpdateTime,
		})
	}

	return &types.QuerySubjectListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    subjectList.Total,
		Code:     "000000",
		Message:  "查询专题成功",
	}, nil
}
