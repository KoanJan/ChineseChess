package redis

// Do executes a redis command
func Do(commandName string, args ...interface{}) (interface{}, error) {

	conn := pool.Get()
	return conn.Do(commandName, args...)
}

// Hset executes the command 'HSET' of Redis
func Hset(field, key, v interface{}) (interface{}, error) {

	return Do("hset", field, key, v)
}

// Hget executes the command 'HGET' of Redis
func Hget(field, key interface{}) (interface{}, error) {

	return Do("hget", field, key)
}

// Set executes the command 'SET' of Redis
func Set(key, v interface{}) (interface{}, error) {

	return Do("set", key, v)
}

// Get executes the command 'GET' of Redis
func Get(key interface{}) (interface{}, error) {

	return Do("get", key)
}

// Del executes the command 'DEL' of Redis
func Del(key interface{}) (interface{}, error) {

	return Do("del", key)
}

// Expire executes the command 'EXPIRE' of Redis
func Expire(key, seconds interface{}) (interface{}, error) {

	return Do("expire", key, seconds)
}
