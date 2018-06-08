var Tasks = {}

function InfoTasksToTable(infos) {
    var prefix=' \
<table class="table table-dark"> \
  <thead> \
    <tr> \
      <th scope="col">Task</th> \
      <th scope="col">Query</th> \
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

	rec=rec + '<td style="width:25%;">';
	
	rec=rec + '<ul class="list-group" style="font-size:8pt;">';
	rec=rec + '<li class="list-group-item list-group-item-info">ID: ' + infos[i].TaskId + '</li>';
	rec=rec + '<li class="list-group-item list-group-item-info">Status: ' + infos[i].Status + '</li>';
	rec=rec + '<li class="list-group-item list-group-item-info">Priority: ' + infos[i].Priority + '</li>';
	rec=rec + '<li class="list-group-item list-group-item-info">Begin: ' + infos[i].BeginTime + '</li>';
	rec=rec + '<li class="list-group-item list-group-item-info">End: ' + infos[i].EndTime + '</li>';
	rec=rec + '<li class="list-group-item list-group-item-info">Commit: ' + infos[i].CommitTime + '</li>';
	rec=rec + "</ul>";

	rec=rec + '<div align="right">';
	rec=rec + '<div class="btn-group" role="group">';
	rec = rec + '<button type="button" class="btn btn-info" onclick=\'ShowDetail("' + infos[i].TaskId + '")\'>Details</button>';
	if(infos[i].Status=="DOING" || infos[i].Status=="TODO"){
	    rec = rec + '<button type="button" class="btn btn-danger" onclick=\'CancelTask("' + infos[i].TaskId + '")\'>Cancel</button>' 
	}
	rec=rec + "</div>";
	rec=rec + "</div>";
	
	rec=rec + '</td>';
	


	//query
	var progressBar =  '<div class="progress">';
	rec=rec + '<td>' + '<div class="alert alert-info">' + infos[i].Query + "</div>";
	if(infos[i].Status=="FAILED"){
	    progressBar = progressBar + '<div class="progress-bar progress-bar-danger progress-bar-striped" role="progressbar" aria-valuenow="' + infos[i].Progress + '" aria-valuemin="0" aria-valuemax="100" style="width: ' + infos[i].Progress + '%;">'
	}else if (infos[i].Status=="DOING"){
	    progressBar = progressBar + '<div class="progress-bar progress-bar-success progress-bar-striped active" role="progressbar" aria-valuenow="' + infos[i].Progress + '" aria-valuemin="0" aria-valuemax="100" style="width: ' + infos[i].Progress + '%;">'
	}else{
	    progressBar = progressBar + '<div class="progress-bar progress-bar-success progress-bar-striped" role="progressbar" aria-valuenow="' + infos[i].Progress + '" aria-valuemin="0" aria-valuemax="100" style="width: ' + infos[i].Progress + '%;">'	    
	}
        progressBar = progressBar + infos[i].Progress + '%' + '</div></div>';

	
	rec = rec + progressBar;
	rec = rec + '</td>';
	
	rec=rec+'</tr>';
	content = content + rec
    }
    return prefix + content + suffix;
}

function CancelTask(taskid) {
    $.post('control', {'cmd':'canceltask', 'taskid':taskid},
	   function(){});
}

function ShowDetail(taskId) {
    var taskInfo = Tasks[taskId];
    $("#sqlDiv").html(taskInfo.Query);
    $("#planTreeDiv").html(taskInfo.PlanTree);
    $('#taskDetailDialog').modal('show');
}

