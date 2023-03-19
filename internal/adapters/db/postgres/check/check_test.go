package check

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"testing"
)

func testOrder() entity.OrderCheck {
	m := map[string]interface{}{"vegetables": 1, "pork": 3}
	indent, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return entity.OrderCheck{}
	}
	orderCheck := entity.NewCheckBuilder().SetCheckType(entity.Kitchen).SetStatus("created").SetPrinterId("111").SetOrder(string(indent)).Build()
	return orderCheck
}

func Test_storage_Create(t *testing.T) {
	DSN := "host=localhost port=5432 user=golobar password=password dbname=golo sslmode=disable"
	driver := "postgres"

	tests := []struct {
		name     string
		dbClient *sqlx.DB
		check    entity.OrderCheck
		wantErr  bool
	}{
		{name: "Simple", dbClient: sqlx.MustConnect(driver, DSN), check: testOrder(), wantErr: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage{
				dbClient: tt.dbClient,
			}
			got, err := s.Create(tt.check)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func Test_storage_GetAllGeneratedChecks(t *testing.T) {
	DSN := "host=localhost port=5432 user=golobar password=password dbname=golo sslmode=disable"
	driver := "postgres"

	tests := []struct {
		name     string
		dbClient *sqlx.DB
		wantErr  bool
	}{
		{name: "Simple", dbClient: sqlx.MustConnect(driver, DSN), wantErr: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage{
				dbClient: tt.dbClient,
			}
			got, err := s.GetAllGeneratedChecks()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllGeneratedChecks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func Test_storage_Get(t *testing.T) {
	DSN := "host=localhost port=5432 user=golobar password=password dbname=golo sslmode=disable"
	driver := "postgres"

	tests := []struct {
		name     string
		dbClient *sqlx.DB
		id       int
		wantErr  bool
	}{
		{name: "Simple", dbClient: sqlx.MustConnect(driver, DSN), id: 16, wantErr: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage{
				dbClient: tt.dbClient,
			}
			got, err := s.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
