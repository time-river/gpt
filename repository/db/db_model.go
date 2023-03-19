package db

//BizId
type BizId struct {
	Bizid			    int64			`gorm:"type:bigint(20);"`
}

type BizInfo struct {
	Bizid			    int64			`gorm:"type:bigint(20);"`
	Appid               int64			`gorm:"type:bigint(20);"`
}
