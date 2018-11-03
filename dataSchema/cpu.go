package dataSchema

type CpuInfo struct {
	Processor       int     `json:"processor"`       //系统中逻辑处理核的编号。对于单核处理器，认为是其CPU编号，对于多核处理器则可以是物理核、或者使用超线程技术虚拟的逻辑核
	VendorId        int     `json:"vendorId"`        //CPU制造商
	CpuFamily       int     `json:"cpuFamily"`       //CPU产品系列代号
	ModelName       string  `json:"modelName"`       //CPU属于的名字及其编号、标称主频
	Stepping        float64 `json:"stepping"`        //CPU属于制作更新版本
	CpuMHz          float64 `json:"cpuMHz"`          //CPU的实际使用主频
	CacheSize       string  `json:"cacheSize"`       //CPU二级缓存大小
	PhysicalId      int     `json:"physicalId"`      //单个CPU的标号
	Siblings        int     `json:"siblings"`        //单个CPU逻辑物理核数
	CoreId          int     `json:"coreId"`          //当前物理核在其所处CPU中的编号，这个编号不一定连续
	CpuCores        int     `json:"cpuCores"`        //该逻辑核所处CPU的物理核数
	Apicid          int     `json:"apicid"`          //用来区分不同逻辑核的编号，系统中每个逻辑核的此编号必然不同，此编号不一定连续
	Fpu             bool    `json:"fpu"`             //是否具有浮点运算单元（Floating Point Unit）
	FpuException    bool    `json:"fpuException"`    //是否支持浮点计算异常
	CpuidLevel      int     `json:"cpuidLevel"`      //执行cpuid指令前，eax寄存器中的值，根据不同的值cpuid指令会返回不同的内容
	Wp              bool    `json:"wp"`              //表明当前CPU是否在内核态支持对用户空间的写保护（Write Protection）
	Flags           string  `json:"flags"`           //当前CPU支持的功能
	Bogomips        float64 `json:"bogomips"`        //在系统内核启动时粗略测算的CPU速度（Million Instructions Per Second）
	ClflushSize     int     `json:"clflushSize"`     //每次刷新缓存的大小单位
	CacheAlignment  int     `json:"cacheAlignment"`  //缓存地址对齐单位
	AddressSizes    string  `json:"addressSizes"`    //可访问地址空间位数
	PowerManagement string  `json:"powerManagement"` //对能源管理的支持
	UserPer         float64 `json:"userPer"`         //用户空间占用cpu的百分比
	SysPer          float64 `json:"sysPer"`          //system 内核空间占用cpu的百分比
	NicedPer        float64 `json:"nicedPer"`        //niced 改变过优先级的进程占用cpu的百分比
	IdPer           float64 `json:"idPer"`           //空闲cpu百分比
	WaitPer         float64 `json:"waitPer"`         //IO wait IO等待占用cpu的百分比
	HardCPer        float64 `json:"hardCPer"`        //Hardware IRQ 硬中断 占用cpu的百分比
	SoftCPer        float64 `json:"softCPer"`        //software 软中断 占用cpu的百分比
	StPer           float64 `json:"stPer"`           //被hypervisor偷去的时间
}
