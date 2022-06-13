package conf

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/gorm"
)

var (
	MySQL *gorm.DB
	OSS   *oss.Bucket
)
