package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// API入参参数
type HttpRequest struct {
	Name string `json:"name"`
}

// API响应参数
type HttpRespone struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

/*
实现一个入参为name，响应为：hello name的api
这个例子中，异常信息通过status和message返回，api响应状态正常，如果需要响应400等异常状态，可以更换c.JSON(http.StatusOK, res)中的StatusOK
*/
func Hello(c *gin.Context) {
	// 声明req
	var req HttpRequest
	// 声明res并初始化
	var res = HttpRespone{}

	// 获取api请求参数
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	// 出现错误，则响应错误信息
	if err != nil {
		res.Status = 10
		res.Message = "读取请求参数错误"
		c.JSON(http.StatusOK, res)
		return
	}
	// 判断是否入参name
	if req.Name == "" {
		res.Status = 20
		res.Message = "参数name为空"
		c.JSON(http.StatusOK, res)
		return
	}
	// 正常响应 hello name
	res.Status = 0
	res.Message = "成功"
	res.Data = fmt.Sprintf("hello %v", req.Name)
	c.JSON(http.StatusOK, res)
}
