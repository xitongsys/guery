# Guery v0.1
Guery is a pure-go implementation of distributed SQL query engine for big data (like Presto). It is composed of one master and many agents and supports to query vast amounts of data using distributed queries.

## Status
This project is still a work in progress.
### Todo list
* Support more SQL statement
* Improve performance
* Improve error control
* Add ut and regression test
* Add Wiki
* ...

### NEED YOUR HELP !!!
This project is in progress and a lot of work need to do. If you are interested in it, please give me your hands and anything is welcome! I also list some issues.
If you want to improve something, please open an issue to let others know :)
* [slack](https://join.slack.com/t/guery-group/shared_invite/enQtNDA1MjM0MTA4OTYzLTljYjlmZjNkZTdiMWQ1YmNkMWNlOGQwMTA1YTg0ZTk1MTNhOWUxNzc1N2Y1MmQ4MThhNDMyZDliZWNmOTY4OGI)
* [google-group](https://groups.google.com/forum/#!forum/guery-group)

## Supported
* Datasource: hive
* FileSystem: local/hdfs/s3
* FileType: csv/parquet/orc
* DataType: bool/int32/int64/float32/float64/string/date/timestamp

## Web UI
Web UI is the web interface of a Guery cluster to monitor and inspect the task executions in a web browser.
It provides following modules:

Task information/Agent information & management/Execute plan

![ui](https://github.com/xitongsys/guery/blob/master/doc/images/ui.png)

## Deploy
* build guery `cd $GOPATH/src/github.com/xitongsys/guery; make build`
* run master `./guery master --address 127.0.0.1:1111 --config ./config.json >> m.log &`
* run agent `./guery agent --master 127.0.0.1:1111 --config ./config.json >> e.log &`
* web ui `http://127.0.0.1:1111`


## Demo
```sh
cd $GOPATH/src/github.com/xitongsys/guery
make run
curl -XPOST -d"sql=select * from test.test.csv where var1=1" 127.0.0.1:1111/query
open web browser: http://127.0.0.1:1111
```








