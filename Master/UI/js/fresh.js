var dataLength=50;

var runningData = new Array(dataLength);
var queuedData = new Array(dataLength);
var finishedData = new Array(dataLength);
var activeData = new Array(dataLength);
var busyData = new Array(dataLength);
var freeData = new Array(dataLength);

var runningDB = new Dashboard("running_board", "RUNNING QUERIES");
var queuedDB = new Dashboard("queued_board", "QUEUED QUERIES");
var finishedDB = new Dashboard("finished_board", "FINISHED QUERIES");
var activeDB = new Dashboard("active_board", "ACTIVE EXECUTORS");
var busyDB = new Dashboard("busy_board", "BUSY EXECUTORS");
var freeDB = new Dashboard("free_board", "FREE EXECUTORS");

for(i=0; i<dataLength; i++){
	runningData[i]=0;
	queuedData[i]=0;
	finishedData[i]=0;
	activeData[i]=0;
	busyData[i]=0;
	freeData[i]=0;		
}

function fresh() {
	xmlhttp=new XMLHttpRequest();
	xmlhttp.open("GET", "getinfo", false);
	xmlhttp.send();
	infoJson = xmlhttp.responseText;
	info = eval('(' + infoJson + ')');

	runningData.shift(); runningData.push(info.Running);
	queuedData.shift(); queuedData.push(info.Queued);
	finishedData.shift(); finishedData.push(info.Finished);
	activeData.shift(); activeData.push(info.Active);
	busyData.shift(); busyData.push(info.Busy);
	freeData.shift(); freeData.push(info.Free);

	runningDB.Plot(runningData, info.Running);
	queuedDB.Plot(runningData, info.Queued);
	finishedDB.Plot(finishedData, info.Finished);
	activeDB.Plot(activeData, info.Active);
	busyDB.Plot(busyData, info.Busy);
	freeDB.Plot(freeData, info.Free);	
	
	setTimeout("fresh()", 1000);
}

fresh();
