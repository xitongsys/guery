function HideAllInfo() {
	$("#running_list").hide();
	$("#queued_list").hide();
	$("#finished_list").hide();
	$("#executor_list").hide();	
}

function RunningButtonClick() {
	HideAllInfo();
	$("#running_list").show();
}

function QueuedButtonClick() {
	HideAllInfo();
	$("#queued_list").show();
}

function FinishedButtonClick() {
	HideAllInfo();
	$("#finished_list").show();
}

function ExecutorButtonClick() {
	HideAllInfo();
	$("#executor_list").show();
}
