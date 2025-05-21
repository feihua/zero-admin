package memberinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	query "github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateMemberInfoLogic 更新会员信息
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type UpdateMemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberInfoLogic {
	return &UpdateMemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberInfo 更新会员信息
func (l *UpdateMemberInfoLogic) UpdateMemberInfo(in *umsclient.UpdateMemberInfoReq) (*umsclient.UpdateMemberInfoResp, error) {
	q := query.UmsMemberInfo.WithContext(l.ctx)

	// 1.根据会员信息id查询会员信息是否已存在
	_, err := q.Where(query.UmsMemberInfo.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员异常")
	}

	birthday, _ := time.Parse("2006-01-02", in.Birthday)
	now := time.Now()
	item := &model.UmsMemberInfo{
		ID:         in.Id,        // 主键ID
		Nickname:   in.Nickname,  // 昵称
		Mobile:     in.Mobile,    // 手机号码
		Avatar:     in.Avatar,    // 头像
		Signature:  in.Signature, // 个性签名
		Gender:     in.Gender,    // 性别：0-未知，1-男，2-女
		Birthday:   &birthday,    // 生日
		UpdateTime: &now,         // 更新时间
	}

	// 2.会员信息存在时,则直接更新会员信息
	_, err = q.Where(query.UmsMemberInfo.ID.Eq(in.Id)).Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新会员信息失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新会员信息失败")
	}

	return &umsclient.UpdateMemberInfoResp{}, nil
}
