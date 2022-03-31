package daily_record

import (
	"go-stock/app/business/daily_record"
)

type Handler struct {
	dailyRecordBus daily_record.Interface
}

func NewsHandler() *Handler {
	return &Handler{
		dailyRecordBus: daily_record.NewBusiness(),
	}
}
