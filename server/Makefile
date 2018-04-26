
all:
	go build *.go
test:
	go run *.go -clientID=5001 -cassIP=127.0.0.2

run1:
	go run *.go -clientID=5001 -cassIP=127.0.0.2 &
run2:
	go run *.go -clientID=5002 -cassIP=127.0.0.3 &
run3:
	go run *.go -clientID=5003 -cassIP=127.0.0.4 &

clear:
	rm cassConn
 
