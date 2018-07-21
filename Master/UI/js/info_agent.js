function InfoAgentsToTable(infos) {
	var prefix=' \
<table class="table" style="margin:0px; padding:0px;"> \
  <thead> \
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
	    rec='';
	    
	    rec=rec+'<tr>'
	    rec=rec + '<td>Name</td>';
	    rec=rec + '<td>' + infos[i].Name + '</td>';
	    rec=rec + '</tr>';

	    rec=rec+'<tr>'
	    rec=rec + '<td>Address</td>';
	    rec=rec + '<td>' + infos[i].Address + '</td>';
	    rec=rec + '</tr>';

	    rec=rec+'<tr>'
	    rec=rec + '<td>CPU</td>';
	    rec=rec + '<td>' + 'Number: ' + infos[i].CpuNumber + 'Usage: ' + infos[i].CpuUsage + '</td>';
	    rec=rec + '</tr>';
	    
	    rec=rec+'<tr>'
	    rec=rec + '<td>Memory</td>';
	    rec=rec + '<td>' + 'Total: ' + infos[i].TotalMemory + 'Free: ' + infos[i].FreeMemory + '</td>';
	    rec=rec + '</tr>';

	    rec=rec+'<tr>'
	    rec=rec + '<td>Executor</td>';
	    rec=rec + '<td>' + 'Max: ' + infos[i].MaxExecutorNumber + 'Running: ' + infos[i].ExecutorNumber + '</td>';
	    rec=rec + '</tr>';

	    rec=rec+'<tr>'
	    rec=rec + '<td colspan=2 align="right">' + '<div class="btn-group" role="group">' +
		'<button type="button" class="btn btn-success" onclick=\'DuplicateAgent("' + infos[i].Name + '")\'>Duplicate</button>' +
		'<button type="button" class="btn btn-primary" onclick=\'RestartAgent("' + infos[i].Name + '")\'>Restart</button>' +
		'<button type="button" class="btn btn-danger" onclick=\'KillAgent("' + infos[i].Name + '")\'>Kill</button>' +		
		'</div>' + '</td>';
	    
	    rec=rec+'</tr>';
	    
	    content = content + rec;
	}
	return prefix + content + suffix;
}

function KillAgent(name){
    $.post('control', {'cmd':'killagent', 'name':name},
	   function(){});
}

function DuplicateAgent(name){
    $.post('control', {'cmd':'duplicateagent', 'name':name},
	   function(){});
}
function RestartAgent(name){
    $.post('control', {'cmd':'restartagent', 'name':name},
	   function(){});
}

