package helpers

import "strconv"

type DbConfig struct {
	User                    string
	Host                    string
	Database                string
	Password                string
	MaxConnections          int
	ConnectionTimeoutMillis int
	IdleTimeoutMillis       int
	Port                    int
}

func GenerateParsedURLFromConfig(dbConfig DbConfig) string {
	if dbConfig.User == "" || dbConfig.Password == "" || dbConfig.Host == "" || dbConfig.Database == "" {
		return ""
	}

	return "postgres://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host +
		":" + strconv.Itoa(dbConfig.Port) + "/" + dbConfig.Database + "?pool_max_conns=" + strconv.Itoa(dbConfig.MaxConnections) +
		"&pool_max_conn_lifetime=" + strconv.Itoa(dbConfig.ConnectionTimeoutMillis) + "ms&pool_max_conn_idle_time=" + strconv.Itoa(dbConfig.IdleTimeoutMillis) + "ms"
}
