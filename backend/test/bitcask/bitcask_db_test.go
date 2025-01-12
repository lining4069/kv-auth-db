package bitcask

import (
	bitcaskdb "github.com/lining4069/kv-auth-db/backend/app/bitcask"
	"github.com/lining4069/kv-auth-db/backend/app/bitcask/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func destroyDB(db *bitcaskdb.DB, dirPath string) {
	if db != nil {
		_ = db.Close()
		err := os.RemoveAll(dirPath)
		if err != nil {
			panic(err)
		}
	}
}

func TestOpen(t *testing.T) {
	opts := bitcaskdb.DefaultOptions
	opts.DirPath = "./tmp"
	db, err := bitcaskdb.Open(opts)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestBitcaskDB_Put(t *testing.T) {
	opts := bitcaskdb.DefaultOptions
	opts.DirPath = "./tmp"
	db, err := bitcaskdb.Open(opts)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	// 1.正常 Put 一条数据
	err = db.Put(utils.GetTestKey(1), utils.RandomValue(24))
	assert.Nil(t, err)
	val1, err := db.Get(utils.GetTestKey(1))
	assert.Nil(t, err)
	assert.NotNil(t, val1)

	// 2.重复 Put key 相同的数据
	err = db.Put(utils.GetTestKey(1), utils.RandomValue(24))
	assert.Nil(t, err)
	val2, err := db.Get(utils.GetTestKey(1))
	assert.Nil(t, err)
	assert.NotNil(t, val2)

	// 3.key 为空
	err = db.Put(nil, utils.RandomValue(24))
	assert.Equal(t, bitcaskdb.ErrKeyIsEmpty, err)

	// 4.value 为空
	err = db.Put(utils.GetTestKey(22), nil)
	assert.Nil(t, err)
	val3, err := db.Get(utils.GetTestKey(22))
	assert.Equal(t, 0, len(val3))
	assert.Nil(t, err)

	// 6.重启后再 Put 数据
	err = db.Close()
	assert.Nil(t, err)

	// 重启数据库
	db2, err := bitcaskdb.Open(opts)
	assert.Nil(t, err)
	assert.NotNil(t, db2)
	val4 := utils.RandomValue(128)
	err = db2.Put(utils.GetTestKey(55), val4)
	assert.Nil(t, err)
	val5, err := db2.Get(utils.GetTestKey(55))
	assert.Nil(t, err)
	assert.Equal(t, val4, val5)

}

func TestBitcaskDB_Get(t *testing.T) {
	opts := bitcaskdb.DefaultOptions
	opts.DirPath = "./tmp"
	db, err := bitcaskdb.Open(opts)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	// 1.正常读取一条数据
	err = db.Put(utils.GetTestKey(11), utils.RandomValue(24))
	assert.Nil(t, err)
	val1, err := db.Get(utils.GetTestKey(11))
	assert.Nil(t, err)
	assert.NotNil(t, val1)

	// 2.读取一个不存在的 key
	val2, err := db.Get([]byte("some key unknown"))
	assert.Nil(t, val2)
	assert.Equal(t, bitcaskdb.ErrKeyNotFound, err)

	// 3.值被重复 Put 后在读取
	err = db.Put(utils.GetTestKey(22), utils.RandomValue(24))
	assert.Nil(t, err)
	err = db.Put(utils.GetTestKey(22), utils.RandomValue(24))
	val3, err := db.Get(utils.GetTestKey(22))
	assert.Nil(t, err)
	assert.NotNil(t, val3)

	// 4.值被删除后再 Get
	err = db.Put(utils.GetTestKey(33), utils.RandomValue(24))
	assert.Nil(t, err)
	err = db.Delete(utils.GetTestKey(33))
	assert.Nil(t, err)
	val4, err := db.Get(utils.GetTestKey(33))
	assert.Equal(t, 0, len(val4))
	assert.Equal(t, bitcaskdb.ErrKeyNotFound, err)

	// 5.重启后，前面写入的数据都能拿到
	err = db.Close()
	assert.Nil(t, err)

	// 重启数据库
	db2, err := bitcaskdb.Open(opts)
	val6, err := db2.Get(utils.GetTestKey(11))
	assert.Nil(t, err)
	assert.NotNil(t, val6)
	assert.Equal(t, val1, val6)

	val7, err := db2.Get(utils.GetTestKey(22))
	assert.Nil(t, err)
	assert.NotNil(t, val7)
	assert.Equal(t, val3, val7)

	val8, err := db2.Get(utils.GetTestKey(33))
	assert.Equal(t, 0, len(val8))
	assert.Equal(t, bitcaskdb.ErrKeyNotFound, err)
}

func TestBitcaskDB_Delete(t *testing.T) {
	opts := bitcaskdb.DefaultOptions
	opts.DirPath = "./tmp"
	err := os.Mkdir(opts.DirPath, os.ModePerm)
	assert.Nil(t, err)
	db, err := bitcaskdb.Open(opts)
	defer destroyDB(db, opts.DirPath)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	// 1.正常删除一个存在的 key
	err = db.Put(utils.GetTestKey(11), utils.RandomValue(128))
	assert.Nil(t, err)
	err = db.Delete(utils.GetTestKey(11))
	assert.Nil(t, err)
	_, err = db.Get(utils.GetTestKey(11))
	assert.Equal(t, bitcaskdb.ErrKeyNotFound, err)

	// 2.删除一个不存在的 key
	err = db.Delete([]byte("unknown key"))
	assert.Nil(t, err)

	// 3.删除一个空的 key
	err = db.Delete(nil)
	assert.Equal(t, bitcaskdb.ErrKeyIsEmpty, err)

	// 4.值被删除之后重新 Put
	err = db.Put(utils.GetTestKey(22), utils.RandomValue(128))
	assert.Nil(t, err)
	err = db.Delete(utils.GetTestKey(22))
	assert.Nil(t, err)

	err = db.Put(utils.GetTestKey(22), utils.RandomValue(128))
	assert.Nil(t, err)
	val1, err := db.Get(utils.GetTestKey(22))
	assert.NotNil(t, val1)
	assert.Nil(t, err)

	// 5.重启之后，再进行校验
	err = db.Close()
	assert.Nil(t, err)

	// 重启数据库
	db2, err := bitcaskdb.Open(opts)
	_, err = db2.Get(utils.GetTestKey(11))
	assert.Equal(t, bitcaskdb.ErrKeyNotFound, err)

	val2, err := db2.Get(utils.GetTestKey(22))
	assert.Nil(t, err)
	assert.Equal(t, val1, val2)
}

func TestBitcaskDB_Sync(t *testing.T) {
	opts := bitcaskdb.DefaultOptions
	opts.DirPath = "./tmp"
	err := os.Mkdir(opts.DirPath, os.ModePerm)
	assert.Nil(t, err)
	db, err := bitcaskdb.Open(opts)
	defer destroyDB(db, opts.DirPath)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.Put(utils.GetTestKey(11), utils.RandomValue(20))
	assert.Nil(t, err)

	err = db.Sync()
	assert.Nil(t, err)
}

func TestDB_Backup(t *testing.T) {
	opts := bitcaskdb.DefaultOptions
	opts.DirPath = "./tmp"
	err := os.Mkdir(opts.DirPath, os.ModePerm)
	assert.Nil(t, err)
	db, err := bitcaskdb.Open(opts)
	defer destroyDB(db, opts.DirPath)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	for i := 1; i < 1000000; i++ {
		err := db.Put(utils.GetTestKey(i), utils.RandomValue(128))
		assert.Nil(t, err)
	}

	backupDir, _ := os.MkdirTemp("", "bitcask-go-backup-test")
	err = db.Backup(backupDir)
	assert.Nil(t, err)

	opts1 := bitcaskdb.DefaultOptions
	opts1.DirPath = backupDir
	db2, err := bitcaskdb.Open(opts1)
	defer destroyDB(db2, backupDir)
	assert.Nil(t, err)
	assert.NotNil(t, db2)
}
