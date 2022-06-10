package request

type FavoriteActionRequest struct {
	UserId     int64 `json:"user_id" binding:"required"`
	VideoId    int64 `json:"video_id" binding:"required"`
	ActionType int64 `json:"action_type" binding:"required"`
}

type FavoriteListRequest struct {
	UserId int64 `json:"user_id" binding:"required" form:"user_id"`
}
