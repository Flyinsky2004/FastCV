package service

import (
	"example/config"
	"example/entity"
	"example/utility"
	"fmt"
	"log"
	"strings"

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

type GetAdveriseUniversalRequest struct {
	Source     string `json:"source"`             // 需要修改的原文内容
	ModifyDes  string `json:"modify_description"` // 修改要求描述
	FieldType  string `json:"field_type"`         // 简历字段类型（如：summary, experience, education等）
	Style      string `json:"style,omitempty"`    // 可选的风格要求
	TemplateId int    `json:"templateId,omitempty"`
}

func GetAdveriseUniversal(c *gin.Context) {
	var request GetAdveriseUniversalRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, entity.ErrorResponse[string](401, "请求参数错误"))
		return
	}

	// 根据字段类型构建不同的提示词
	basePrompt := "你是一名专业的简历优化专家，擅长帮助求职者打造有竞争力的简历。"
	fieldPrompt := ""
	switch request.FieldType {
	case "summary":
		fieldPrompt = "你需要帮助优化简历的个人总结/自我介绍部分，使其更加突出个人特点和职业目标。"
	case "experience":
		fieldPrompt = "你需要帮助优化工作经验描述，突出具体成就和贡献，使用有力的行动动词和量化数据。"
	case "education":
		fieldPrompt = "你需要帮助优化教育背景部分，突出相关的学术成就和专业技能。"
	case "skills":
		fieldPrompt = "你需要帮助优化技能部分，使其更加符合目标职位的要求，突出核心竞争力。"
	case "projects":
		fieldPrompt = "你需要帮助优化项目经验描述，突出你的角色、贡献和项目成果。"
	default:
		fieldPrompt = "你需要帮助优化简历内容，提升其整体专业性和竞争力。"
	}

	stylePrompt := ""
	if request.Style != "" {
		stylePrompt = fmt.Sprintf("请按照以下风格要求进行优化：%s", request.Style)
	}

	modifyPrompt := ""
	if request.ModifyDes != "" {
		modifyPrompt = fmt.Sprintf("具体优化要求：%s", request.ModifyDes)
	}

	sourcePrompt := ""
	if request.TemplateId != 0 {
		var source string
		if err := config.MysqlDataBase.Table("fake_templates").Select("content").Where("id = ?", request.TemplateId).Find(&source).Error; err != nil {
			log.Println(err.Error())
		} else {
			sourcePrompt = fmt.Sprintf("以下是原文内容，仅供参考以保持文章的连贯性和避免产生不实信息：%s", source)
		}
	}

	prompt := strings.Join([]string{basePrompt, fieldPrompt, stylePrompt, modifyPrompt, sourcePrompt}, " ")
	questionStr := fmt.Sprintf("这是需要优化的%s部分：%s", request.FieldType, request.Source)

	score, err := utility.ChatHandler(utility.ChatRequest{
		Model:    config.Config.OpenAI.UseModelName,
		Messages: []utility.Message{},
		Prompt:   prompt,
		Question: questionStr,
	})
	if err != nil {
		c.JSON(200, entity.ErrorResponse[string](500, "内容优化失败"))
		return
	}

	resStr := score.Choices[0].Message.Content
	c.JSON(200, entity.SuccessResponse(resStr))
}
