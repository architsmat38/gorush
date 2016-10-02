package leveldb

import (
	c "github.com/architsmat38/gorush/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLevelDBEngine(t *testing.T) {
	var val int64

	config := c.BuildDefaultPushConf()

	if _, err := os.Stat(config.Stat.LevelDB.Path); os.IsNotExist(err) {
		os.RemoveAll(config.Stat.LevelDB.Path)
	}

	levelDB := New(config)
	levelDB.Init()
	levelDB.Reset()

	levelDB.AddTotalCount(10)
	val = levelDB.GetTotalCount()
	assert.Equal(t, int64(10), val)
	levelDB.AddTotalCount(10)
	val = levelDB.GetTotalCount()
	assert.Equal(t, int64(20), val)

	levelDB.AddIosSuccess(20)
	val = levelDB.GetIosSuccess()
	assert.Equal(t, int64(20), val)

	levelDB.AddIosError(30)
	val = levelDB.GetIosError()
	assert.Equal(t, int64(30), val)

	levelDB.AddAndroidSuccess(40)
	val = levelDB.GetAndroidSuccess()
	assert.Equal(t, int64(40), val)

	levelDB.AddAndroidError(50)
	val = levelDB.GetAndroidError()
	assert.Equal(t, int64(50), val)

	levelDB.Reset()
	val = levelDB.GetAndroidError()
	assert.Equal(t, int64(0), val)
}
