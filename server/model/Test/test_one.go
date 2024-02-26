// 自动生成模板TestOne
package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 测试一 结构体  TestOne
type TestOne struct {
	global.GVA_MODEL
	Test_one string `json:"test_one" form:"test_one" gorm:"column:test_one;comment:;"` //测试一
}

// TableName 测试一 TestOne自定义表名 test_one
func (TestOne) TableName() string {
	return "test_one"
}
