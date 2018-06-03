function InfoExecutorsToTable(infos) {
	var prefix=' \
<table class="table table-dark"> \
  <thead> \
    <tr> \
      <th scope="col">Name</th> \
      <th scope="col">Status</th> \
      <th scope="col">Location</th> \
      <th scope="col">Task</th> \
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
		if(infos[i].Status=="Idle"){
			rec='<tr class="success">';
		}else if(infos[i].Status=="Busy"){
			rec='<tr class="warning">';
		}
		
		rec=rec + '<td>' + infos[i].Name + '</td>';
		rec=rec + '<td>' + infos[i].Status + '</td>';
		rec=rec + '<td>' + infos[i].Location + '</td>';
		rec=rec + '<td>' + infos[i].TaskId + '</td>';
		rec=rec+'</tr>';
		content = content + rec;
	}
	return prefix + content + suffix;
}
