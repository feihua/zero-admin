package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskListLogic {
	return &MemberTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskListLogic) MemberTaskList(in *umsclient.MemberTaskListReq) (*umsclient.MemberTaskListResp, error) {
	q := query.UmsMemberTask.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员任务列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*umsclient.MemberTaskListData
	for _, item := range result {

		list = append(list, &umsclient.MemberTaskListData{
			Id:           item.ID,
			Name:         item.Name,
			Growth:       item.Growth,
			Intergration: item.Intergration,
			Type:         item.Type,
		})
	}

	logc.Infof(l.ctx, "查询会员任务列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberTaskListResp{
		Total: count,
		List:  list,
	}, nil

}
