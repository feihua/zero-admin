package commentreplayservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCommentReplayListLogic 查询产品评价回复表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:34
*/
type QueryCommentReplayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCommentReplayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCommentReplayListLogic {
	return &QueryCommentReplayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCommentReplayList 查询产品评价回复表列表
func (l *QueryCommentReplayListLogic) QueryCommentReplayList(in *pmsclient.QueryCommentReplayListReq) (*pmsclient.QueryCommentReplayListResp, error) {
	q := query.PmsCommentReplay.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询评价回复列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.CommentReplayListData
	for _, item := range result {

		list = append(list, &pmsclient.CommentReplayListData{
			Id:             item.ID,
			CommentId:      item.CommentID,
			MemberNickName: item.MemberNickName,
			MemberIcon:     item.MemberIcon,
			Content:        item.Content,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			Type:           item.Type,
		})
	}

	return &pmsclient.QueryCommentReplayListResp{
		Total: count,
		List:  list,
	}, nil

}
