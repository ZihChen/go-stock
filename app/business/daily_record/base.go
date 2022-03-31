package daily_record

import (
	"go-stock/app/global/request"
	"go-stock/app/repository/daily_record"
	"sync"
)

var singleton *business
var once sync.Once

type Interface interface {
	CreateDailyRecord(option request.CreateDailyRecordOption) (err error)
}

type business struct {
	dailyRecord daily_record.Interface
}

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			dailyRecord: daily_record.NewRepo(),
		}
	})
	return singleton
}
