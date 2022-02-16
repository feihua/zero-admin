package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagListLogic {
	return MemberTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagListLogic) MemberTagList(req types.ListMemberTagReq) (*types.ListMemberTagResp, error) {
	resp, err := l.svcCtx.Ums.MemberTagList(l.ctx, &umsclient.MemberTagListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询会员标签列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询会员标签失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberTagData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberTagData{
			Id:                item.Id,
			Name:              item.Name,
			FinishOrderCount:  item.FinishOrderCount,
			FinishOrderAmount: float64(item.FinishOrderAmount),
		})
	}

	return &types.ListMemberTagResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员标签成功",
	}, nil
}
