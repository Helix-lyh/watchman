package dataSchema

type MemInfo struct {
	Total int64 `json:"total"`//物理内存总大小
	Used int64 `json:"used"`//已经使用的物理内存大小
	Free int64 `json:"free"`//空闲的物理内存
	Shared int64 `json:"shared"`//多个进程共享内存的大小
	Buffers int64 `json:"buffers"`//做为缓存的内存大小
	Swap int64 `json:"swap"`//交互空间的使用状态
}

