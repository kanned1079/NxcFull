package Document

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
