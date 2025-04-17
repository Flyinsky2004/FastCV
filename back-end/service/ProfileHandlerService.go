package service

import (
	"example/config"
	"example/entity"
	"example/utility"

	"github.com/gin-gonic/gin"
)

type ProfileCreateRequestBody struct {
	Content string `json:"content" binding:"required"`
}

type ProfileUpdateRequestBody struct {
	ID      int    `json:"profile_id"`
	Content string `json:"content" binding:"required"`
}

// CreateProfile 创建个人资料
func CreateProfile(c *gin.Context) {
	var body ProfileCreateRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(200, entity.ErrorResponse[string](400, "提交了错误的JSON格式"))
		return
	}

	// 从 token 中获取用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(200, entity.ErrorResponse[string](401, "未登录或 token 已过期"))
		return
	}

	profile := entity.Profile{
		UserID:  userID.(int),
		Content: body.Content,
	}

	if err := config.MysqlDataBase.Create(&profile).Error; err != nil {
		c.JSON(200, entity.ErrorResponse[string](500, "创建个人资料失败"))
		return
	}

	c.JSON(200, entity.SuccessResponse("个人资料创建成功"))
}

// GetProfile 获取个人资料
func GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt := utility.GetUserIDFromContext(userID)

	var profile []entity.Profile
	if err := config.MysqlDataBase.Where("user_id = ?", userIDInt).Find(&profile).Error; err != nil {
		c.JSON(200, entity.ErrorResponse[string](404, "未找到个人资料"))
		return
	}

	c.JSON(200, entity.SuccessResponse(profile))
}

// GetProfile 通过ID获取个人资料
func GetProfileByID(c *gin.Context) {
	profileID := c.Param("id")
	var profile entity.Profile
	if err := config.MysqlDataBase.Where("id = ?", profileID).Find(&profile).Error; err != nil {
		c.JSON(200, entity.ErrorResponse[string](404, "未找到个人资料"))
		return
	}

	c.JSON(200, entity.SuccessResponse(profile))
}

// UpdateProfile 更新个人资料
func UpdateProfile(c *gin.Context) {
	var body ProfileUpdateRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(200, entity.ErrorResponse[string](400, "提交了错误的JSON格式"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(200, entity.ErrorResponse[string](401, "未登录或 token 已过期"))
		return
	}

	var profile entity.Profile
	if err := config.MysqlDataBase.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(200, entity.ErrorResponse[string](404, "未找到个人资料"))
		return
	}

	profile.Content = body.Content
	if err := config.MysqlDataBase.Save(&profile).Error; err != nil {
		c.JSON(200, entity.ErrorResponse[string](500, "更新个人资料失败"))
		return
	}

	c.JSON(200, entity.SuccessResponse("个人资料更新成功"))
}

// DeleteProfile 删除个人资料
func DeleteProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(200, entity.ErrorResponse[string](401, "未登录或 token 已过期"))
		return
	}

	if err := config.MysqlDataBase.Where("user_id = ?", userID).Delete(&entity.Profile{}).Error; err != nil {
		c.JSON(200, entity.ErrorResponse[string](500, "删除个人资料失败"))
		return
	}

	c.JSON(200, entity.SuccessResponse("个人资料删除成功"))
}
