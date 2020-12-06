package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryListLogic {
	return OperateHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryListLogic) OperateHistoryList(req types.ListOperateHistoryReq) (*types.ListOperateHistoryResp, error) {
	resp, err := l.svcCtx.Oms.OrderOperateHistoryList(l.ctx, &omsclient.OrderOperateHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	//var list []*types.ListUserData

	//for _, user := range resp.List {
	//	list = append(list, &types.ListUserData{
	//		Id:             user.Id,
	//		Name:           user.Name,
	//		NickName:       user.NickName,
	//		Password:       user.Password,
	//		Salt:           user.Salt,
	//		Email:          user.Email,
	//		Mobile:         user.Mobile,
	//		DeptId:         user.DeptId,
	//		CreateBy:       user.CreateBy,
	//		CreateTime:     user.CreateTime,
	//		LastUpdateBy:   user.LastUpdateBy,
	//		LastUpdateTime: user.LastUpdateTime,
	//		DelFlag:        user.DelFlag,
	//	})
	//}

	return &types.ListOperateHistoryResp{
		Current:  req.Current,
		Data:     nil,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
