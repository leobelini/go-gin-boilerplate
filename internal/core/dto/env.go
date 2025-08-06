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

type DtoEnvAppApp struct {
	Name string
	URL  string
}

type DtoEnvAppSmtp struct {
	Host string
	Port int
	From string
}

type DtoEnvApp struct {
	Server   DtoEnvAppServer
	Database DtoEnvAppDatabase
	Redis    DtoEnvAppRedis
	App      DtoEnvAppApp
	Smtp     DtoEnvAppSmtp
	IsProd   bool
}
