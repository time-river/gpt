package db

import(
	"gorm.io/gorm"
	"sync"
)

var dbConnection *Connect
var onceConf sync.Once

type CtlClient struct{

}

type Connect struct {
	TestDataDb	*gorm.DB
}

// NewDbClient 获取db连接
func NewDbClient() *Connect{
	if dbConnection == nil{
		onceConf.Do(func() {
			dbConnection = &Connect{
				TestDataDb :connectTestDataDb(),
			}
		})
	}
	return dbConnection
}

