package configs

func init() {
	LoadEnv()
	DB = ConnectDB()
	AutoMigrate(DB)
}
