package Master

import ()

type UIInfo struct {
	ClusterInfo   *UIClusterInfo
	TaskInfos     map[string][]*UITaskInfo
	ExecutorInfos []*UIExecutorInfo
}

func (self *Master) GetInfo() *UIInfo {
	res := &UIInfo{}
	res.ClusterInfo = self.GetUIClusterInfo()
	res.TaskInfos = self.GetUITaskInfos()
	res.ExecutorInfos = self.GetUIExecutorInfos()
	return res
}
