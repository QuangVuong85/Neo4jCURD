package main

import (
	"Neo4jCURD/appconfig"
	_ "Neo4jCURD/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	config := &appconfig.AppConfig{}
	config.Neo4jSid = beego.AppConfig.String("bigset_neo4j::sid")
	config.Neo4jHost = beego.AppConfig.String("bigset_neo4j::host")
	config.Neo4jPort = beego.AppConfig.String("bigset_neo4j::port")
	config.EtcdServerEndpoints = beego.AppConfig.Strings("default::etcd")
	appconfig.Config = config

	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
