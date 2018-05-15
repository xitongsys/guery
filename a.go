package main

import (
	"fmt"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Optimizer"
	"github.com/xitongsys/guery/pb"
)

func main() {
	sql := `SELECT CA.ID FROM CLASS AS CA`

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
