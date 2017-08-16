package redis

// Do executes a redis command
func Do(commandName string, args ...interface{}) (interface{}, error) {

	conn := pool.Get()
	return conn.Do(commandName, args...)
}
