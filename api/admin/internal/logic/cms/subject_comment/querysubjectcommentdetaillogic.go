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

// QuerySubjectCommentDetailLogic 查询专题评论详情
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QuerySubjectCommentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubjectCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCommentDetailLogic {
	return &QuerySubjectCommentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySubjectCommentDetail 查询专题评论详情
func (l *QuerySubjectCommentDetailLogic) QuerySubjectCommentDetail(req *types.QuerySubjectCommentDetailReq) (resp *types.QuerySubjectCommentDetailResp, err error) {

	detail, err := l.svcCtx.SubjectCommentService.QuerySubjectCommentDetail(l.ctx, &cmsclient.QuerySubjectCommentDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询专题评论详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QuerySubjectCommentDetailData{
		Id:             detail.Id,             // 编号
		SubjectId:      detail.SubjectId,      // 关联专题id
		MemberNickName: detail.MemberNickName, // 关联会员昵称
		MemberIcon:     detail.MemberIcon,     // 会员头像
		Content:        detail.Content,        // 评论内容
		CreateTime:     detail.CreateTime,     // 创建时间
		ShowStatus:     detail.ShowStatus,     // 是否显示，0->不显示；1->显示
	}
	return &types.QuerySubjectCommentDetailResp{
		Code:    "000000",
		Message: "查询专题评论成功",
		Data:    data,
	}, nil
}
