function InfoExecutorsToTable(infos) {
	var prefix=' \
<table class="table" style="margin:0px; padding:0px;"> \
  <thead> \
    <tr> \
      <th scope="col">Name</th> \
      <th scope="col">Status</th> \
      <th scope="col">Location</th> \
      <th scope="col">Task</th> \
      <th scope="col">Control</th> \
    </tr> \
  </thead> \
  <tbody>';

	var suffix = ' \
  </tbody> \
</table>';

	infos.sort(function(a,b){
		if(a.Name > b.Name){return 1;}
		else if(a.Name==b.Name){return 0;}
		else{return -1;}
	});

	var content="";
	for(var i=0; i<infos.length; i++){
	    rec='<tr>';
	    if(infos[i].Status=="FREE"){
		rec='<tr class="success">';
	    }else if(infos[i].Status=="BUSY"){
		rec='<tr class="danger">';
	    }
		
	    rec=rec + '<td>' + infos[i].Name + '</td>';
	    rec=rec + '<td>' + infos[i].Status + '</td>';
	    rec=rec + '<td>' + infos[i].Location + '</td>';
	    rec=rec + '<td>' + infos[i].TaskId + '</td>';
	    rec=rec + '<td>' + '<div class="btn-group" role="group">' +
		'<button type="button" class="btn btn-success" onclick=\'DuplicateExecutor("' + infos[i].Name + '")\'>Duplicate</button>' +
		'<button type="button" class="btn btn-primary" onclick=\'RestartExecutor("' + infos[i].Name + '")\'>Restart</button>' +
		'<button type="button" class="btn btn-danger" onclick=\'KillExecutor("' + infos[i].Name + '")\'>Kill</button>' +		
		'</div>' + '</td>';
	    
	    rec=rec+'</tr>';
	    content = content + rec;
	}
	return prefix + content + suffix;
}

function KillExecutor(name){
    $.post('control', {'cmd':'killexecutor', 'name':name},
	   function(){});
}

function DuplicateExecutor(name){
    $.post('control', {'cmd':'duplicateexecutor', 'name':name},
	   function(){});
}
function RestartExecutor(name){
    $.post('control', {'cmd':'restartexecutor', 'name':name},
	   function(){});
}

