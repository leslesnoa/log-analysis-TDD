package usecase

import (
	"log-analysis/models"
	"net"
	"time"
)

type LogAnalysisUseCase struct {
	ServerLogs []models.ServerLog
}

func NewLogAnalysisUseCase(serverLog []models.ServerLog) *LogAnalysisUseCase {
	return &LogAnalysisUseCase{
		ServerLogs: serverLog,
	}
}

type BreakServer struct {
	Address   net.IP
	BreakSpan *time.Duration
}

const pingOutStr = "-"

func (l *LogAnalysisUseCase) SearchBreakServer(limitCount int) []BreakServer {
	ret := []BreakServer{}

	pingOutMap := make(map[string]*struct {
		FirstBreakTime   time.Time
		Counter          int
		RecoveryInterval *time.Duration
	}, len(l.ServerLogs))
	for _, log := range l.ServerLogs {
		if log.Result == pingOutStr {
			if t, exists := pingOutMap[log.Address.String()]; exists {
				pingOutMap[log.Address.String()].Counter = t.Counter + 1
			} else {
				pingOutMap[log.Address.String()] = &struct {
					FirstBreakTime   time.Time
					Counter          int
					RecoveryInterval *time.Duration
				}{log.Datetime, 1, nil}
			}
		} else if t, ok := pingOutMap[log.Address.String()]; ok {
			// 指定回数失敗していた場合、最初のpingアウトから復旧するまで時間を格納する
			if t.Counter >= limitCount {
				sub := log.Datetime.Sub(t.FirstBreakTime)
				pingOutMap[log.Address.String()].RecoveryInterval = &sub
			} else {
				delete(pingOutMap, log.Address.String())
			}
		}
	}
	// 指定回数以上連続してpingアウトしたサーバのみ故障サーバとみなす
	for k, v := range pingOutMap {
		if v.Counter >= limitCount {
			ret = append(ret, BreakServer{
				Address:   net.ParseIP(k),
				BreakSpan: v.RecoveryInterval,
			})
		}
	}
	return ret
}
