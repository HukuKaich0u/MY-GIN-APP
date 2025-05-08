package controllers

import (
	"my-gin-app/database"
	"my-gin-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 新規投稿の作成を行うハンドラー CREATE
func CreatePostHandler(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "投稿の作成に失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "投稿が作成されました", "post": post})
}

// 全ての投稿を取得して一覧表示する READ
func ListPostHandler(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "投稿の取得に失敗"})
		return
	}

	// HTMLテンプレートに投稿一覧をレンダリングする
	c.HTML(http.StatusOK, "posts/list.html", gin.H{"posts": posts})
}

// 個別の投稿を取得して表示する READ
func ShowPostHandler(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "投稿が見つかりません"})
		return
	}

	c.HTML(http.StatusOK, "posts/detail.html", gin.H{"post": post})
}

// 投稿の編集画面を表示するGETリクエスト用のハンドラー
func EditPostHandler(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "投稿が見つかりません"})
		return
	}

	// 編集フォームに現在の投稿データを渡す
	c.HTML(http.StatusOK, "posts/edit.html", gin.H{"post": post})
}

// 既存の投稿を更新する UPDATE
func UpdatePostHandler(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	// 既存の投稿を取得
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "投稿が見つかりません"})
		return
	}

	// フォームから送信されたデータで投稿を更新
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.UpdatedAt = time.Now()
	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新に失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "投稿が更新されました", "post": post})
}

// 不要な投稿を削除する処理
func DeletePostHandler(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Post{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "投稿が削除されました"})
	}
}
