package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name     string
	testcase func(*testing.T)
}

func Test_NewSQLiteInstnace(t *testing.T) {

	tests := []testcase{
		{
			name: "SUCCESS CASE",
			testcase: func(t *testing.T) {

				testDBFileName := "test_nodea.db"

				defer func() {
					os.Remove(testDBFileName)
				}()

				sqlite, err := NewSQLiteInstnace()
				assert.NoError(t, err, "Should not error when initializing SQLite instance")
				assert.NotNil(t, sqlite.DB, "DB should not be nil after successful initialization")
				defer sqlite.DB.Db.Close()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testcase)
	}
}
