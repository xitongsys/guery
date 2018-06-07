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
    pnode.text={
	name: pnode.NodeType,
	data_location: pnode.Location
    };
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
    chart_config = {
	chart:{
	    container: "#PlanTree",
	    scrollbar: "native", 
	    rootOrientation: "NORTH",
	    nodeAlign: "BOTTOM",
	    connectors:{
		type: "step",
		style: {
		    "stroke-width": 2
		}
	    },
	},
	nodeStructure: PlanTreeToNodeStructure(taskInfo.PlanTree)
    };

    var chart = new Treant(chart_config);
    
    $('#taskDetailDialog').modal('show');
}
