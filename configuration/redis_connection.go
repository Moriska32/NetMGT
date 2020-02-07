package configuration

// RedisDatabaseCfg - конфигурация подключения к БД Redis
type RedisDatabaseCfg struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	Password   string `json:"password"`
	DBIndecies []int  `json:"db_indecies"`
}
