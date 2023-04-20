package logic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardRelationListLogic {
	return &MemberPrepaidCardRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardRelationListLogic) MemberPrepaidCardRelationList(in *umsclient.MemberPrepaidCardRelationListReq) (*umsclient.MemberPrepaidCardRelationListResp, error) {
	all, err := l.svcCtx.UmsMemberPrepaidCardRelationModel.FindListByMemberId(l.ctx, in.MemberId, in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberPrepaidCardRelationModel.CountByMemberId(l.ctx, in.MemberId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员地址列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*ums.MemberPrepaidCardRelationListData
	for _, item := range *all {

		list = append(list, &ums.MemberPrepaidCardRelationListData{
			Id:            item.Id,
			MemberId:      item.MemberId,
			PrepaidCardId: item.PrepaidCardId,
			PrepaidCardSn: item.PrepaidCardSn,
			Balance:       item.Balance,
			Status:        item.Status,
			CreateTime:    item.CreateTime.Format("2006-01-02 15:04:05"),
			ExpiredTime:   item.ExpiredTime.Format("2006-01-02 15:04:05"),
			UpdateTime:    item.UpdateTime.Format("2006-01-02 15:04:05"),
			Note:          item.Note,
		})
	}
	// reqStr, _ := json.Marshal(in)
	// listStr, _ := json.Marshal(list)
	// logx.WithContext(l.ctx).Infof("查询会员地址列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.MemberPrepaidCardRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
