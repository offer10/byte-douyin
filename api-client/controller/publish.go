package controller

import (
	"crypto/md5"
	"fmt"
	"mime/multipart"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/request"
	"github.com/offer10/byte-douyin/api-client/response"
	"github.com/offer10/byte-douyin/api-client/service"
	"github.com/offer10/byte-douyin/basic-server/conf"
	"github.com/offer10/byte-douyin/pb"
)

type IPublishController interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type PublishController struct{}

func NewPublishController() IPublishController {
	return PublishController{}
}

func (u PublishController) Action(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	title := form.Value["title"]
	userId := ctx.GetInt64("user_id")

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
		})
		return
	}

	ossUrl, err := UploadFile(file)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
		})
		return
	}

	// 封面图截取第一秒帧数
	// @doc https://help.aliyun.com/document_detail/64555.html?spm=a2c6h.12873639.article-detail.6.164d43eeuJdd08
	coverUrl := fmt.Sprintf("%s%s", ossUrl, "?x-oss-process=video/snapshot,t_1000,f_jpg,m_fast")

	resp, err := service.PublishClient.Action(ctx, &pb.PublishActionRequest{
		AuthorID: userId,
		Title:    title[0],
		PlayUrl:  ossUrl,
		CoverUrl: coverUrl,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"video_id":    resp.VideoID,
	})
}

func UploadFile(file *multipart.FileHeader) (ossUrl string, err error) {
	// file中没有单独列出扩展名，所以此处需要单独取一次
	fileExt := path.Ext(file.Filename)
	// 此处重命名文件名 取此时的时间戳的MD5为上传OSS的文件名
	data := []byte(time.Now().String())
	md5FileName := fmt.Sprintf("%x", md5.Sum(data))
	// 以年月为文件目录进行分类
	tTime := time.Now().Format("200601")
	// 年月/文件名.扩展名（注意不要再定义的目录前面加/）
	ossFilePath := fmt.Sprintf("videos/%s/%s%s", tTime, md5FileName, fileExt)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	err = conf.OSS.PutObject(ossFilePath, src)
	if err != nil {
		return "", err
	}

	ossUrl = fmt.Sprintf("https://%s.%s/%s", "byte-douyin", "oss-cn-hangzhou.aliyuncs.com", ossFilePath)

	return ossUrl, err
}

func (u PublishController) List(ctx *gin.Context) {
	payload := request.PublishListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.PublishClient.List(ctx, &pb.PublishListRequest{
		UserID: payload.UserId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
			"error":       err.Error(),
		})
		return
	}

	// 组装数据返回
	videoList := response.VideoList{}
	for _, video := range resp.List {
		user, _ := GetUser(ctx, video.AuthorId, GetLoginUserId(ctx))
		isFav, _ := GetIsFav(ctx, GetLoginUserId(ctx), video.Id)
		videoList = append(videoList, response.Video{
			Author:        user,
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			IsFavorite:    isFav,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"video_list":  videoList,
	})
}
