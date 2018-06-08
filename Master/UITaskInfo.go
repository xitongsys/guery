package Master

import (
	"bytes"
	"fmt"
	"math"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Scheduler"
	"github.com/xitongsys/svgo"
)

var (
	NODER int = 10
	ROWH  int = 100
)

type SVGNode struct {
	Inputs   []*SVGNode
	Outputs  []*SVGNode
	NodeType string
	Location string
	Executor string
	X, Y     int
}

func NewSVGNodeFromENode(node EPlan.ENode) *SVGNode {
	loc := node.GetLocation()
	res := &SVGNode{
		NodeType: node.GetNodeType().String(),
		Location: (&loc).GetURL(),
		Executor: (&loc).Name,
	}
	return res

}

func CreateSVGNode(nodes []EPlan.ENode) *SVGNode {
	if len(nodes) <= 0 {
		return nil
	}
	svgNodes := []*SVGNode{}
	nodeMap := map[string]int{}

	for i, node := range nodes {
		svgNode := NewSVGNodeFromENode(node)
		svgNodes = append(svgNodes, svgNode)
		nodeMap[svgNode.Executor] = i
	}

	for i, node := range nodes {
		for _, loc := range node.GetInputs() {
			ad := loc.Name
			j := nodeMap[ad]
			svgNodes[i].Inputs = append(svgNodes[i].Inputs, svgNodes[j])
		}
		for _, loc := range node.GetOutputs() {
			ad := loc.Name
			j := nodeMap[ad]
			svgNodes[i].Outputs = append(svgNodes[i].Outputs, svgNodes[j])
		}
	}
	res := svgNodes[len(svgNodes)-1]
	return res
}

func SetSVGNodePos(node *SVGNode, tW int) int {
	if node == nil {
		return 0
	}
	rec := map[string]bool{}
	q := []*SVGNode{node}
	depth := 0
	for len(q) > 0 {
		ln := len(q)
		curNodes := []*SVGNode{}
		for i := 0; i < ln; i++ {
			node := q[i]
			q = append(q, node.Inputs...)
			if _, ok := rec[node.Executor]; !ok {
				rec[node.Executor] = true
				curNodes = append(curNodes, node)
			}
		}
		nnum := len(curNodes)
		for i, node := range curNodes {
			node.X = (tW/nnum)*i + (tW/nnum)/2
			node.Y = depth*ROWH + 20
		}
		depth++
		q = q[ln:]
	}
	return depth
}

func DrawNode(canvas *svg.SVG, node *SVGNode) {
	canvas.Circle(node.X, node.Y, NODER, "stroke:rgb(0,200,255);stroke-width:0; fill:rgb(0,200,255);")
	canvas.TextWithTitle(node.X+NODER, node.Y, node.NodeType, fmt.Sprintf("Executor: %s\n URL: %s", node.Executor, node.Location), "font-size:12pt;")
}

func DrawArrow(canvas *svg.SVG, nodeFrom *SVGNode, nodeTo *SVGNode) {
	x1, y1 := nodeFrom.X, nodeFrom.Y-NODER/2
	x2, y2 := nodeTo.X, nodeTo.Y+NODER/2
	L := math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
	l := float64(20)
	a := math.Pi / 10
	dx, dy := float64(x1-x2), float64(y1-y2)
	xs := []int{int(x2), int(float64(x2) + l/L*(math.Cos(a)*dx-math.Sin(a)*dy)), int(float64(x2) + l/L*(math.Cos(a)*dx+math.Sin(a)*dy))}
	ys := []int{int(y2), int(float64(y2) + l/L*(math.Sin(a)*dx+math.Cos(a)*dy)), int(float64(y2) + l/L*(-math.Sin(a)*dx+math.Cos(a)*dy))}
	canvas.Polygon(xs, ys, "stroke:rgb(0,0,0); stroke-width:0; fill:rgb(0,0,0);")
	canvas.Line(x1, y1, x2, y2, "stroke:rgb(0,0,0);stroke-width:2")
}

func DrawSVG(node *SVGNode, tW int) string {
	if node == nil {
		return ""
	}
	h := SetSVGNodePos(node, tW) * ROWH

	buf := new(bytes.Buffer)
	canvas := svg.New(buf)
	canvas.Start(900, h)

	rec := map[string]bool{}
	q := []*SVGNode{node}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if _, ok := rec[node.Executor]; !ok {
			for _, input := range node.Inputs {
				DrawArrow(canvas, input, node)
				q = append(q, input)
			}
			DrawNode(canvas, node)
		}
	}
	canvas.End()
	return string(buf.Bytes())
}

///////////////////////////

type UITaskInfo struct {
	TaskId     int64
	Status     string
	Query      string
	PlanTree   string
	Priority   int32
	CommitTime string
	ErrInfo    string
}

func NewUITaskInfoFromTask(task *Scheduler.Task) *UITaskInfo {
	res := &UITaskInfo{
		TaskId:     task.TaskId,
		Status:     task.Status.String(),
		Query:      task.Query,
		PlanTree:   DrawSVG(CreateSVGNode(task.EPlanNodes), 850),
		Priority:   task.Priority,
		CommitTime: task.CommitTime.Format("2006-01-02 15:04:05"),
		ErrInfo:    fmt.Sprintf("%v", task.Errs),
	}
	return res
}

func (self *Master) GetUITaskInfos() map[string][]*UITaskInfo {
	res := make(map[string][]*UITaskInfo)

	res["TODO"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Todos {
		res["TODO"] = append(res["TODO"], NewUITaskInfoFromTask(t))
	}

	res["DOING"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Doings {
		res["DOING"] = append(res["DOING"], NewUITaskInfoFromTask(t))
	}

	res["DONE"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Dones {
		res["DONE"] = append(res["DONE"], NewUITaskInfoFromTask(t))
	}

	res["FAILED"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Fails {
		res["FAILED"] = append(res["FAILED"], NewUITaskInfoFromTask(t))
	}

	return res
}
