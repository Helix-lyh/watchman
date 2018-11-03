package utils

import "watchman/dataSchema"

//获取当前系统的Cpu占用率
func CpuInfo() (info dataSchema.CpuInfo) {
	cmd := `top -b -n 1 | grep Cpu`
	info = Run(cmd)
	return info
}

//获取当前系统的负载信息
func LoadInfo() (info dataSchema.Load) {
	cmd := `top -b -n 1 | grep Cpu`
	info = Run(cmd)
	return info
}

//获取当前系统的内存使用信息
func MemInfo() (info dataSchema.MemInfo) {
	cmd := `top -b -n 1 | grep Cpu`
	info = Run(cmd)
	return info
}

//获取当前系统的网络连接信息
func NetInfo() (info dataSchema.NetStatus) {
	cmd := `top -b -n 1 | grep Cpu`
	info = Run(cmd)
	return info
}

