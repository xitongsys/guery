package Master

import ()

type UIAgentInfo struct {
	Name     string
	Location string
	Status   string
	TaskId   int64
}

func (self *Master) GetUIAgentInfos() []*UIAgentInfo {
	agents := self.Topology.GetAgents()
	res := []*UIAgentInfo{}
	for _, agent := range agents {
		res = append(res,
			&UIAgentInfo{
				Name:     agent.Name,
				Location: agent.GetURL(),
			},
		)
	}
	return res
}
