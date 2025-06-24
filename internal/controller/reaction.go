package controller

import (
	"net/http"
	"strconv"

	"emaction/internal/model"
	"emaction/internal/service"
	"emaction/internal/until"

	"github.com/gin-gonic/gin"
)

// GetReactions 获取 Emoji 点击列表
func GetReactions(c *gin.Context) {
	targetID := c.Query("targetId")
	if targetID == "" {
		c.JSON(http.StatusBadRequest, until.FailWithMessage("Empty targetId"))
		return
	}

	reactions, err := service.GetReactions(targetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, until.FailWithMessage("Failed to get reactions"))
		return
	}

	// 确保返回空数组而不是 null
	if reactions == nil {
		reactions = []model.ReactionResponse{}
	}

	c.JSON(http.StatusOK, until.OkWithData(map[string]interface{}{
		"reactionsGot": reactions,
	}))
}

// UpdateReaction 更新 Emoji 点击
func UpdateReaction(c *gin.Context) {
	targetID := c.Query("targetId")
	reactionName := c.Query("reaction_name")
	diffStr := c.Query("diff")

	if targetID == "" || reactionName == "" || diffStr == "" {
		c.JSON(http.StatusBadRequest, until.FailWithMessage("Invalid Response."))
		return
	}

	diff, err := strconv.Atoi(diffStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, until.FailWithMessage("Invalid diff parameter"))
		return
	}

	// 标准化 diff 值
	if diff > 0 {
		diff = 1
	} else if diff < 0 {
		diff = -1
	} else {
		c.JSON(http.StatusBadRequest, until.FailWithMessage("Invalid diff value"))
		return
	}

	err = service.UpdateReaction(targetID, reactionName, diff)
	if err != nil {
		c.JSON(http.StatusInternalServerError, until.Fail())
		return
	}

	c.JSON(http.StatusOK, until.Ok())
}
