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
	NodeStyle := map[string]string{
		"SCAN":                 "stroke-width:0; fill:rgb(187,255,255);",
		"SELECT":               "stroke-width:0; fill:rgb(152,251,152);",
		"GROUP BY":             "stroke-width:0; fill:rgb(255,255,0);",
		"GROUP BY LOCAL":       "stroke-width:0; fill:rgb(205,205,0);",
		"FILTER":               "stroke-width:0; fill:rgb(0,255,255);",
		"UNION":                "stroke-width:0; fill:rgb(255,106,106);",
		"LIMIT":                "stroke-width:0; fill:rgb(135,206,250);",
		"ORDER BY":             "stroke-width:0; fill:rgb(0,191,255);",
		"ORDER BY LOCAL":       "stroke-width:0; fill:rgb(0,154,205);",
		"JOIN":                 "stroke-width:0; fill:rgb(224,102,255);",
		"HASH JOIN":            "stroke-width:0; fill:rgb(180,82,205);",
		"HASH JOIN SHUFFLE":    "stroke-width:0; fill:rgb(180,62,105);",
		"HAVING":               "stroke-width:0; fill:rgb(191,239,255);",
		"COMBINE":              "stroke-width:0; fill:rgb(255,114,86);",
		"DUPLICATE":            "stroke-width:0; fill:rgb(244,164,96);",
		"AGGREGATE":            "stroke-width:0; fill:rgb(255,48,48);",
		"AGGREGATE FUNC LOCAL": "stroke-width:0; fill:rgb(155,48,48);",
		"SHOW":                 "stroke-width:0; fill:rgb(125,148,148);",
		"UNKNOWN":              "stroke-width:0; fill:rgb(181,181,181);",
	}
	style := NodeStyle[node.NodeType]
	//canvas.Circle(node.X, node.Y, NODER, "stroke-width:0; fill:rgb(0,200,255);")
	canvas.Circle(node.X, node.Y, NODER, style)
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
			rec[node.Executor] = true
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
	BeginTime  string
	EndTime    string
	ErrInfo    string
	Executors  []string
	Progress   int32
}

func NewUITaskInfoFromTask(task *Scheduler.Task) *UITaskInfo {
	res := &UITaskInfo{
		TaskId:     task.TaskId,
		Status:     task.Status.String(),
		Query:      task.Query,
		PlanTree:   DrawSVG(CreateSVGNode(task.EPlanNodes), 850),
		Priority:   task.Runtime.Priority,
		CommitTime: task.CommitTime.Format("2006-01-02 15:04:05"),
		BeginTime:  task.BeginTime.Format("2006-01-02 15:04:05"),
		EndTime:    task.EndTime.Format("2006-01-02 15:04:05"),
		ErrInfo:    fmt.Sprintf("%v", task.Errs),
		Executors:  task.Executors,
		Progress:   0,
	}
	return res
}

func (self *Master) GetUITaskInfos(exeInfos []*UIAgentInfo) map[string][]*UITaskInfo {
	res := make(map[string][]*UITaskInfo)
	exeInfoMap := map[string]string{}
	for _, exeInfo := range exeInfos {
		exeInfoMap[exeInfo.Name] = exeInfo.Status
	}

	res["TODO"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Todos {
		res["TODO"] = append(res["TODO"], NewUITaskInfoFromTask(t))
	}

	res["DOING"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Doings {
		tinfo := NewUITaskInfoFromTask(t)
		freeNum := 0
		for _, name := range tinfo.Executors {
			if exeInfoMap[name] == "FREE" {
				freeNum++
			}
		}
		if len(tinfo.Executors) > 0 {
			tinfo.Progress = int32(freeNum * 100 / len(tinfo.Executors))
		}
		res["DOING"] = append(res["DOING"], tinfo)
	}

	res["DONE"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Dones {
		tinfo := NewUITaskInfoFromTask(t)
		tinfo.Progress = 100
		res["DONE"] = append(res["DONE"], tinfo)
	}

	res["FAILED"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Fails {
		tinfo := NewUITaskInfoFromTask(t)
		tinfo.Progress = 100
		res["FAILED"] = append(res["FAILED"], tinfo)
	}

	return res
}
