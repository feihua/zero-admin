package upload

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UploadLogic 文件上传
/*
Author: LiuFeiHua
Date: 2024/5/27 9:23
*/
type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

// Upload 文件上传
// 参考：https://github.com/zeromicro/zero-examples/blob/main/monolithic/internal/logic/uploadlogic.go
func (l *UploadLogic) Upload() (resp *types.UploadResp, err error) {
	file, handler, err := l.r.FormFile("file")
	logx.WithContext(l.ctx).Infof("File uploaded successfully: %s", handler.Filename)
	if err != nil {
		logc.Errorf(l.ctx, "Error retrieving the file,异常:%s", err.Error())
		return nil, err
	}
	defer file.Close()

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		logc.Errorf(l.ctx, "Failed to create directory,异常:%s", err.Error())
		return nil, err
	}
	// 创建一个新的文件来保存上传的文件
	uploadedFile, err := os.Create(filepath.Join("./uploads", handler.Filename))
	if err != nil {
		logc.Errorf(l.ctx, "Error creating the filee,异常:%s", err.Error())
		return nil, err
	}
	defer uploadedFile.Close()

	// 将上传的文件拷贝到新文件中
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		logc.Errorf(l.ctx, "Error copying the file,异常:%s", err.Error())
		return nil, err
	}

	logx.WithContext(l.ctx).Infof("File uploaded successfully: %s", handler.Filename)
	return &types.UploadResp{
		Code:    "000000",
		Message: "上传文件成功",
		Data:    "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20181113/movie_ad.jpg",
	}, nil
}
