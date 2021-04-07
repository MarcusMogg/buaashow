package utils

import (
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Server struct {
	Os   `json:"os"`
	Cpu  `json:"cpu"`
	Ram  `json:"ram"`
	Disk `json:"disk"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Ram struct {
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}

func (s *Server) InitOS() {
	s.Os.GOOS = runtime.GOOS
	s.Os.NumCPU = runtime.NumCPU()
	s.Os.Compiler = runtime.Compiler
	s.Os.GoVersion = runtime.Version()
	s.Os.NumGoroutine = runtime.NumGoroutine()
}

func (s *Server) InitCPU() {
	if cores, err := cpu.Counts(false); err != nil {
		return
	} else {
		s.Cpu.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return
	} else {
		s.Cpu.Cpus = cpus
	}
}

func (s *Server) InitRAM() {
	if u, err := mem.VirtualMemory(); err != nil {
		return
	} else {
		s.Ram.UsedMB = int(u.Used) / MB
		s.Ram.TotalMB = int(u.Total) / MB
		s.Ram.UsedPercent = int(u.UsedPercent)
	}
}

func (s *Server) InitDisk() {
	if u, err := disk.Usage("/buaashow"); err != nil {
		return
	} else {
		s.Disk.UsedMB = int(u.Used) / MB
		s.Disk.UsedGB = int(u.Used) / GB
		s.Disk.TotalMB = int(u.Total) / MB
		s.Disk.TotalGB = int(u.Total) / GB
		s.Disk.UsedPercent = int(u.UsedPercent)
	}
}
