build:prepare
	cd run
	go build ../Main/guery.go

run:build
	./guery master --address 127.0.0.1:1111
	./guery executor --master 127.0.0.1:1111
	./guery executor --master 127.0.0.1:1111
	./guery executor --master 127.0.0.1:1111

prepare:clean
	mkdir build
	cp -rvf Mater/UI ./build/

clean:
	[ -d build ] && rm -rvf build
