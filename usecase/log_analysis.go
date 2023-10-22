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
	BreakSpan time.Duration
}

const pingOutStr = "-"

func (l *LogAnalysisUseCase) SearchBreakServer() []BreakServer {
	ret := []BreakServer{}
	pingOutMap := make(map[string]time.Time, len(l.ServerLogs))
	for _, log := range l.ServerLogs {
		if log.Result == pingOutStr {
			pingOutMap[log.Address.String()] = log.Datetime
		} else if t, ok := pingOutMap[log.Address.String()]; ok {
			breakSpan := log.Datetime.Sub(t)
			ret = append(ret, BreakServer{
				Address:   log.Address,
				BreakSpan: breakSpan,
			})
		}
	}
	return ret
}
