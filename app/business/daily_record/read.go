package daily_record

import (
	"go-stock/app/global/helper"
	"go-stock/app/global/request"
)

func (b *business) CreateDailyRecord(option request.CreateDailyRecordOption) (err error) {
	dataMap, _ := helper.StructToMap(option)

	b.dailyRecord.CreateByMap(dataMap)

	return
}
