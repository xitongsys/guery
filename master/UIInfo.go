package Master

import ()

type UIInfo struct {
	ClusterInfo *UIClusterInfo
	AgentInfos  []*UIAgentInfo
	TaskInfos   map[string][]*UITaskInfo
}

func (self *Master) GetInfo() *UIInfo {
	res := &UIInfo{}
	res.AgentInfos = self.GetUIAgentInfos()
	res.TaskInfos = self.GetUITaskInfos()

	res.ClusterInfo = self.GetUIClusterInfo()
	res.ClusterInfo.Agent = len(res.AgentInfos)
	for _, info := range res.AgentInfos {
		res.ClusterInfo.Busy += info.ExecutorNumber
		res.ClusterInfo.Total += info.MaxExecutorNumber
	}

	return res
}
