package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestOneRouter struct {
}

// InitTestOneRouter 初始化 测试一 路由信息
func (s *TestOneRouter) InitTestOneRouter(Router *gin.RouterGroup) {
	test_oneRouter := Router.Group("test_one").Use(middleware.OperationRecord())
	test_oneRouterWithoutRecord := Router.Group("test_one")
	var test_oneApi = v1.ApiGroupApp.TestApiGroup.TestOneApi
	{
		test_oneRouter.POST("createTestOne", test_oneApi.CreateTestOne)             // 新建测试一
		test_oneRouter.DELETE("deleteTestOne", test_oneApi.DeleteTestOne)           // 删除测试一
		test_oneRouter.DELETE("deleteTestOneByIds", test_oneApi.DeleteTestOneByIds) // 批量删除测试一
		test_oneRouter.PUT("updateTestOne", test_oneApi.UpdateTestOne)              // 更新测试一
	}
	{
		test_oneRouterWithoutRecord.GET("findTestOne", test_oneApi.FindTestOne)       // 根据ID获取测试一
		test_oneRouterWithoutRecord.GET("getTestOneList", test_oneApi.GetTestOneList) // 获取测试一列表
	}
}
