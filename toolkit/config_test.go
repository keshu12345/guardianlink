package toolkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type usecase struct {
	name     string
	testcase func(*testing.T)
}

type Configuration struct {
	EnvironmentName string
	Server          Server  `mapstructure:"server"`
	Swagger         Swagger `mapstructure:"swagger"`
	SQLite          DB      `mapstructure:"db"`
}
type Swagger struct {
	Host string
}

type Server struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type DB struct {
	Filepath     string
	Databasename string
}

func Test_NewConfig(t *testing.T) {

	tests := []usecase{

		{
			name: "SUCCESS CASE",
			testcase: func(t *testing.T) {

				cfg := Configuration{
					EnvironmentName: "",
					Server: Server{
						Port:         0,
						ReadTimeout:  0,
						WriteTimeout: 0,
						IdleTimeout:  0,
					},
					Swagger: Swagger{
						Host: "",
					},
					SQLite: DB{
						Filepath:     "",
						Databasename: "",
					},
				}
				err := NewConfig(&cfg, "local/server.yml", "local/server.yml")
				assert.NoError(t, err)
			},
		},

		{
			name: "FAILURE  CASE",
			testcase: func(t *testing.T) {

				cfg := Configuration{
					EnvironmentName: "",
					Server: Server{
						Port:         0,
						ReadTimeout:  0,
						WriteTimeout: 0,
						IdleTimeout:  0,
					},
					Swagger: Swagger{
						Host: "",
					},
					SQLite: DB{
						Filepath:     "",
						Databasename: "",
					},
				}
				err := NewConfig(&cfg, "config/local/server.yml", "config/local/server.yml")
				assert.Error(t, err)
			},
		},

		{
			name: "FAILURE  CASE",
			testcase: func(t *testing.T) {

				cfg := Configuration{
					EnvironmentName: "",
					Server: Server{
						Port:         0,
						ReadTimeout:  0,
						WriteTimeout: 0,
						IdleTimeout:  0,
					},
					Swagger: Swagger{
						Host: "",
					},
					SQLite: DB{
						Filepath:     "",
						Databasename: "",
					},
				}

				listOfMaps := map[string]string{

					"key1": "value1",
					"key2": "value2",

					"key3": "value3",
					"key4": "value4",
				}

				err := NewConfig(&cfg, "config/local/server.yml", "config/local/server.yml", listOfMaps)
				assert.Error(t, err)
			},
		},

		{
			name: "FAILURE  CASE",
			testcase: func(t *testing.T) {

				cfg := Configuration{
					EnvironmentName: "",
					Server: Server{
						Port:         0,
						ReadTimeout:  0,
						WriteTimeout: 0,
						IdleTimeout:  0,
					},
					Swagger: Swagger{
						Host: "",
					},
					SQLite: DB{
						Filepath:     "",
						Databasename: "",
					},
				}

				mp := map[string]string{}
				err := NewConfig(&cfg, "local/server.yml", "local/server.yml", mp)
				assert.NoError(t, err)
			},
		},

		{
			name: "PATH EMPTY",
			testcase: func(t *testing.T) {

				cfg := Configuration{
					EnvironmentName: "",
					Server: Server{
						Port:         0,
						ReadTimeout:  0,
						WriteTimeout: 0,
						IdleTimeout:  0,
					},
					Swagger: Swagger{
						Host: "",
					},
					SQLite: DB{
						Filepath:     "",
						Databasename: "",
					},
				}

				err := NewConfig(&cfg, "", "", nil)
				assert.Error(t, err)
			},
		},

		{
			name: "EMPTY CASE",
			testcase: func(t *testing.T) {
				err := NewConfig(nil, "", "")
				assert.Error(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testcase)
	}
}
