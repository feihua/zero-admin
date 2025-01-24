package commentreplayservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCommentReplayDetailLogic 查询产品评价回复详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:05
*/
type QueryCommentReplayDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCommentReplayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCommentReplayDetailLogic {
	return &QueryCommentReplayDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCommentReplayDetail 查询产品评价回复详情
func (l *QueryCommentReplayDetailLogic) QueryCommentReplayDetail(in *pmsclient.QueryCommentReplayDetailReq) (*pmsclient.QueryCommentReplayDetailResp, error) {
	item, err := query.PmsCommentReplay.WithContext(l.ctx).Where(query.PmsCommentReplay.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品评价回复详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询产品评价回复详情失败")
	}

	data := &pmsclient.QueryCommentReplayDetailResp{
		Id:             item.ID,                              //
		CommentId:      item.CommentID,                       // 评论id
		MemberNickName: item.MemberNickName,                  // 评论人员昵称
		MemberIcon:     item.MemberIcon,                      // 评论人员头像
		Content:        item.Content,                         // 内容
		CreateTime:     time_util.TimeToStr(item.CreateTime), // 评论时间
		Type:           item.Type,                            // 评论人员类型；0->会员；1->管理员
	}

	logc.Infof(l.ctx, "查询产品评价回复详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
