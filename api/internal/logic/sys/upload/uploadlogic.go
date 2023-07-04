package upload

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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
//参考：https://github.com/zeromicro/zero-examples/blob/main/monolithic/internal/logic/uploadlogic.go
func (l *UploadLogic) Upload() (resp *types.UploadResp, err error) {
	file, handler, err := l.r.FormFile("file")
	logx.WithContext(l.ctx).Infof("File uploaded successfully: %s", handler.Filename)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error retrieving the file,异常:%s", err.Error())
		return nil, err
	}
	defer file.Close()

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Failed to create directory,异常:%s", err.Error())
		return nil, err
	}
	// 创建一个新的文件来保存上传的文件
	uploadedFile, err := os.Create(filepath.Join("./uploads", handler.Filename))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error creating the filee,异常:%s", err.Error())
		return nil, err
	}
	defer uploadedFile.Close()

	// 将上传的文件拷贝到新文件中
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error copying the file,异常:%s", err.Error())
		return nil, err
	}

	logx.WithContext(l.ctx).Infof("File uploaded successfully: %s", handler.Filename)
	return &types.UploadResp{
		Code:    "000000",
		Message: "上传文件成功",
		Data:    "https://img14.360buyimg.com/n1/s450x450_jfs/t1/158675/37/28388/40807/649eaad4Fc122761e/ee6fa54cf0c458db.jpg",
	}, nil
}
