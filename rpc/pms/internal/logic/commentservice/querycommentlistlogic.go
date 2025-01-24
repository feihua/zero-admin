package commentservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCommentListLogic 查询商品评价表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:36
*/
type QueryCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCommentListLogic {
	return &QueryCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCommentList 查询商品评价表列表
func (l *QueryCommentListLogic) QueryCommentList(in *pmsclient.QueryCommentListReq) (*pmsclient.QueryCommentListResp, error) {
	q := query.PmsComment.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品评价列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.CommentListData
	for _, item := range result {

		list = append(list, &pmsclient.CommentListData{
			Id:               item.ID,                              //
			ProductId:        item.ProductID,                       // 商品id
			MemberNickName:   item.MemberNickName,                  // 评价者昵称
			ProductName:      item.ProductName,                     // 商品名称
			Star:             item.Star,                            // 评价星数：0->5
			MemberIp:         item.MemberIP,                        // 评价的ip
			CreateTime:       time_util.TimeToStr(item.CreateTime), // 评价时间
			ShowStatus:       item.ShowStatus,                      // 是否显示，0-不显示，1-显示
			ProductAttribute: item.ProductAttribute,                // 购买时的商品属性
			CollectCount:     item.CollectCount,                    // 点赞数
			ReadCount:        item.ReadCount,                       // 阅读数
			Content:          item.Content,                         // 内容
			Pics:             item.Pics,                            // 上传图片地址，以逗号隔开
			MemberIcon:       item.MemberIcon,                      // 评论用户头像
			ReplayCount:      item.ReplayCount,                     // 回复数量
		})
	}

	return &pmsclient.QueryCommentListResp{
		Total: count,
		List:  list,
	}, nil

}
