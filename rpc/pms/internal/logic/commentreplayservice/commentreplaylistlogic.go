package commentreplayservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	q := query.PmsCommentReplay.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

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

	logc.Infof(l.ctx, "查询评价回复列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.CommentReplayListResp{
		Total: count,
		List:  list,
	}, nil
}
