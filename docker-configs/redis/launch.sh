cur_path=`pwd`
docker rm redis
docker run --name redis -p 6379:6379 -v $cur_path/redis.conf:/etc/redis/redis.conf -d redis:latest redis-server /etc/redis/redis.conf