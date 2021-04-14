package router

import (
	"github.com/Kubernetes-extension/preset-pod/imple"
	"github.com/Kubernetes-extension/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	//重新定义404
	r.NoRoute(NoRoute)
	// Pod IP 地址固定
	r.POST(common.PresetPath+"mutate/fixpodip", impl.MutateFixPodIP)
	r.POST(common.PresetPath+"validate/fixpodip", impl.ValidateFixPodIP)

	return r
}

func NoRoute(c *gin.Context) {
	responseData := common.ResponseData{Code: http.StatusNotFound, Msg: "404 Not Found"}
	c.JSON(http.StatusNotFound, responseData)
}
