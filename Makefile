EN=30
build:prepare
	cd build; go build ../Main/guery.go; cp ../Config/config.json ./

run:stop build
	cd build; ./guery master --address 127.0.0.1:1111 --config config.json >> m.log &
	cd build; \
	for i in `seq 1 $(EN)`;do \
		./guery executor --master 127.0.0.1:1111 --config config.json >> e.log &	\
	done

stop:
	-killall guery

prepare:clean
	mkdir build
	cp -rvf Master/UI ./build/

clean:stop
	-rm -rvf ./build
