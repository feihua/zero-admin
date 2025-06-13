package commentreplayservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCommentReplayListLogic 查询产品评价回复列表
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

// QueryCommentReplayList 查询产品评价回复列表
func (l *QueryCommentReplayListLogic) QueryCommentReplayList(in *pmsclient.QueryCommentReplayListReq) (*pmsclient.QueryCommentReplayListResp, error) {
	result, err := l.svcCtx.ProductCommentReplayModel.FindPage(l.ctx, in.CommentId, in.PageNum, in.PageSize)

	if err != nil {
		logc.Errorf(l.ctx, "查询产品评价回复列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询产品评价回复列表失败")
	}

	var list []*pmsclient.CommentReplayListData
	for _, item := range result {

		list = append(list, &pmsclient.CommentReplayListData{
			Id:             item.ID.Hex(),                      //
			CommentId:      item.CommentId,                     // 评论id
			MemberNickName: item.MemberNickName,                // 评论人员昵称
			MemberIcon:     item.MemberIcon,                    // 评论人员头像
			Content:        item.Content,                       // 内容
			CreateTime:     time_util.TimeToStr(item.CreateAt), // 评论时间
			Type:           item.Type,                          // 评论人员类型；0->会员；1->管理员
		})
	}

	return &pmsclient.QueryCommentReplayListResp{
		Total: 0,
		List:  list,
	}, nil

}
