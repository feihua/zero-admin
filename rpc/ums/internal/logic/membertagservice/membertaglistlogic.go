package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	q := query.UmsMemberTag.WithContext(l.ctx)
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员标签列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberTagListData
	for _, item := range result {

		list = append(list, &umsclient.MemberTagListData{
			Id:                item.ID,
			Name:              item.Name,
			FinishOrderCount:  item.FinishOrderCount,
			FinishOrderAmount: float32(item.FinishOrderAmount),
		})
	}

	logc.Infof(l.ctx, "查询会员标签列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberTagListResp{
		Total: count,
		List:  list,
	}, nil

}
