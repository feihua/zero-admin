package logic

import (
	"context"

	"zero-admin/common/xerr"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/test/model"
)

type MemberLevelInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelInfoLogic {
	return &MemberLevelInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelInfoLogic) MemberLevelInfo(in *umsclient.MemberLevelInfoReq) (*umsclient.MemberLevelInfoResp, error) {
	memberLevel, err := l.svcCtx.UmsMemberLevelModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", in.Id, err)
	}
	if memberLevel == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.Id)
	}
	var resMemberLevel umsclient.MemberLevelListData
	_ = copier.Copy(&resMemberLevel, memberLevel)

	return &umsclient.MemberLevelInfoResp{
		Info: &resMemberLevel,
	}, nil
}
