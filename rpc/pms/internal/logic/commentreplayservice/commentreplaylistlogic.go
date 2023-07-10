package commentreplayservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentReplayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayListLogic {
	return &CommentReplayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentReplayListLogic) CommentReplayList(in *pmsclient.CommentReplayListReq) (*pmsclient.CommentReplayListResp, error) {
	all, err := l.svcCtx.PmsCommentReplayModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsCommentReplayModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询评价回复列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pmsclient.CommentReplayListData
	for _, item := range *all {

		list = append(list, &pmsclient.CommentReplayListData{
			Id:             item.Id,
			CommentId:      item.CommentId,
			MemberNickName: item.MemberNickName,
			MemberIcon:     item.MemberIcon,
			Content:        item.Content,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			Type:           item.Type,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询评价回复列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pmsclient.CommentReplayListResp{
		Total: count,
		List:  list,
	}, nil
}
