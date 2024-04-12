package constant

import (
	"testing"

	"github.com/keshu12345/guardianlink/nodeb/constant"
	"github.com/stretchr/testify/assert"
)

func TestConstantsAndToString(t *testing.T) {

	tests := []struct {
		name        string
		constant    Nodea
		expectedStr string
	}{
		{
			name:        "Databasename",
			constant:    Databasename,
			expectedStr: "nodea.db",
		},
		{
			name:        "Drivername",
			constant:    Drivername,
			expectedStr: "sqlite3",
		},
		{
			name:        "NodeAURL",
			constant:    NodeBURL,
			expectedStr: "http://localhost:8081/api/blocks",
		},
		{
			name:        "ToString",
			constant:    Nodea("test"),
			expectedStr: "test",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStr, tc.constant.ToString(), "Unexpected string representation of constant")
			assert.Equal(t, tc.expectedStr, constant.Nodea(tc.expectedStr).ToString(), "Unexpected string representation of constant")
		})
	}
}
