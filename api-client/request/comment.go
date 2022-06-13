package request

type CommentActionRequest struct {
	UserId      int64  `json:"user_id" binding:"required"`
	VideoId     int64  `json:"video_id" binding:"required"`
	ActionType  int64  `json:"action_type" binding:"required"`
	CommentText string `json:"comment_text,omitempty"`
	CommentID   int64  `json:"comment_id,omitempty" `
}

type CommentListRequest struct {
	VideoId int64 `json:"video_id" binding:"required"  form:"video_id"`
}
