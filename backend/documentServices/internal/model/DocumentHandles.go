package model

import (
	"NxcFull/nxc-backend/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllDocument(context *gin.Context) {
	lang := context.Query("lang")
	var documentsList []Document
	if result := dao.Db.Model(&Document{}).Where("language = ?", lang).Find(&documentsList); result.Error != nil {
		log.Println("获取所有文档失败")
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取所有文档失败",
		})
	}

	// 构建指定格式的数据
	result := make([]map[string]interface{}, 0)
	categories := make(map[string][]map[string]string)
	for _, doc := range documentsList {
		item := map[string]string{
			"date":  doc.CreatedAt.Format("2006-01-02"),
			"title": doc.Title,
			"body":  doc.Body,
		}
		if _, ok := categories[doc.Category]; !ok {
			categories[doc.Category] = make([]map[string]string, 0)
		}
		categories[doc.Category] = append(categories[doc.Category], item)
	}
	for category, data := range categories {
		result = append(result, map[string]interface{}{
			"category": category,
			"data":     data,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"doc_list": result,
	})
}

func HandleAddNewDocument(context *gin.Context) {
	var postData = &struct {
		Title    string `json:"title"`
		Category string `json:"category"`
		Language string `json:"language"`
		MdText   string `json:"md_text"`
	}{}
	if err := context.ShouldBindJSON(postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
	}
	//log.Println(postData)
	newDoc := Document{
		Title:    postData.Title,
		Category: postData.Category,
		Language: postData.Language,
		Body:     postData.MdText,
	}
	if err := dao.Db.Model(&Document{}).Create(&newDoc).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "插入数据库失败",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "添加成功",
	})
}

func HandleUpdate(context *gin.Context) {
	var postData = &struct {
		Id       int64  `json:"id"`
		Title    string `json:"title"`
		Category string `json:"category"`
		Language string `json:"language"`
		MdText   string `json:"md_text"`
	}{}

	// 绑定传入的 JSON 数据
	if err := context.ShouldBindJSON(postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
		return
	}

	// 查找需要更新的文档
	var doc Document
	if err := dao.Db.Model(&Document{}).Where("id = ?", postData.Id).First(&doc).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "文档未找到",
		})
		return
	}

	// 更新文档字段
	doc.Title = postData.Title
	doc.Category = postData.Category
	doc.Language = postData.Language
	doc.Body = postData.MdText

	// 保存更新
	if err := dao.Db.Save(&doc).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "更新文档失败",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "更新成功",
	})
}
