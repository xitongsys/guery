package Master

import ()

type UIInfo struct {
	ClusterInfo *UIClusterInfo
	TaskInfos   map[string][]*UITaskInfo
	AgentInfos  []*UIAgentInfo
}

func (self *Master) GetInfo() *UIInfo {
	res := &UIInfo{}
	res.ClusterInfo = self.GetUIClusterInfo()
	res.AgentInfos = self.GetUIAgentInfos()
	res.TaskInfos = self.GetUITaskInfos(res.AgentInfos)

	return res
}
