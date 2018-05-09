build:prepare
	cd build; go build ../Main/guery.go

run:stop build
	cd build; ./guery master --address 127.0.0.1:1111 >& m.log &
	cd build; ./guery executor --master 127.0.0.1:1111 >e1.log &
	cd build; ./guery executor --master 127.0.0.1:1111 >e2.log &
	cd build; ./guery executor --master 127.0.0.1:1111 >e3.log &

stop:
	-killall guery

prepare:clean
	mkdir build
	cp -rvf Master/UI ./build/

clean:stop
	-rm -rvf ./build
