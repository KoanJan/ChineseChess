-- 'unlock'
-- this script is used by copying as a string into the target .go file

if redis.call("get", KEYS[1]) == ARGV[1]
then
    return redis.call("del", KEYS[1])
else
    return 0
end
