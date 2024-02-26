package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Test"
	TestReq "github.com/flipped-aurora/gin-vue-admin/server/model/Test/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestOneApi struct {
}

var test_oneService = service.ServiceGroupApp.TestServiceGroup.TestOneService

// CreateTestOne 创建测试一
// @Tags TestOne
// @Summary 创建测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Test.TestOne true "创建测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /test_one/createTestOne [post]
func (test_oneApi *TestOneApi) CreateTestOne(c *gin.Context) {
	var test_one Test.TestOne
	err := c.ShouldBindJSON(&test_one)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := test_oneService.CreateTestOne(&test_one); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestOne 删除测试一
// @Tags TestOne
// @Summary 删除测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Test.TestOne true "删除测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_one/deleteTestOne [delete]
func (test_oneApi *TestOneApi) DeleteTestOne(c *gin.Context) {
	ID := c.Query("ID")
	if err := test_oneService.DeleteTestOne(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestOneByIds 批量删除测试一
// @Tags TestOne
// @Summary 批量删除测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /test_one/deleteTestOneByIds [delete]
func (test_oneApi *TestOneApi) DeleteTestOneByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := test_oneService.DeleteTestOneByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestOne 更新测试一
// @Tags TestOne
// @Summary 更新测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Test.TestOne true "更新测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test_one/updateTestOne [put]
func (test_oneApi *TestOneApi) UpdateTestOne(c *gin.Context) {
	var test_one Test.TestOne
	err := c.ShouldBindJSON(&test_one)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := test_oneService.UpdateTestOne(test_one); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestOne 用id查询测试一
// @Tags TestOne
// @Summary 用id查询测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Test.TestOne true "用id查询测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test_one/findTestOne [get]
func (test_oneApi *TestOneApi) FindTestOne(c *gin.Context) {
	ID := c.Query("ID")
	if retest_one, err := test_oneService.GetTestOne(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retest_one": retest_one}, c)
	}
}

// GetTestOneList 分页获取测试一列表
// @Tags TestOne
// @Summary 分页获取测试一列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TestReq.TestOneSearch true "分页获取测试一列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_one/getTestOneList [get]
func (test_oneApi *TestOneApi) GetTestOneList(c *gin.Context) {
	var pageInfo TestReq.TestOneSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := test_oneService.GetTestOneInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
