package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

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

func (l *MemberTaskListLogic) MemberTaskList(in *ums.MemberTaskListReq) (*ums.MemberTaskListResp, error) {
	all, err := l.svcCtx.UmsMemberTaskModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberTaskModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员任务列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*ums.MemberTaskListData
	for _, item := range *all {

		list = append(list, &ums.MemberTaskListData{
			Id:           item.Id,
			Name:         item.Name,
			Growth:       item.Growth,
			Intergration: item.Intergration,
			Type:         item.Type,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员任务列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &ums.MemberTaskListResp{
		Total: count,
		List:  list,
	}, nil

}
