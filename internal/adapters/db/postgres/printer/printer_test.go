package printer

import (
	"github.com/jmoiron/sqlx"
	"testing"
)

func Test_printerStorage_GetByPoint(t *testing.T) {
	DSN := "host=localhost port=5432 user=golobar password=password dbname=golo sslmode=disable"
	driver := "postgres"
	tests := []struct {
		name     string
		dbClient *sqlx.DB
		pointID  int
		wantErr  bool
	}{
		{name: "simple", dbClient: sqlx.MustConnect(driver, DSN), pointID: 1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := printerStorage{
				dbClient: tt.dbClient,
			}
			gotPrinters, err := p.GetByPoint(tt.pointID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotPrinters)
		})
	}
}

func Test_printerStorage_GetAll(t *testing.T) {
	DSN := "host=localhost port=5432 user=golobar password=password dbname=golo sslmode=disable"
	driver := "postgres"
	tests := []struct {
		name     string
		dbClient *sqlx.DB

		wantErr bool
	}{
		{name: "simple", dbClient: sqlx.MustConnect(driver, DSN), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := printerStorage{
				dbClient: tt.dbClient,
			}
			gotPrinters, err := p.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotPrinters)
		})
	}
}
