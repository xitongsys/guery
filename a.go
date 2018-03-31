package main

import (
	"fmt"
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/Plan"
)

func main() {
	t := DataSource.NewTableSource("t", []string{}, []Common.Type{})
	fmt.Println("hehe", t)
}
