package nclink

//Node 节点接口
type Node interface {
	NodeID() string
	NodeType() string
	NodeName() string
	NodeNumber() string
	NodeDescription() string
	NodePath() string
	NodeDataType() string
	SetNodePath(path string)
}
type ContainerNode interface {
	Node
	NodeConfigs() []Config
	NodeDataItems() []DataItem
	NodeComponents() []Component
}
type BaseNode struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Number      string `json:"number,omitempty"`
	Name        string `json:"name,omitempty"`
	DataType    string `json:"datatype,omitempty"`
	Description string `json:"description,omitempty"`
	Path        string `json:"-"`
}
type BaseLeafNode struct {
	BaseNode
	Value    interface{} `json:"value,omitempty"`
	Settable bool        `json:"setable,omitempty"`
	Source   string      `json:"source,omitempty"`
	Units    string      `json:"units,omitempty"`
}
type BaseContainerNode struct {
	BaseNode
	Configs    []Config    `json:"configs,omitempty"`
	DataItems  []DataItem  `json:"dataItems,omitempty"`
	Components []Component `json:"components,omitempty"`
}
type DataItem struct {
	BaseLeafNode
}
type SampleItem struct {
	QueryItem
}
type Config struct {
	BaseLeafNode
	SampleInterval uint16       `json:"sampleInterval,omitempty"`
	UploadInterval uint16       `json:"uploadInterval,omitempty"`
	SampleItems    []SampleItem `json:"ids,omitempty"`
}
type Component struct {
	BaseContainerNode
}
type Device struct {
	BaseContainerNode
	GUID        string       `json:"guid,omitempty"`
	Version     string       `json:"version,omitempty"`
	SampleTasks []SampleTask `json:"-"`
}
type Root struct {
	BaseNode
	Devices   []Device        `json:"devices,omitempty"`
	idNodeMap map[string]Node `json:"-"`
}

func (_this *BaseNode) NodeID() string {
	return _this.ID
}
func (_this *BaseNode) NodeType() string {
	return _this.Type
}
func (_this *BaseNode) NodeNumber() string {
	return _this.Number
}
func (_this *BaseNode) NodeName() string {
	return _this.Name
}
func (_this *BaseNode) NodeDescription() string {
	return _this.Description
}
func (_this *BaseNode) NodePath() string {
	return _this.Path
}
func (_this *BaseNode) NodeDataType() string {
	return _this.DataType
}
func (_this *BaseNode) SetNodePath(path string) {
	_this.Path = path
}
func (_this *BaseContainerNode) NodeConfigs() []Config {
	return _this.Configs
}
func (_this *BaseContainerNode) NodeDataItems() []DataItem {
	return _this.DataItems
}
func (_this *BaseContainerNode) NodeComponents() []Component {
	return _this.Components
}
