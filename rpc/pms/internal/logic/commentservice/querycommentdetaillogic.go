package commentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCommentDetailLogic 查询商品评价详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:05
*/
type QueryCommentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCommentDetailLogic {
	return &QueryCommentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCommentDetail 查询商品评价详情
func (l *QueryCommentDetailLogic) QueryCommentDetail(in *pmsclient.QueryCommentDetailReq) (*pmsclient.QueryCommentDetailResp, error) {
	item, err := l.svcCtx.ProductCommentModel.FindOne(l.ctx, in.Id)

	if err != nil {
		logc.Errorf(l.ctx, "查询商品评价详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品评价详情失败")
	}

	data := &pmsclient.QueryCommentDetailResp{
		Id:               item.ID.Hex(),                      //
		ProductId:        item.ProductId,                     // 商品id
		MemberNickName:   item.MemberNickName,                // 评价者昵称
		ProductName:      item.ProductName,                   // 商品名称
		Star:             item.Star,                          // 评价星数：0->5
		MemberIp:         item.MemberIp,                      // 评价的ip
		CreateTime:       time_util.TimeToStr(item.CreateAt), // 评价时间
		ShowStatus:       item.ShowStatus,                    // 是否显示，0-不显示，1-显示
		ProductAttribute: item.ProductAttribute,              // 购买时的商品属性
		CollectCount:     item.CollectCount,                  // 点赞数
		ReadCount:        item.ReadCount,                     // 阅读数
		Content:          item.Content,                       // 内容
		Pics:             item.Pics,                          // 上传图片地址，以逗号隔开
		MemberIcon:       item.MemberIcon,                    // 评论用户头像
		ReplayCount:      item.ReplayCount,                   // 回复数量
	}

	return data, nil
}
