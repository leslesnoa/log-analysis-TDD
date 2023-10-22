package config

import "flag"

type Config struct {
	CSVFile      string
	PingOutCount int
}

func ReadConfig() Config {
	var (
		f = flag.String("target", "testdata/breakServer/exists.csv", "-file test.csv")
		p = flag.Int("out", 3, "-out 3")
	)
	flag.Parse()

	return Config{
		CSVFile:      *f,
		PingOutCount: *p,
	}
}
