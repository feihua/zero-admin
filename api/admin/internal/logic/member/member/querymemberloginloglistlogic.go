package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLoginLogListLogic {
	return &QueryMemberLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMemberLoginLogListLogic) QueryMemberLoginLogList(req *types.ListMemberLoginLogReq) (resp *types.ListMemberLoginLogResp, err error) {
	result, err := l.svcCtx.MemberLoginLogService.MemberLoginLogList(l.ctx, &umsclient.MemberLoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		MemberId: req.MemberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员登录记录列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询员登录记录列表失败")
	}

	var list []*types.ListMemberLoginLogData

	for _, item := range result.List {
		list = append(list, &types.ListMemberLoginLogData{
			Id:         item.Id,
			MemberId:   item.MemberId,
			CreateTime: item.CreateTime,
			Ip:         item.MemberIP,
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
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询员登录记录列表成功",
	}, nil
}
