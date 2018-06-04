package main

import (
	"fmt"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Optimizer"
	"github.com/xitongsys/guery/pb"
)

func main() {
	sql := `select INT64 from test.test.test as t1 join test.test.test as t2 on t1.ID=t2.ID and t1.INT64=t2.INT64 group by ID`

	fmt.Println(sql)

	logicalPlanTree, err := Optimizer.CreateLogicalTree(sql)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	freeExecutor := []pb.Location{}
	for i := 0; i < 100; i++ {
		freeExecutor = append(freeExecutor, pb.Location{Name: fmt.Sprintf("%v", i)})
	}
	res := []EPlan.ENode{}

	EPlan.CreateEPlan(logicalPlanTree, &res, &freeExecutor, 2)
	fmt.Println(logicalPlanTree)
	fmt.Println("--------------")
	fmt.Println(logicalPlanTree.GetMetadata())

}
