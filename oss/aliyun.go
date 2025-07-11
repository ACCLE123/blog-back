package oss

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"path/filepath"
)

type AliyunOSS struct {
	client     *oss.Client
	bucketName string
	endpoint   string
}

func NewAliyunOSS(endpoint, accessKeyId, accessKeySecret, bucketName string) (*AliyunOSS, error) {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	return &AliyunOSS{
		client:     client,
		bucketName: bucketName,
		endpoint:   endpoint,
	}, nil
}

func (a *AliyunOSS) UploadImage(file multipart.File, originalFileName string) (string, error) {
	bucket, err := a.client.Bucket(a.bucketName)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	ext := filepath.Ext(originalFileName)    // 获取原始扩展名
	newFileName := uuid.New().String() + ext // 生成唯一文件名
	err = bucket.PutObject(newFileName, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.%s/%s", a.bucketName, a.endpoint, newFileName)
	return url, nil
}
