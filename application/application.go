package application

import (
	"fmt"
	"log-analysis/config"
	csvreader "log-analysis/csv_reader"
	"log-analysis/models"
	"log-analysis/usecase"
)

type Result struct {
	BreakServers []usecase.BreakServer
}

func Run() {
	conf := config.ReadConfig()
	rows := csvreader.ReadCSV(conf)
	sl := models.NewServerLog(rows)
	uc := usecase.NewLogAnalysisUseCase(sl)
	bs := uc.SearchBreakServer(conf.PingOutCount)

	result := Result{bs}
	for _, bs2 := range result.BreakServers {
		fmt.Printf("故障サーバー: %v 故障期間: %v\n", bs2.Address, bs2.BreakSpan)
	}
	fmt.Printf("合計故障サーバー数 %d\n", len(result.BreakServers))
}
