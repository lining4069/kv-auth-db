package bootstrap

import (
	bitcaskdb "github.com/lining4069/kv-auth-db/backend/app/bitcask"
	"github.com/lining4069/kv-auth-db/backend/global"
	"go.uber.org/zap"
)

func InitialBitcaskDB() *bitcaskdb.DB {
	opts := bitcaskdb.DefaultOptions
	opts.DirPath = global.App.Config.BitcaskDB.DirPath
	db, err := bitcaskdb.Open(opts)
	if err != nil {
		global.App.Log.Error("bitcask kv database init failed ", zap.Any("err", err))
	}
	return db
}
