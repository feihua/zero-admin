package logic

import (
	"context"

	"zero-admin/common/xerr"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/test/model"
)

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

type MemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberInfoLogic {
	return &MemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberInfoLogic) MemberInfo(in *umsclient.MemberInfoReq) (*umsclient.MemberInfoResp, error) {
	member, err := l.svcCtx.UmsMemberModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", in.Id, err)
	}
	if member == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.Id)
	}
	// var respMember umsclient.Member
	// _ = copier.Copy(&respMember, member)
	// return &umsclient.MemberInfoResp{
	// 	Member: &respMember,
	// }, nil

	return &umsclient.MemberInfoResp{
		Member: &umsclient.Member{
			Id:             member.Id,
			Phone:          member.Phone,
			Nickname:       member.Nickname,
			Icon:           member.Icon,
			LevelId:        member.MemberLevelId,
			ParentId:       member.ParentId,
			Balance:        member.Balance,
			InvitationCode: member.InvitationCode,
		},
	}, nil
}
