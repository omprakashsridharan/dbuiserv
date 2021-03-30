package main

type ConnectionOptions struct {
	DbType   string `json:"dbType" binding:"required"`
	Hostname string `json:"hostname" binding:"required"`
	Port     int32  `json:"port" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
