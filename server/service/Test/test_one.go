package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Test"
	TestReq "github.com/flipped-aurora/gin-vue-admin/server/model/Test/request"
)

type TestOneService struct {
}

// CreateTestOne 创建测试一记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_oneService *TestOneService) CreateTestOne(test_one *Test.TestOne) (err error) {
	err = global.GVA_DB.Create(test_one).Error
	return err
}

// DeleteTestOne 删除测试一记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_oneService *TestOneService) DeleteTestOne(ID string) (err error) {
	err = global.GVA_DB.Delete(&Test.TestOne{}, "id = ?", ID).Error
	return err
}

// DeleteTestOneByIds 批量删除测试一记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_oneService *TestOneService) DeleteTestOneByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Test.TestOne{}, "id in ?", IDs).Error
	return err
}

// UpdateTestOne 更新测试一记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_oneService *TestOneService) UpdateTestOne(test_one Test.TestOne) (err error) {
	err = global.GVA_DB.Save(&test_one).Error
	return err
}

// GetTestOne 根据ID获取测试一记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_oneService *TestOneService) GetTestOne(ID string) (test_one Test.TestOne, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&test_one).Error
	return
}

// GetTestOneInfoList 分页获取测试一记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_oneService *TestOneService) GetTestOneInfoList(info TestReq.TestOneSearch) (list []Test.TestOne, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Test.TestOne{})
	var test_ones []Test.TestOne
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&test_ones).Error
	return test_ones, total, err
}
