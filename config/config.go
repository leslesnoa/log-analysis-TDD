package config

import "flag"

type Config struct {
	CSVFile string
}

func ReadConfig() Config {
	var (
		f = flag.String("target", "testdata/breakServer/exists.csv", "test.csv")
	)
	flag.Parse()

	return Config{
		CSVFile: *f,
	}
}
