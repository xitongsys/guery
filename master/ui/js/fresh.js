var dataLength=100;

var runningData = new Array(dataLength);
var queuedData = new Array(dataLength);
var finishedData = new Array(dataLength);
var agentData = new Array(dataLength);
var busyData = new Array(dataLength);
var totalData = new Array(dataLength);

var runningDB = new Dashboard("running_board", "");
var queuedDB = new Dashboard("queued_board", "");
var finishedDB = new Dashboard("finished_board", "");
var agentDB = new Dashboard("agent_board", "");
var busyDB = new Dashboard("busy_board", "");
var totalDB = new Dashboard("total_board", "");

for(i=0; i<dataLength; i++){
	runningData[i]=0;
	queuedData[i]=0;
	finishedData[i]=0;
	agentData[i]=0;
	busyData[i]=0;
	totalData[i]=0;		
}

//HideAllInfo();
//RunningButtonClick();

function freshClusterInfo(info) {
	runningData.shift(); runningData.push(info.Running);
	queuedData.shift(); queuedData.push(info.Queued);
	finishedData.shift(); finishedData.push(info.Finished);
	agentData.shift(); agentData.push(info.Agent);
	busyData.shift(); busyData.push(info.Busy);
	totalData.shift(); totalData.push(info.Total);

	runningDB.Plot(runningData, info.Running);
	queuedDB.Plot(queuedData, info.Queued);
	finishedDB.Plot(finishedData, info.Finished);
	agentDB.Plot(agentData, info.Agent);
	busyDB.Plot(busyData, info.Busy);
	totalDB.Plot(totalData, info.Total);

	$("#starttime").html(info.StartTime);
}

function freshAgentInfo(infos) {
	$("#agent_list").html(InfoAgentsToTable(infos));
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
	freshAgentInfo(info.AgentInfos)
	freshRunningTaskInfo(info.TaskInfos["DOING"])
	freshQueuedTaskInfo(info.TaskInfos["TODO"])
	
	var finishedInfos = new Array();
	finishedInfos = finishedInfos.concat(info.TaskInfos["DONE"], info.TaskInfos["FAILED"]);
	freshFinishedTaskInfo(finishedInfos);	
	
	setTimeout("fresh()", 1000);
}

fresh();


