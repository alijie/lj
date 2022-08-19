package mysqlDB

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_getDBMysql(t *testing.T) {
	tests := []struct {
		name string
		want *sql.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDBMysql(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDBMysql() = %v, want %v", got, tt.want)
			}
		})
	}
}
