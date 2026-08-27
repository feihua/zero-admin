package subject_comment

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

// QuerySubjectCommentListLogic 查询专题评论列表
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QuerySubjectCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubjectCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCommentListLogic {
	return &QuerySubjectCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySubjectCommentList 查询专题评论列表
func (l *QuerySubjectCommentListLogic) QuerySubjectCommentList(req *types.QuerySubjectCommentListReq) (resp *types.QuerySubjectCommentListResp, err error) {
	result, err := l.svcCtx.SubjectCommentService.QuerySubjectCommentList(l.ctx, &cmsclient.QuerySubjectCommentListReq{
		PageNum:        int32(req.Current),
		PageSize:       int32(req.PageSize),
		SubjectId:      req.SubjectId,      // 关联专题id
		MemberNickName: req.MemberNickName, // 关联会员昵称
		ShowStatus:     req.ShowStatus,     // 是否显示，0->不显示；1->显示

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字专题评论列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QuerySubjectCommentListData

	for _, item := range result.List {
		list = append(list, &types.QuerySubjectCommentListData{
			Id:             item.Id,             // 编号
			SubjectId:      item.SubjectId,      // 关联专题id
			MemberNickName: item.MemberNickName, // 关联会员昵称
			MemberIcon:     item.MemberIcon,     // 会员头像
			Content:        item.Content,        // 评论内容
			CreateTime:     item.CreateTime,     // 创建时间
			ShowStatus:     item.ShowStatus,     // 是否显示，0->不显示；1->显示
		})
	}

	return &types.QuerySubjectCommentListResp{
		Code:     "000000",
		Message:  "查询专题评论列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
