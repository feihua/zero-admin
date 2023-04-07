package auth

import (
	"context"
	"fmt"

	"zero-admin/common/tool"
	"zero-admin/common/xerr"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/ums/umsclient"

	"github.com/pkg/errors"
	"github.com/silenceper/wechat/cache"
	wechat "github.com/silenceper/wechat/v2"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrWxMiniAuthFailError = xerr.NewErrMsg("wechat mini auth fail")
var UserAuthTypeSystem string = "system"  //平台内部
var UserAuthTypeSmallWX string = "wxMini" //微信小程序

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMiniAuthReq) (resp *types.LoginAndRegisterResp, err error) {
	//1、Wechat-Mini
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.AppId,
		AppSecret: l.svcCtx.Config.WxMiniConf.Secret,
		Cache:     cache.NewMemory(),
	})

	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "发起授权请求失败 err : %v , code : %s  , authResult : %+v", err, req.Code, authResult)
	}

	//2、Parsing WeChat-Mini return data
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "解析数据失败 req : %+v , err: %v , authResult:%+v ", req, err, authResult)
	}

	//3、bind user or login.
	var userId int64
	rpcRsp, err := l.svcCtx.Ums.MemberAuthByAuthKey(l.ctx, &umsclient.MemberAuthByAuthKeyReq{
		AuthType: UserAuthTypeSmallWX,
		AuthKey:  authResult.OpenID,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "rpc call userAuthByAuthKey err : %v , authResult : %+v", err, authResult)
	}

	var data types.LoginAndRegisterData
	if rpcRsp.MemberAuth == nil || rpcRsp.MemberAuth.Id == 0 {
		//bind user.
		authPhoneNumberInfo, err := miniprogram.GetAuth().GetPhoneNumberContext(l.ctx, req.PNCode)
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "发起手机号授权请求失败 err : %v , code : %s  , authPhoneNumberInfo : %+v", err, req.PNCode, authPhoneNumberInfo)
		}
		phone := authPhoneNumberInfo.PhoneInfo.PhoneNumber
		if phone == "" {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "发起手机号授权请求失败 err : %v , code : %s  , authPhoneNumberInfo : %+v", err, req.PNCode, authPhoneNumberInfo)
		}

		nickName := userData.NickName
		// icon := userData.AvatarURL
		registerRsp, err := l.svcCtx.Ums.Register(l.ctx, &umsclient.RegisterReq{
			Phone:    phone,
			Nickname: nickName,
			Password: tool.GenerateUuid(),
			AuthKey:  authResult.OpenID,
			AuthType: UserAuthTypeSmallWX,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "UsercenterRpc.Register err :%v, authResult : %+v", err, authResult)
		}
		userId = registerRsp.MemberId
		data = types.LoginAndRegisterData{
			Token:        registerRsp.AccessToken,
			AccessExpire: registerRsp.AccessExpire,
			RefreshAfter: registerRsp.RefreshAfter,
		}

	} else {
		userId = rpcRsp.MemberAuth.MemberId
		tokenResp, err := l.svcCtx.Ums.GenerateToken(l.ctx, &umsclient.GenerateTokenReq{
			MemberId: userId,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "usercenterRpc.GenerateToken err :%v, userId : %d", err, userId)
		}
		data = types.LoginAndRegisterData{
			Token:        tokenResp.AccessToken,
			AccessExpire: tokenResp.AccessExpire,
			RefreshAfter: tokenResp.RefreshAfter,
		}
	}

	memberinfo, err := l.svcCtx.Ums.MemberInfo(l.ctx, &umsclient.MemberInfoReq{Id: userId})
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "usercenterRpc.GenerateToken err :%v, userId : %d", err, userId)
	}
	fmt.Println(memberinfo)
	data.UserInfo = types.UserInfo{
		NickName: memberinfo.Member.Nickname,
		Phone:    memberinfo.Member.Phone,
		Icon:     memberinfo.Member.Icon,
	}
	return &types.LoginAndRegisterResp{
		Errno:  0,
		Data:   data,
		Errmsg: "登录成功",
	}, nil
}
