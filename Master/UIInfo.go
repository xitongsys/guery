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
	res.ExecutorInfos = self.GetUIExecutorInfos()
	res.TaskInfos = self.GetUITaskInfos(res.ExecutorInfos)

	return res
}
