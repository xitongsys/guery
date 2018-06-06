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
	    rec=rec + '<td>' + infos[i].Query + '</td>';
	    rec=rec + '<td>' + infos[i].Priority + '</td>';
	    rec=rec + '<td>' + infos[i].CommitTime + '</td>';
	    rec=rec + '<td>' + infos[i].ErrInfo + '</td>';
	    
	    if(infos[i].Status=="DOING" || infos[i].Status=="TODO"){
		rec = rec + '<td>' + '<div class="btn-group" role="group">' +
		    '<button type="button" class="btn btn-danger">Cancel</button>' +
		    '</div>' + '</td>';
	    }
	    
	    rec=rec+'</tr>';
	    content = content + rec
	}
	return prefix + content + suffix;
}
