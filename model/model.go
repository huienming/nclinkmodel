package nclink

import (
	"encoding/json"
	"fmt"
)

const rootType string = "NC_LINK_ROOT"
const sampleChannelType string = "SAMPLE_CHANNEL"

func buildPath(node Node, parent Node) {
	if node.NodeNumber() != "" {
		node.SetNodePath(fmt.Sprintf("%s:%s#%s", parent.NodePath(), node.NodeType(), node.NodeNumber()))
	} else {
		node.SetNodePath(fmt.Sprintf("%s:%s", parent.NodePath(), node.NodeType()))
	}
}
func buildContainderIDMap(c ContainerNode, parent Node, idMap map[string]Node) bool {
	if len(c.NodeID()) == 0 {
		return false
	}
	_, ok := idMap[c.NodeID()]
	if ok {
		return false
	}
	CurNode := c
	idMap[CurNode.NodeID()] = CurNode
	buildPath(CurNode, parent)
	for _, cfg := range CurNode.NodeConfigs() {
		if len(cfg.ID) == 0 {
			return false
		}
		_, ok := idMap[cfg.ID]
		if ok {
			return false
		}
		idMap[cfg.ID] = &cfg
		buildPath(&cfg, CurNode)
	}
	for _, item := range CurNode.NodeDataItems() {
		if len(item.ID) == 0 {
			return false
		}
		_, ok := idMap[item.ID]
		if ok {
			return false
		}
		idMap[item.ID] = &item
		buildPath(&item, CurNode)
	}
	for _, subComponent := range CurNode.NodeComponents() {
		buildPath(&subComponent, CurNode)
		if !buildContainderIDMap(&subComponent, CurNode, idMap) {
			return false
		}
	}
	return true
}
func buildRootIDMap(m *Root, idMap map[string]Node) bool {
	if len(m.ID) == 0 || idMap[m.ID] != nil {
		return false
	}
	idMap[m.ID] = m
	m.Path = m.Type
	for _, dev := range m.Devices {
		if !buildContainderIDMap(&dev, m, idMap) {
			return false
		}
	}
	return true
}

func buildModelIDMap(m *Root) map[string]Node {
	idMap := make(map[string]Node)
	if !buildRootIDMap(m, idMap) {
		return nil
	}
	return idMap
}

//NewModel 新建模型对象
func NewModel(modelText []byte) *Root {
	nclinkObject := &Root{}
	err := json.Unmarshal(modelText, nclinkObject)
	if err != nil {
		fmt.Printf("json.Unmarshal:%v\n", err.Error())
		return nil
	}
	if !checkModel(nclinkObject) {
		fmt.Printf("CheckModel Failed\n")
		return nil
	}
	nclinkObject.idNodeMap = buildModelIDMap(nclinkObject)
	if nclinkObject.idNodeMap == nil {
		return nil
	}
	if len(nclinkObject.Devices) == 1 {
		device := &nclinkObject.Devices[0]
		for cfindex, Config := range device.Configs {
			if Config.Type == sampleChannelType {
				ConfigPointer := device.Configs[cfindex]
				for itemIndex, item := range ConfigPointer.SampleItems {
					node, ok := nclinkObject.idNodeMap[item.ID]
					if !ok {
						continue
					}
					ConfigPointer.SampleItems[itemIndex].Node = node
				}
				device.SampleTasks = append(device.SampleTasks, SampleTask{ID: Config.ID, SampleInterval: Config.SampleInterval, UploadInterval: Config.UploadInterval, SampleItems: Config.SampleItems})
			}
		}
	}
	return nclinkObject
}
func checkModel(m *Root) bool {
	if m.Type != rootType || m.Devices == nil || len(m.Devices) != 1 {
		return false
	}
	return true
}
