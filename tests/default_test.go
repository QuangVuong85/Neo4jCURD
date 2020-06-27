package test

import (
	"Neo4jCURD/appconfig"
	"Neo4jCURD/consts"
	"Neo4jCURD/models"
	_ "Neo4jCURD/routers"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"testing"
)

func init() {
	err := beego.LoadAppConfig("ini", consts.PATH_CONFIG_FILE)
	if err != nil {
		log.Println(err.Error(), ": err")
		return
	}
	config := &appconfig.AppConfig{}
	config.Neo4jSid = beego.AppConfig.String("bigset_neo4j::sid")
	config.Neo4jHost = beego.AppConfig.String("bigset_neo4j::host")
	config.Neo4jPort = beego.AppConfig.String("bigset_neo4j::port")
	config.EtcdServerEndpoints = beego.AppConfig.Strings("default::etcd")
	appconfig.Config = config

	models.InitModel()
}

func TestCreate(t *testing.T) {
	coinIns := models.Coin{
		Coin:                "test",
		Symbol:              "test",
		CoinImage:           "test",
		Name:                "test",
		Confirmation:        0,
		Decimals:            0,
		Type:                "",
		ContractAddress:     "",
		TransactionTxPath:   "",
		TransactionExplorer: "",
		WithdrawalThreshold: 0,
	}

	err := coinIns.Create()
	if err != nil {
		log.Println(err.Error(), ": err.error")
	}
}

func TestGet(t *testing.T) {
	coinIns := models.Coin{
		Coin: "test",
	}

	coin, err := coinIns.GetFromKey(coinIns.String())
	if err != nil {
		log.Println(err.Error(), ": err.error")
		return
	}
	bData, err := json.Marshal(coin)
	if err != nil {
		log.Println(err.Error(), ": erasdfajsdf")
		return
	}

	log.Println(string(bData))
}