package main

import (
	"net/http"
	"tmp-code/cache"
	"tmp-code/data"
	"tmp-code/dist"
	"tmp-code/dist/assets"
	"tmp-code/tool"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置生产模式
	gin.SetMode(gin.ReleaseMode)

	// 获取一个id work
	work := tool.NewWorker()

	// 获取一个cache
	cache := cache.NewCache()

	// 基本路由
	r := gin.Default()

	// 处理跨域
	r.Use(tool.Cors())

	r.StaticFS("/assets", http.FS(assets.Assets))

	r.GET("/", func(c *gin.Context) {
		c.Header("content-type", "text/html;charset=utf-8")
		c.String(200, dist.Index)
	})
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Content-Type", "image/vnd.microsoft.icon")
		c.String(200, string(dist.Icon))
	})
	// 写
	r.POST("/", func(c *gin.Context) {

		// 获取code
		code := c.PostForm("code")

		// 获取类型
		typeV := c.PostForm("type")

		// 过期时间
		expire := tool.HandleTime(c.PostForm("expire"))

		// 获取一个id
		id := work.ShortId()

		// 存储code
		cache.Set(id+"-code", code, expire)

		// 存储类型
		cache.Set(id+"-type", typeV, expire)

		c.String(http.StatusOK, id)
	})

	// 读
	r.GET("/get/:id", func(c *gin.Context) {
		id := c.Param("id")

		// 获取code
		code, err1 := cache.Get(id + "-code")
		// 获取type
		typeV, err2 := cache.Get(id + "-type")
		// 判断获取是否出错
		if err1 != nil || err2 != nil {
			c.String(http.StatusNotFound, "failed")
			return
		}

		newCode := data.TmpCode{
			Code: string(code.([]byte)),
			Type: string(typeV.([]byte)),
		}
		c.JSON(http.StatusOK, newCode)

	})
	r.Run("localhost:8000")
}
