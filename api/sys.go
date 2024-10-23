package api

import (
	"encoding/json"
	"net/http"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	//"github.com/shirou/gopsutil/host"
)

type SystemInfo struct {
	CPU   string        `json:"cpu"`
	CPUUsage   float64        `json:"cpu_usage"`
	Memory     MemoryInfo     `json:"memory"`
	OS    string   `json:"os"`
}

type MemoryInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}


func GetSystemInfoHandler(w http.ResponseWriter, r *http.Request) {
	cpuUsage, _ := cpu.Percent(0, false)
	c, _ := cpu.Info()
	memStats, _ := mem.VirtualMemory()
	//n, _ := host.Info()
	os := "Ubuntu 22.04.5 LTS (Jammy Jellyfish) aarch64(Py3.7.16)"
	systemInfo := SystemInfo{
		CPU: c[0].ModelName,
		CPUUsage: cpuUsage[0],
		Memory: MemoryInfo{
			Total: memStats.Total,
			Used:  memStats.Used,
			Free:  memStats.Free,
		},
		OS: os,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(systemInfo)
}