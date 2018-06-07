var Tasks = {}

function InfoTasksToTable(infos) {
    var prefix=' \
<table class="table table-dark"> \
  <thead> \
    <tr> \
      <th scope="col">ID</th> \
      <th scope="col">Status</th> \
      <th scope="col">Query</th> \
      <th scope="col">Priority</th> \
      <th scope="col">CommitTime</th> \
      <th scope="col">Err</th> \
      <th scope="col">Control</th> \
    </tr> \
  </thead> \
  <tbody>';

    var suffix = ' \
  </tbody> \
</table>';

    var content='';
    infos.sort(function(a,b){return b.TaskId - a.TaskId;});
    for(var i=0; i<infos.length; i++){
	Tasks[infos[i].TaskId] = infos[i]
	
	rec='<tr>';
	if(infos[i].Status=="DONE"){
	    rec='<tr class="active">'
	}else if (infos[i].Status=="DOING"){
	    rec='<tr class="success">'
	}else if (infos[i].Status=="TODO") {
	    rec='<tr class="info">'
	}else if (infos[i].Status=="FAILED"){
	    rec='<tr class="danger">'
	}
	
	rec=rec + '<td>' + infos[i].TaskId + '</td>';
	rec=rec + '<td>' + infos[i].Status + '</td>';
	rec=rec + '<td>' + infos[i].Query + ' </td>';
	rec=rec + '<td>' + infos[i].Priority + '</td>';
	rec=rec + '<td>' + infos[i].CommitTime + '</td>';
	rec=rec + '<td>' + infos[i].ErrInfo + '</td>';
	
	rec = rec + '<td>' + '<div class="btn-group" role="group">';
	rec = rec + '<button type="button"  class="btn btn-info" onclick=\'ShowDetail("' + infos[i].TaskId + '")\'>Detail</button>';
	if(infos[i].Status=="DOING" || infos[i].Status=="TODO"){
	    rec = rec + '<button type="button" class="btn btn-danger" onclick=\'CancelTask("' + infos[i].TaskId + '")\'>Cancel</button>' 
	}
	red = rec + '</div>' + '</td>';
	
	rec=rec+'</tr>';
	content = content + rec
    }
    return prefix + content + suffix;
}

function CancelTask(taskid) {
    $.post('control', {'cmd':'canceltask', 'taskid':taskid},
	   function(){});
}

function PlanTreeToNodeStructure(pnode) {
    if(pnode == undefined || pnode==null){
	return {};
    }
    pnode.name=pnode.NodeType;
    pnode.executor="Executor:" + pnode.Executor;
    pnode.location="Location:" + pnode.Location;
   
    if(pnode.Inputs != null){
	pnode.children=new Array();
	for(var i=0; i<pnode.Inputs.length; i++){
	    pnode.children.push(PlanTreeToNodeStructure(pnode.Inputs[i]));
	}
    }
    return pnode
}

function ShowDetail(taskId) {
    var taskInfo = Tasks[taskId];
    var plan=PlanTreeToNodeStructure(taskInfo.PlanTree);
    update(plan);
    $('#taskDetailDialog').modal('show');
}

// ************** Generate the tree diagram	 *****************
var margin = {top: 20, right: 120, bottom: 20, left: 120},
    width = 960 - margin.right - margin.left,
    height = 500 - margin.top - margin.bottom;

var i = 0;
var tree = d3.layout.tree()
    .size([height, width]);
var diagonal = d3.svg.diagonal()
    .projection(function(d) { return [d.x, d.y]; });
var svg = d3.select("svg")
    .attr("width", width + margin.right + margin.left)
    .attr("height", height + margin.top + margin.bottom)
    .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

function update(source) {
    // Compute the new tree layout.
    var nodes = tree.nodes(source).reverse(),
	links = tree.links(nodes);

    // Normalize for fixed-depth.
    nodes.forEach(function(d) { d.y = d.depth * 50;});

    // Declare the nodes…
    var node = svg.selectAll("g.node")
	.data(nodes, function(d) { return d.id || (d.id = ++i); });

    // Enter the nodes.
    var nodeEnter = node.enter().append("g")
	.attr("class", "node")
	.attr("transform", function(d) { 
	    return "translate(" + d.x + "," + d.y + ")"; });

    nodeEnter.append("circle")
	.attr("r", 10)
	.style("fill", "#fff");

    nodeEnter.append("text")
	.attr("y", function(d) { 
	    return d.children || d._children ? -13 : 13; })
	.attr("dy", ".35em")
	.attr("text-anchor", function(d) { 
	    return d.children || d._children ? "end" : "start"; })
	.text(function(d) { return d.name; })
	.style("fill-opacity", 1);

    // Declare the links…
    var link = svg.selectAll("path.link")
	.data(links, function(d) { return d.target.id; });

    // Enter the links.
    link.enter().insert("path", "g")
	.attr("class", "link")
	.attr("d", diagonal);

}
