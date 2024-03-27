package memberservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberByIdLogic
/*
Author: LiuFeiHua
Date: 2023/11/28 15:07
*/
type QueryMemberByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberByIdLogic {
	return &QueryMemberByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberById 获取个人信息
func (l *QueryMemberByIdLogic) QueryMemberById(in *umsclient.MemberByIdReq) (*umsclient.MemberListData, error) {
	member, err := l.svcCtx.UmsMemberModel.FindOne(l.ctx, in.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "根据会员id查询会员信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &umsclient.MemberListData{
		Id:                    member.Id,
		MemberLevelId:         member.MemberLevelId,
		Username:              member.Username,
		Password:              member.Password,
		Nickname:              member.Nickname,
		Phone:                 member.Phone,
		Status:                member.Status,
		CreateTime:            member.CreateTime.Format("2006-01-02 15:04:05"),
		Icon:                  member.Icon.String,
		Gender:                member.Gender.Int64,
		Birthday:              member.Birthday.Time.Format("2006-01-02 15:04:05"),
		City:                  member.City.String,
		Job:                   member.Job.String,
		PersonalizedSignature: member.PersonalizedSignature.String,
		SourceType:            member.SourceType,
		Integration:           member.Integration,
		Growth:                member.Growth,
		LuckeyCount:           member.LuckeyCount,
		HistoryIntegration:    member.HistoryIntegration,
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logc.Infof(l.ctx, "根据会员id查询会员信息,参数：%s,响应：%s", reqStr, listStr)

	return resp, nil
}
