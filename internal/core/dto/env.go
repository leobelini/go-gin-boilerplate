package dto

type DtoEnvAppServer struct {
	Port int
	Host string
}

type DtoEnvAppDatabase struct {
	File        string
	AutoMigrate bool
}

type DtoEnvAppRedis struct {
	Host string
	Port int
}

type DtoEnvApp struct {
	Server   DtoEnvAppServer
	Database DtoEnvAppDatabase
	Redis    DtoEnvAppRedis
	IsProd   bool
}
