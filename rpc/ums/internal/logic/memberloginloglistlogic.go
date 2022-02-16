package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogListLogic {
	return &MemberLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogListLogic) MemberLoginLogList(in *ums.MemberLoginLogListReq) (*ums.MemberLoginLogListResp, error) {
	all, err := l.svcCtx.UmsMemberLoginLogModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberLoginLogModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员登录记录列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*ums.MemberLoginLogListData
	for _, item := range *all {
		list = append(list, &ums.MemberLoginLogListData{
			Id:         item.Id,
			MemberId:   item.MemberId,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			Ip:         item.Ip,
			City:       item.City,
			LoginType:  item.LoginType,
			Province:   item.Province,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员登录记录列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &ums.MemberLoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
