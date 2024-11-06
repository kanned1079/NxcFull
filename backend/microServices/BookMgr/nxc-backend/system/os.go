package system

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type OsInfo struct {
	// 图表的
	CpuPercent  float64 `json:"cpu_percent"`  // cpu占用%
	MemPercent  float64 `json:"mem_percent"`  // 内存占用%
	DiskPercent float64 `json:"disk_percent"` // 磁盘占用%
	// 左边部分
	CpuModel string  `json:"cpu_model"`
	Cores    int     `json:"cores"`
	MemUsed  float64 `json:"mem_used"`
	MemSize  float64 `json:"mem_size"`
	DiskUsed float64 `json:"disk_used"`
	DiskSize float64 `json:"disk_size"`
	// 右边部分
	OsVersion       string    `json:"os_version"`
	KernelVersion   string    `json:"kernel_version"`
	OsArch          string    `json:"os_arch"`
	ProcessId       int       `json:"process_id"`
	NumsOfGoroutine int       `json:"nums_of_goroutine"`
	LastGCTime      time.Time `json:"last_gc_time"`
}

// getOsVersion 获取操作系统版本信息
func getOsVersion() string {
	out, err := exec.Command("cat", "/etc/redhat-release").Output()
	if err != nil {
		return runtime.GOOS
	}
	return string(out)
}

// getKernelVersion 获取内核版本信息
func getKernelVersion() string {
	out, err := exec.Command("uname", "-rs").Output()
	if err != nil {
		return ""
	}
	return string(out)
}

func getArchName() string {
	out, err := exec.Command("uname", "-m").Output()
	if err != nil {
		return ""
	}
	return string(out)
}

func (this *OsInfo) GetOsInfo() {
	memInfo, _ := mem.VirtualMemory()
	percentages, _ := cpu.Percent(0, false)
	cpuInfo, _ := cpu.Info()
	rootPart, _ := disk.Usage("/")
	//netInfo, _ := net.Interfaces()
	//log.Println(netInfo)
	// ServerLoad
	//log.Println("操作系统", runtime.GOOS)
	// chart
	this.CpuPercent = percentages[0]
	this.MemPercent = memInfo.UsedPercent
	this.DiskPercent = rootPart.UsedPercent
	// HardwareInfo
	this.CpuModel = cpuInfo[0].ModelName
	this.Cores = runtime.NumCPU()
	this.MemUsed = float64(memInfo.Used) / 1024 / 1024 / 1024
	this.MemSize = float64(memInfo.Total) / 1024 / 1024 / 1024
	this.DiskUsed = float64(rootPart.Used) / 1024 / 1024 / 1024
	this.DiskSize = float64(rootPart.Total) / 1024 / 1024 / 1024

	// OsInfo
	this.OsVersion = getOsVersion()
	this.KernelVersion = strings.TrimSpace(getKernelVersion())
	this.OsArch = strings.TrimSpace(getArchName())
	this.ProcessId = os.Getpid()
	this.NumsOfGoroutine = runtime.NumGoroutine()

	//log.Println(memInfo.SwapTotal, memInfo.SwapFree)

}
