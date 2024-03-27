package subjectproductrelationservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductRelationListLogic {
	return &SubjectProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SubjectProductRelationList 查询关联专题
func (l *SubjectProductRelationListLogic) SubjectProductRelationList(in *cmsclient.SubjectProductRelationListReq) (*cmsclient.SubjectProductRelationListResp, error) {
	all, err := l.svcCtx.CmsPrefrenceAreaProductRelationModel.FindAll(l.ctx, in.ProductId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询关联专题列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []int64
	for _, item := range *all {
		list = append(list, item.PrefrenceAreaId)
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询关联专题列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &cmsclient.SubjectProductRelationListResp{
		SubjectId: list,
	}, nil
}
