package request

type RelationActionRequest struct {
	UserId     int64 `json:"user_id" binding:"required"`
	FollowId   int64 `json:"follow_id"  binding:"required"`
	ActionType int32 `json:"action_type" binding:"required"`
}
type RelationFollowListRequest struct {
	UserId int64 `json:"user_id" binding:"required" form:"user_id"`
}
type RelationFollowerListRequest struct {
	UserId int64 `json:"user_id" binding:"required" form:"user_id"`
}
