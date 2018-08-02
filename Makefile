build:prepare
	cd build; go build ../main/guery.go; cp ../config/config.json ./; cp -rf ../test/* /tmp/

run:stop build
	cd build; ./guery master --address 127.0.0.1:1111 --config ./config.json >> m.log &
	cd build; ./guery agent --master 127.0.0.1:1111 --config ./config.json >> a1.log &	
	cd build; ./guery agent --master 127.0.0.1:1111 --config ./config.json >> a2.log &	


stop:
	-killall guery

prepare:clean
	mkdir build
	cp -rvf master/ui ./build/

clean:stop
	-rm -rvf ./build
