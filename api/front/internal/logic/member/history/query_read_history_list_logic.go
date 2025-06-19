package history

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryReadHistoryListLogic 查询会员浏览商品的记录
/*
Author: LiuFeiHua
Date: 2025/6/19 11:01
*/
type QueryReadHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryReadHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryReadHistoryListLogic {
	return &QueryReadHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryReadHistoryList 查询会员浏览商品的记录
func (l *QueryReadHistoryListLogic) QueryReadHistoryList() (resp *types.ReadHistoryListResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	historyList, err := l.svcCtx.MemberReadHistoryService.QueryMemberReadHistoryList(l.ctx, &umsclient.QueryMemberReadHistoryListReq{
		PageNum:  1,
		PageSize: 100,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员浏览商品的记录失败,参数memberId: %d,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []types.ReadHistoryList

	for _, detail := range historyList.List {
		list = append(list, types.ReadHistoryList{
			Id:              detail.Id,              //
			MemberId:        detail.MemberId,        // 会员id
			MemberNickName:  detail.MemberNickName,  // 会员姓名
			MemberIcon:      detail.MemberIcon,      // 会员头像
			ProductId:       detail.ProductId,       // 商品id
			ProductName:     detail.ProductName,     // 商品名称
			ProductPic:      detail.ProductPic,      // 商品图片
			ProductSubTitle: detail.ProductSubTitle, // 商品标题
			ProductPrice:    detail.ProductPrice,    // 商品价格
			CreateTime:      detail.CreateTime,      // 浏览时间
		})
	}

	return &types.ReadHistoryListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
