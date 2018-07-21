package Master

import ()

type UIInfo struct {
	ClusterInfo *UIClusterInfo
	AgentInfos  []*UIAgentInfo
	TaskInfos   map[string][]*UITaskInfo
}

func (self *Master) GetInfo() *UIInfo {
	res := &UIInfo{}
	res.ClusterInfo = self.GetUIClusterInfo()
	res.AgentInfos = self.GetUIAgentInfos()
	res.TaskInfos = self.GetUITaskInfos(res.AgentInfos)
	return res
}
