package memberservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberUpdateLogic {
	return &MemberUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberUpdateLogic) MemberUpdate(in *umsclient.MemberUpdateReq) (*umsclient.MemberUpdateResp, error) {
	createTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	birthday, _ := time.Parse("2006-01-02 15:04:05", in.Birthday)
	l.svcCtx.UmsMemberModel.Update(l.ctx, &umsmodel.UmsMember{
		Id:                    in.Id,
		MemberLevelId:         in.MemberLevelId,
		Username:              in.Username,
		Password:              in.Password,
		Nickname:              in.Nickname,
		Phone:                 in.Phone,
		Status:                in.Status,
		CreateTime:            createTime,
		Icon:                  sql.NullString{String: in.Icon, Valid: true},
		Gender:                sql.NullInt64{Int64: in.Gender, Valid: true},
		Birthday:              sql.NullTime{Time: birthday, Valid: true},
		City:                  sql.NullString{String: in.City, Valid: true},
		Job:                   sql.NullString{String: in.Job, Valid: true},
		PersonalizedSignature: sql.NullString{String: in.PersonalizedSignature, Valid: true},
		SourceType:            in.SourceType,
		Integration:           in.Integration,
		Growth:                in.Growth,
		LuckeyCount:           in.LuckeyCount,
		HistoryIntegration:    in.HistoryIntegration,
	})

	return &umsclient.MemberUpdateResp{}, nil
}
