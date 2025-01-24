package commentreplayservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
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
			Id:             item.ID,                              //
			CommentId:      item.CommentID,                       // 评论id
			MemberNickName: item.MemberNickName,                  // 评论人员昵称
			MemberIcon:     item.MemberIcon,                      // 评论人员头像
			Content:        item.Content,                         // 内容
			CreateTime:     time_util.TimeToStr(item.CreateTime), // 评论时间
			Type:           item.Type,                            // 评论人员类型；0->会员；1->管理员
		})
	}

	return &pmsclient.QueryCommentReplayListResp{
		Total: count,
		List:  list,
	}, nil

}
