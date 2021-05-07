```sh
docker image ls
redis                    
docker run --rm --name redis3 -p 0.0.0.0:6379:6379 -d redis:latest
docker run --rm -p 0.0.0.0:6379:6379 -d redis:6.0.9
docker ps


go test -bench=. -benchtime=1s
# BenchmarkRedigo-8                                                 734088              1844 ns/op
# BenchmarkGoRedis-8                                                  8282            142040 ns/op
# BenchmarkRadix3-8                                                   8036            149718 ns/op

# BenchmarkSerialGetSet/radix-8               						  4540            235175 ns/op
# BenchmarkSerialGetSet/redigo-8              						  4986            228132 ns/op
# BenchmarkSerialGetSet/redispipe-8            						   427           3110015 ns/op
# BenchmarkSerialGetSet/redispipe_pause0-8                    		  3484            345068 ns/op
# BenchmarkSerialGetSetLargeArgs/radix-8                      		  3358            355671 ns/op
# BenchmarkSerialGetSetLargeArgs/redigo-8                     		  3741            366067 ns/op
# BenchmarkSerialGetSetLargeArgs/redispipe-8                   		   414           3215232 ns/op
# BenchmarkSerialGetSetLargeArgs/redispipe_pause0-8                   3080            495226 ns/op
# BenchmarkParallelGetSet/radix/no_pipeline-8                        31476             37881 ns/op
# BenchmarkParallelGetSet/radix/one_pipeline-8                       23282             48906 ns/op
# BenchmarkParallelGetSet/radix/default-8                            25814             49320 ns/op
# BenchmarkParallelGetSet/redigo-8                                   26710             39576 ns/op
# BenchmarkParallelGetSet/redispipe-8                                25291             53241 ns/op
# PASS
# ok      go-redis-bench  26.668s
```

https://redis.io/clients#go

8 days ago 1863 commits
https://github.com/go-redis/redis

7 days ago 1,071 commits
https://github.com/mediocregopher/radix

on Feb 18 279 commits
https://github.com/gomodule/redigo

11 days ago 228 commits
https://github.com/joomcode/redispipe

 on Mar 21 31 commits
https://github.com/stfnmllr/go-resp3

on Feb 17, 2020  247 commits
https://github.com/insmo/godis

on Feb 11, 2020 112 commits
https://github.com/pascaldekloe/redis

on Aug 8, 2018 51 commits
https://github.com/gistao/RedisGo-Async

on Sep 9, 2017 341 commits
https://github.com/tideland/golib

on Sep 28, 2016 187 commits
https://github.com/xuyu/goredis

on Jul 30, 2016 78 commits
https://github.com/hoisie/redis

on Mar 16, 2016  44 commits
https://github.com/keimoon/gore

on Jun 2, 2015 130 commits
https://github.com/gosexy/redis

on Dec 15, 2014 50 commits
https://github.com/shipwire/redis

155 commits last 2012
https://github.com/alphazero/Go-Redis