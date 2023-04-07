package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagListLogic {
	return &MemberTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagListLogic) MemberTagList(in *umsclient.MemberTagListReq) (*umsclient.MemberTagListResp, error) {
	// todo: add your logic here and delete this line
	// all, err := l.svcCtx.UmsMemberTagModel.FindAll(l.ctx, in.Current, in.PageSize)
	// count, _ := l.svcCtx.UmsMemberTagModel.Count(l.ctx)

	// if err != nil {
	// 	reqStr, _ := json.Marshal(in)
	// 	logx.WithContext(l.ctx).Errorf("查询会员标签列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
	// 	return nil, err
	// }

	// var list []*ums.MemberTagListData
	// for _, item := range *all {

	// 	list = append(list, &ums.MemberTagListData{
	// 		Id:                item.Id,
	// 		Name:              item.Name,
	// 		FinishOrderCount:  item.FinishOrderCount,
	// 		FinishOrderAmount: int64(item.FinishOrderAmount),
	// 	})
	// }

	// reqStr, _ := json.Marshal(in)
	// listStr, _ := json.Marshal(list)
	// logx.WithContext(l.ctx).Infof("查询会员标签列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.MemberTagListResp{
		// Total: count,
		// List:  list,
	}, nil
}
