package vo

// UpdateReactionVO 更新 Emoji 请求参数
type UpdateReactionVO struct {
	TargetID     string `json:"targetId" binding:"required"`
	ReactionName string `json:"reactionName" binding:"required"`
	Diff         int    `json:"diff" binding:"required"`
}

// GetReactionsVO 获取 Emoji 请求参数
type GetReactionsVO struct {
	TargetID string `form:"targetId" binding:"required"`
}
