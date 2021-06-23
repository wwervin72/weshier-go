package sd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	// B 1 bit
	B = 1
	// KB 1KB
	KB = 1024 * B
	// MB 1MB
	MB = 1024 * KB
	// GB 1GB
	GB = 1024 * MB
)

// HealthCheck
// @Summary 服务器健康检测
// @Description 服务器健康检测
// @Tags healthy check
// @Produce plain
// @Success 200 string string "服务器健康检测"
// @Router /api/sd/health [get]
func HealthCheck(c *gin.Context) {
	message := "ok"
	c.String(http.StatusOK, message)
}

// @Summary 服务器磁盘检测
// @Description 服务器磁盘使用检测
// @Tags healthy check
// @Produce plain
// @Success 200 string string "服务器磁盘健康检测"
// @Router /api/sd/disk [get]
func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message)
}

// CPUCheck
// @Summary 服务器 cpu 检测
// @Description 服务器 cpu 检测
// @Tags healthy check
// @Produce plain
// @Success 200 string string "服务器 cpu 健康检测"
// @Router /api/sd/cpu [get]
func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)
	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := "ok"
	message = fmt.Sprintf("\n%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores)
	c.String(status, message)
}

// RAMCheck
// @Summary 服务器内存检测
// @Description 服务器内存检测
// @Tags healthy check
// @Produce plain
// @Router /api/sd/ram [get]
// @Success 200 string string "服务器内存健康检测"
func RAMCheck(c *gin.Context) {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent > 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("\n%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message)
}
