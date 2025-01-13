package services

import (
	"github.com/lining4069/kv-auth-db/backend/app/common/request"
	"github.com/lining4069/kv-auth-db/backend/global"
)

type bitcaskDBService struct {
}

var BitcaskSBService = new(bitcaskDBService)

func (bitcaskDBSer *bitcaskDBService) Put(params request.BitcaskPutRequest) (err error) {
	err = global.App.BitcaskDB.Put([]byte(params.Key), []byte(params.Value))
	return
}

func (bitcaskDBSer *bitcaskDBService) Get(key string) (err error, value string) {
	byteVal, errGet := global.App.BitcaskDB.Get([]byte(key))
	if errGet != nil {
		err = errGet
		return
	}
	value = string(byteVal)
	return
}
