package models

import (
	"log"
	"net"
	"time"
)

type ServerLog struct {
	Datetime time.Time
	Address  net.IP
	Result   string
}

var layout = "20060102150405"

func NewServerLog(csv [][]string) []ServerLog {
	var ret []ServerLog
	for _, v := range csv {
		t, err := time.Parse(layout, v[0])
		if err != nil {
			log.Fatalf("time parsing error. %v\n", err)
		}
		ip, _, err := net.ParseCIDR(v[1])
		if err != nil {
			log.Fatalf("ip parsing error. %v\n", err)
		}
		slog := ServerLog{
			Datetime: t,
			Address:  ip,
			Result:   v[2],
		}
		ret = append(ret, slog)
	}
	return ret
}
