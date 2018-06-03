#
# FFS Makefile
#

ffs:
	go get github.com/go-ini/ini
	go get github.com/gorilla/securecookie
	go get github.com/gorilla/websocket
	go get github.com/lib/pq
	go get gopkg.in/redis.v4

	go build

clean:
	rm ffs