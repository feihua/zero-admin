package prefrenceareaproductrelationservicelogic

import (
	"context"
	"encoding/json"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaProductRelationListLogic {
	return &PrefrenceAreaProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrefrenceAreaProductRelationListLogic) PrefrenceAreaProductRelationList(in *cmsclient.PrefrenceAreaProductRelationListReq) (*cmsclient.PrefrenceAreaProductRelationListResp, error) {
	all, err := l.svcCtx.CmsSubjectProductRelationModel.FindAll(l.ctx, in.ProductId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询关联优选列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []int64
	for _, item := range *all {
		list = append(list, item.SubjectId)
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询关联优选列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &cmsclient.PrefrenceAreaProductRelationListResp{
		PrefrenceAreaId: list,
	}, nil
}
