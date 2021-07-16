package nclink

/*NC-Link指令类别*/
const (
	CommandRegisterRequest    = 1
	CommandProbeVersion       = 2
	CommandDiscoveryRequest   = 3
	CommandDiscoveryResponse  = 4
	CommandProbeQueryRequest  = 5
	CommandProbeQueryResponse = 6
	CommandProbeSetRequest    = 7
	CommandProbeSetResponse   = 8
	CommandQueryRequest       = 9
	CommandQueryResponse      = 10
	CommandSetRequest         = 11
	CommandSetResponse        = 12
	CommandSample             = 13
)
const (
	OK = "OK"
	NG = "NG"
)
const (
	///查询、采样项中的operation字段值,表示取LIST类型数据长度
	OP_get_length = "get_length"
	///查询、采样项中的operation字段默认值,表示取数据的值
	OP_get_value = "get_value"
	///查询、采样项中的operation字段值,表示取HASH类型数据的所有key
	OP_get_keys = "get_keys"
	///查询、采样项中的operation字段值,表示取类型数据的属性
	OP_get_attributes = "get_attributes"
	///设置消息中的operation字段默认值，表示设置值
	OP_set_value = "set_value"
	///设置消息中的operation字段值，表示向LIST或HASH类型数据中添加元素
	OP_set_add = "add"
	///设置消息中的operation字段值，表示删除LIST或HASH类型数据中的元素
	OP_set_delelte = "delete"
)
const (
	///查询、设置消息中的reason字段的值，表示权限不足
	Reason_Debied = "Permission Denied"
	///查询、设置消息中的reason字段的值，表示参数不匹配
	Reason_Not_Match = "Parameter Not Match"
	///查询、设置消息中的reason字段的值，表示不支持的操作
	Reason_Unspported_Operation = "Unsupported Operation"
	///查询、设置消息中的reason字段的值，表示结点不存在
	Reason_Id_Not_Exists = "Id Not Exists"
	///查询、设置消息中的reason字段的值，表示内存不足
	Reason_No_Memory = "No Memory"
)

type BaseCommand struct {
	CommandID string `json:"@id,omitempty"`
	Code      string `json:"code,omitempty"`
	Reason    string `json:"reason,omitempty"`
}
type RegisterRequest struct {
	BaseCommand
	Device string `json:"deviceid"`
}
type ProbeVersion struct {
	BaseCommand
	Version string `json:"version"`
}
type DiscoveryRequest struct {
	BaseCommand
}
type DiscoveryResponse struct {
	BaseCommand
	Deviceids []string `json:"deviceids"`
}

type ProbeQueryRequest struct {
	BaseCommand
}
type ProbeQueryResponse struct {
	BaseCommand
	Probe string `json:"probe"`
}
type ProbeSetRequest struct {
	BaseCommand
	Probe string `json:"probe"`
}
type ProbeSetResponse struct {
	BaseCommand
}
type CommandParameter struct {
	Operation string `json:"operation,omitempty"`
	Offset    *int   `json:"offset,omitempty"`
	Length    *int   `json:"length,omitempty"`
}
type QueryParameter struct {
	CommandParameter
	Indexes []string      `json:"indexes,omitempty"`
	Keys    []string      `json:"keys,omitempty"`
	Values  []interface{} `json:"values,omitempty"`
}
type SetParameter struct {
	CommandParameter
	Index *int        `json:"index,omitempty"`
	Key   string      `json:"key,omitempty"`
	Value interface{} `json:"value"`
}
type CommandItem struct {
	BaseCommand
	ID   string `json:"id"`
	Node Node   `json:"-"`
}
type QueryItem struct {
	CommandItem
	Params *QueryParameter `json:"params,omitempty"`
	Values []interface{}   `json:"values,omitempty"`
}
type QueryRequest struct {
	BaseCommand
	Ids []QueryItem `json:"ids"`
}
type QueryResponse struct {
	BaseCommand
	Values []QueryItem `json:"values"`
}
type SetItem struct {
	CommandItem
	Params SetParameter `json:"params"`
}
type SetRequest struct {
	BaseCommand
	Values []SetItem `json:"values"`
}
type SetResponse struct {
	BaseCommand
	Results []SetItem `json:"results"`
}
type SampleDataItem struct {
	Data []interface{} `json:"data"`
}
type SampleData struct {
	ID        string           `json:"id"`
	BeginTime int64            `json:"beginTime"`
	Data      []SampleDataItem `json:"data"`
}
type SampleTask struct {
	ID             string
	SampleInterval uint16
	UploadInterval uint16
	SampleItems    []SampleItem
	CurTime        int64
}
