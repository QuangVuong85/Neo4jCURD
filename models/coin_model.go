package models

import (
	"Neo4jCURD/consts"
	"Neo4jCURD/helps"
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
	"time"
)

var bigsetIf StringBigsetService.StringBigsetServiceIf

func InitModel() {
	bigsetIf = GetBigSet("test", "127.0.0.1", "18990")
}

//go:generate easytags $GOFILE json,xml

type Coin struct {
	Coin         string  `json:"coin" xml:"coin"`
	Symbol       string  `json:"symbol" xml:"symbol"`
	CoinImage    string  `json:"icon_image" xml:"icon_image"`
	Name         string  `json:"name" xml:"name"`
	Confirmation int     `json:"confirmation" xml:"confirmation"`
	Decimals     float64 `json:"decimals" xml:"decimals"`
	Type         string  `json:"type" xml:"type"`
	//WalletId            uint64  `json:"wallet_id" xml:"wallet_id"`
	//Env                 string  `json:"env" xml:"env"`
	ContractAddress     string  `json:"contract_address" xml:"contract_address"`
	TransactionTxPath   string  `json:"transaction_tx_path" xml:"transaction_tx_path"`
	TransactionExplorer string  `json:"transaction_explorer" xml:"transaction_explorer"`
	WithdrawalThreshold float64 `json:"withdrawal_threshold" xml:"withdrawal_threshold"`
	CreatedAt           int64   `json:"created_at" xml:"created_at"`
	UpdatedAt           int64   `json:"updated_at" xml:"updated_at"`
}

func (this *Coin) String() string {
	return this.Coin
}

func (this *Coin) GetBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s:%s", consts.BS_COIN))
}

func (this *Coin) GetAll() ([]Coin, int64, error) {
	var err error
	if totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey()); totalCount > 0 && (err == nil || !helps.IsError(err)) {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(), 0, int32(totalCount))
		if helps.IsError(err) {
			return make([]Coin, 0), 0, err
		}
		coin, err := this.UnMarshalArrayTItem(slice)
		return coin, totalCount, err
	}

	return make([]Coin, 0), 0, err
}

func (this *Coin) GetPaginate(pos, count int32) ([]Coin, int64, error) {
	totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey())
	if helps.IsError(err) || totalCount < 1 {
		return nil, 0, err
	}

	setItems, err := bigsetIf.BsGetSlice(this.GetBsKey(), pos, count)
	if helps.IsError(err) {
		return nil, 0, err
	}

	coins, err := this.UnMarshalArrayTItem(setItems)
	if err != nil {
		return nil, 0, err
	}

	return coins, totalCount, err
}

func (this *Coin) Create() error {
	now := time.Now().UnixNano() / 1000
	this.CreatedAt = now
	this.UpdatedAt = now

	bCoin, key, err := helps.MarshalBytes(this)
	if err != nil {
		return err
	}

	return bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bCoin,
	})
}

func (this *Coin) PutItem() error {
	this.UpdateTime()

	bCoin, key, err := helps.MarshalBytes(this)

	if err != nil {
		return err
	}

	return bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bCoin,
	})
}

func (this *Coin) Delete() error {
	return bigsetIf.BsRemoveItem(this.GetBsKey(), []byte(this.String()))
}

func (this *Coin) Get() (interface{}, error) {
	bytes, err := this.GetItemBytes()
	if err != nil {
		return nil, err
	}

	return helps.UnMarshalBytes(bytes)
}

func (this *Coin) GetItemBytes() ([]byte, error) {
	tCoin, err := bigsetIf.BsGetItem(this.GetBsKey(), generic.TItemKey(this.String()))
	if helps.IsError(err) {
		return nil, err
	}

	return tCoin.GetValue(), nil
}

func (this *Coin) UnMarshalArrayTItem(objects []*generic.TItem) ([]Coin, error) {
	objs := make([]Coin, 0)

	for _, object := range objects {
		obj := Coin{}
		log.Println(string(object.GetValue()), "-- string(object.GetValue())")
		err := json.Unmarshal(object.GetValue(), &obj)
		log.Println(string(object.Key), ": Key")

		if err != nil {
			return make([]Coin, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (this *Coin) GetFromKey(key string) (*Coin, error) {
	item, err := bigsetIf.BsGetItem(this.GetBsKey(), generic.TItemKey(key))
	if helps.IsError(err) {
		return nil, err
	}
	coin := &Coin{}
	err = json.Unmarshal(item.GetValue(), &coin)
	if err != nil {
		return nil, err
	}
	return coin, nil
}

func (this *Coin) UpdateTime() {
	this.UpdatedAt = time.Now().UnixNano() / 1000
}
