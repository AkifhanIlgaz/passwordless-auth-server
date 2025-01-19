package helper

func CreateRedisKey(collection, key string) string {
	return collection + ":" + key
}
