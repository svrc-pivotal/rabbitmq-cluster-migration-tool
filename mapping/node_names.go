package mapping

import (
	"github.com/pivotal-cf/rabbitmq-cluster-migration-tool/parsers"
)

type NodeMapping map[string]string

func NodeNames(oldIPAddressesWithNodeNames, newIPAddressesWithNodeNames parsers.IPAddressesWithNodeNames) NodeMapping {
	mappingResult := NodeMapping{}
	for ip, node := range oldIPAddressesWithNodeNames {
		if newNodeName := newIPAddressesWithNodeNames.NodeNameByIP(ip); newNodeName != "" {
			if newNodeName != node {
				mappingResult[node] = newNodeName
			}
		}
	}
	return mappingResult
}
