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
	    rec=rec + '<td align="right"><b>Name</b></td>';
	    rec=rec + '<td>' + infos[i].Name + '</td>';
	    rec=rec + '</tr>';

	    rec=rec+'<tr>'
	    rec=rec + '<td align="right"><b>Address</b></td>';
	    rec=rec + '<td>' + infos[i].Address + '</td>';
	    rec=rec + '</tr>';

	    rec=rec+'<tr>'
	    rec=rec + '<td align="right"><b>CPU</b></td>';
	    rec=rec + '<td>';
	    for(var j=0; j<infos[i].CpuUsage.length; j++){
		rec = rec + '<span class="label label-success">'
		rec = rec + FloatFormat(infos[i].CpuUsage[j], 2) + "%"
		rec = rec + '</span> '
	    }
	    rec=rec + '</td>';
	    rec=rec + '</tr>';

	    
	    var totMem=FloatFormat(infos[i].TotalMemory/1024.0/1024.0, 2);
	    var freeMem=FloatFormat(infos[i].FreeMemory/1024.0/1024.0, 2);
	    var usedMem=FloatFormat(totMem-freeMem, 2);
	    rec=rec+'<tr>'
	    rec=rec + '<td align="right"><b>Memory</b></td>';
	    rec=rec + '<td>';
	    rec=rec + '<div class="progress">';
	    rec=rec + '<div class="progress-bar" role="progressbar" aria-valuemin="0" aria-valuemax="100" style="width:' + usedMem*100/totMem + '%;">';
	    rec=rec + FloatFormat(usedMem*100/totMem,2) + '% ' + '(' + usedMem + '/' + totMem + 'MB)';
	    rec=rec + '</div></div>'
	    rec=rec + '</td>';

	    
	    //rec=rec + '<td>' + 'Total: ' + FloatFormat(infos[i].TotalMemory/1024.0/1024.0, 2) + 'MB'  + ' Free: ' + FloatFormat(infos[i].FreeMemory/1024.0/1024.0, 2) + 'MB'  + '</td>';
	    rec=rec + '</tr>';

	    rec=rec+'<tr>'
	    rec=rec + '<td align="right"><b>Executor</b></td>';
	    rec=rec + '<td>';
	    for(var j=0; j<infos[i].ExecutorNumber; j++){
		rec=rec + '<span class="label label-danger"> </span> ' ;
	    }
	    for(var j=0; j<infos[i].MaxExecutorNumber-infos[i].ExecutorNumber; j++){
		rec=rec + '<span class="label label-success"> </span> ';
	    }
	    rec=rec+'</td>';
	    //rec=rec + '<td>' + 'Max: ' + infos[i].MaxExecutorNumber + ' Running: ' + infos[i].ExecutorNumber + '</td>';
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

function FloatFormat(num, bit){
    return Math.round(num * Math.pow(10, bit))/Math.pow(10,bit)
}
