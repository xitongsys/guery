var dataLength=100;

var runningData = new Array(dataLength);
var queuedData = new Array(dataLength);
var finishedData = new Array(dataLength);
var activeData = new Array(dataLength);
var busyData = new Array(dataLength);
var freeData = new Array(dataLength);

var runningDB = new Dashboard("running_board", "");
var queuedDB = new Dashboard("queued_board", "");
var finishedDB = new Dashboard("finished_board", "");
var activeDB = new Dashboard("active_board", "");
var busyDB = new Dashboard("busy_board", "");
var freeDB = new Dashboard("free_board", "");

for(i=0; i<dataLength; i++){
	runningData[i]=0;
	queuedData[i]=0;
	finishedData[i]=0;
	activeData[i]=0;
	busyData[i]=0;
	freeData[i]=0;		
}

//HideAllInfo();
//RunningButtonClick();

function freshClusterInfo(info) {
	runningData.shift(); runningData.push(info.Running);
	queuedData.shift(); queuedData.push(info.Queued);
	finishedData.shift(); finishedData.push(info.Finished);
	activeData.shift(); activeData.push(info.Active);
	busyData.shift(); busyData.push(info.Busy);
	freeData.shift(); freeData.push(info.Free);

	runningDB.Plot(runningData, info.Running);
	queuedDB.Plot(queuedData, info.Queued);
	finishedDB.Plot(finishedData, info.Finished);
	activeDB.Plot(activeData, info.Active);
	busyDB.Plot(busyData, info.Busy);
	freeDB.Plot(freeData, info.Free);

	$("#starttime").html(info.StartTime);
}

function freshExecutorInfo(infos) {
	$("#executor_list").html(InfoExecutorsToTable(infos));
}

function freshRunningTaskInfo(infos) {
	$("#running_list").html(InfoTasksToTable(infos))
}

function freshQueuedTaskInfo(infos) {
	$("#queued_list").html(InfoTasksToTable(infos))
}

function freshFinishedTaskInfo(infos) {
	$("#finished_list").html(InfoTasksToTable(infos))
}

function fresh() {
	xmlhttp=new XMLHttpRequest();
	xmlhttp.open("GET", "getinfo", false);
	xmlhttp.send();
	infoJson = xmlhttp.responseText;
	info = eval('(' + infoJson + ')');



	freshClusterInfo(info.ClusterInfo);
	freshExecutorInfo(info.ExecutorInfos)
	freshRunningTaskInfo(info.TaskInfos["DOING"])
	freshQueuedTaskInfo(info.TaskInfos["TODO"])
	freshFinishedTaskInfo(info.TaskInfos["DONE"])		
	
	setTimeout("fresh()", 1000);
}

fresh();


