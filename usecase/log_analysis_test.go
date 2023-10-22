package usecase

import (
	"log-analysis/config"
	csvreader "log-analysis/csv_reader"
	"log-analysis/models"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearchBreakServer(t *testing.T) {
	t.Run("test break server exists", func(t *testing.T) {
		csv := "../testdata/breakServer/exists.csv"
		conf := config.Config{CSVFile: csv}
		rows := csvreader.ReadCSV(conf)
		l := models.NewServerLog(rows)
		uc := NewLogAnalysisUseCase(l)
		expected := []BreakServer{
			{
				net.ParseIP("192.168.1.1"),
				1 * time.Second,
			},
			{
				net.ParseIP("10.20.30.1"),
				2 * time.Second,
			},
		}
		actual := uc.SearchBreakServer()

		assert.Equal(t, expected, actual)
	})
	t.Run("test not exists break server", func(t *testing.T) {
		csv := "../testdata/breakServer/not_exists.csv"
		conf := config.Config{CSVFile: csv}
		rows := csvreader.ReadCSV(conf)
		l := models.NewServerLog(rows)
		uc := NewLogAnalysisUseCase(l)
		expected := []BreakServer{}

		actual := uc.SearchBreakServer()

		assert.Equal(t, expected, actual)
	})
}
