package sensor

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/disk"
	"time"
)

func GetCPUUsage() (float64, error){
	percentages, err := cpu.Percent(1*time.Second, false)
	if err != nil{
		return 0, err
	}
	return percentages[0], nil
}

func GetMemoryUsage() (usedPercent float64, total uint64, err error){
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}
	return vmStat.UsedPercent, vmStat.Total, nil
}

func GetDiskUsage(path string) (usedPercent float64, total uint64, err error){
	usageStat, err := disk.Usage(path)
	if err != nil {
		return 0, 0, err
	}
	return usageStat.UsedPercent, usageStat.Total, nil
}