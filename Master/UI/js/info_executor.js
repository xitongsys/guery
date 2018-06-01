function InfoExecutorsToTable(infos) {
	var prefix=' \
<table class="table"> \
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

	var res=new Array();
	for(var i=0; i<infos.length; i++){
		rec='<tr>';
		rec=rec + '<td>' + infos[i].Name + '</td>';
		rec=rec + '<td>' + infos[i].Status + '</td>';
		rec=rec + '<td>' + infos[i].Location + '</td>';
		rec=rec + '<td>' + infos[i].TaskId + '</td>';
		rec=rec+'</tr>';
		res.push(rec)
	}
	res.sort();
	content='';
	for(var i=0; i<res.length; i++){
		content = content + res[i];
	}

	return prefix + content + suffix;
}
