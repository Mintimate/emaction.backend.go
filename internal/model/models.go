package model

// Reaction Emoji 数据模型
type Reaction struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	TargetID     string `json:"target_id" gorm:"column:target_id;type:varchar(255);not null;index:idx_target_reaction"`
	ReactionName string `json:"reaction_name" gorm:"column:reaction_name;type:varchar(50);not null;index:idx_target_reaction"`
	Count        int    `json:"count" gorm:"column:count;type:int;not null;default:0"`
	CreatedAt    int64  `json:"created_at" gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt    int64  `json:"updated_at" gorm:"column:updated_at;type:bigint;not null"`
}

// TableName 指定表名
func (Reaction) TableName() string {
	return "reactions"
}

// ReactionResponse Emoji 查询响应
type ReactionResponse struct {
	ReactionName string `json:"reaction_name"`
	Count        int    `json:"count"`
}

// APIResponse 通用API响应格式
type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
