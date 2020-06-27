package models

import (
	"Neo4jCURD/appconfig"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"log"
)

func GetBigSet(sid, host, port string) StringBigsetService.StringBigsetServiceIf {
	key := fmt.Sprintf("%s:%s", host, port)
	log.Println("--------------")
	log.Println(key, "-- key")

	return StringBigsetService.NewStringBigsetServiceModel(appconfig.Config.Neo4jSid,
		appconfig.Config.EtcdServerEndpoints,
		GoEndpointBackendManager.EndPoint{
			Host:      host,
			Port:      port,
			ServiceID: sid,
		})
}
