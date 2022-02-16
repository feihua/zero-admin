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

type MemberLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogListLogic {
	return MemberLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogListLogic) MemberLoginLogList(req types.ListMemberLoginLogReq) (*types.ListMemberLoginLogResp, error) {
	resp, err := l.svcCtx.Ums.MemberLoginLogList(l.ctx, &umsclient.MemberLoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询会员登录记录列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询员登录记录列表失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberLoginLogData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberLoginLogData{
			Id:         item.Id,
			MemberId:   item.MemberId,
			CreateTime: item.CreateTime,
			Ip:         item.Ip,
			City:       item.City,
			LoginType:  item.LoginType,
			Province:   item.Province,
		})
	}

	return &types.ListMemberLoginLogResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询员登录记录列表成功",
	}, nil
}
