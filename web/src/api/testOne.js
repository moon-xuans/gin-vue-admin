import service from '@/utils/request'

// @Tags TestOne
// @Summary 创建测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestOne true "创建测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /test_one/createTestOne [post]
export const createTestOne = (data) => {
  return service({
    url: '/test_one/createTestOne',
    method: 'post',
    data
  })
}

// @Tags TestOne
// @Summary 删除测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestOne true "删除测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_one/deleteTestOne [delete]
export const deleteTestOne = (params) => {
  return service({
    url: '/test_one/deleteTestOne',
    method: 'delete',
    params
  })
}

// @Tags TestOne
// @Summary 批量删除测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_one/deleteTestOne [delete]
export const deleteTestOneByIds = (params) => {
  return service({
    url: '/test_one/deleteTestOneByIds',
    method: 'delete',
    params
  })
}

// @Tags TestOne
// @Summary 更新测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestOne true "更新测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test_one/updateTestOne [put]
export const updateTestOne = (data) => {
  return service({
    url: '/test_one/updateTestOne',
    method: 'put',
    data
  })
}

// @Tags TestOne
// @Summary 用id查询测试一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestOne true "用id查询测试一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test_one/findTestOne [get]
export const findTestOne = (params) => {
  return service({
    url: '/test_one/findTestOne',
    method: 'get',
    params
  })
}

// @Tags TestOne
// @Summary 分页获取测试一列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试一列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_one/getTestOneList [get]
export const getTestOneList = (params) => {
  return service({
    url: '/test_one/getTestOneList',
    method: 'get',
    params
  })
}
