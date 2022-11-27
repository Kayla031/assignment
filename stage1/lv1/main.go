package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 储存用户信息
var data = make(map[string]string)

func adduser(username, password string) {
	data[username] = password
}

func selectuser(username string) bool {
	if data[username] == "" {
		return false
	}
	return true
}

// 注册段
func register(c *gin.Context) {
	var username string
	var password string

	// 传入用户名和密码
	username = c.PostForm("username")
	password = c.PostForm("password")

	data[username] = password
	// 验证用户名是否重复
	flag := selectuser(username)
	// 重复则退出
	if flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "用户已存在",
		}) //StatusInternalServerError
		return
	}

	adduser(username, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "添加用户成功",
	})
}

// 登录段
func login(c *gin.Context) {
	var username string
	var password string

	// 传入用户名和密码
	username = c.PostForm("username")
	password = c.PostForm("password")

	value, ok := data[username]
	if ok == false {
		c.JSON(200, "用户未注册")
		return
	}
	if value != password {
		c.JSON(200, "密码错误！")
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "登录成功",
	})
}
func main() {
	router := gin.Default()
	user := router.Group("/user")
	{
		user.POST("/register", register)
		user.POST("/login", login)
	}

	// 跑在 9000 端口上
	router.Run(":9000")
}
