package redis

import (
	"github.com/architsmat38/gorush/config"
	"gopkg.in/redis.v4"
	"log"
	"strconv"
)

// Stat variable for redis
const (
	TotalCountKey     = "gorush-total-count"
	IosSuccessKey     = "gorush-ios-success-count"
	IosErrorKey       = "gorush-ios-error-count"
	AndroidSuccessKey = "gorush-android-success-count"
	AndroidErrorKey   = "gorush-android-error-count"
)

//
var redisClient *redis.Client

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

func getInt64(key string, count *int64) {
	val, _ := redisClient.Get(key).Result()
	*count, _ = strconv.ParseInt(val, 10, 64)
}

// Storage is interface structure
type Storage struct {
	config config.ConfYaml
}

// Init client storage.
func (s *Storage) Init() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     s.config.Stat.Redis.Addr,
		Password: s.config.Stat.Redis.Password,
		DB:       s.config.Stat.Redis.DB,
	})

	_, err := redisClient.Ping().Result()

	if err != nil {
		// redis server error
		log.Println("Can't connect redis server: " + err.Error())

		return err
	}

	return nil
}

// Reset Client storage.
func (s *Storage) Reset() {
	redisClient.Set(TotalCountKey, strconv.Itoa(0), 0)
	redisClient.Set(IosSuccessKey, strconv.Itoa(0), 0)
	redisClient.Set(IosErrorKey, strconv.Itoa(0), 0)
	redisClient.Set(AndroidSuccessKey, strconv.Itoa(0), 0)
	redisClient.Set(AndroidErrorKey, strconv.Itoa(0), 0)
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	redisClient.Set(TotalCountKey, strconv.Itoa(int(total)), 0)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	redisClient.Set(IosSuccessKey, strconv.Itoa(int(total)), 0)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	redisClient.Set(IosErrorKey, strconv.Itoa(int(total)), 0)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	redisClient.Set(AndroidSuccessKey, strconv.Itoa(int(total)), 0)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	redisClient.Set(AndroidErrorKey, strconv.Itoa(int(total)), 0)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	getInt64(TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	getInt64(IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	getInt64(IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	getInt64(AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	getInt64(AndroidErrorKey, &count)

	return count
}
