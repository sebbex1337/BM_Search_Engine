package metrics

import (
    "log"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
)

//Collect system metrics

func CollectSystemMetrics(){
	go func() {
		for {
			//CPU Usage/Perfomance
			percent, err := cpu.Percent(0, false)
			if err == nil && len(percent) > 0 {
				CPUUsage.Set(percent[0])
			}

			//Memory Usage/Perfomance
			vmStat, err := mem.VirtualMemory()
			if err == nil {
				MemoryUsage.Set(vmStat.UsedPercent)
		} else {
			log.Printf("Error getting memory usage: %v", err)
		}
		//Disk Usage
		partitions, err := disk.Partitions(false)
		if err == nil {
			for _, p := range partitions{
				usage, err := disk.Usage(p.Mountpoint)
				if err == nil {
					DiskUsage.WithLabelValues(p.Device).Set(float64(usage.UsedPercent))
				} else {
					log.Printf("Error getting disk usage: %v", err)
				}
				time.Sleep(10 * time.Second) // Sleep for 10 seconds before collecting the next disk usage, can be adjusted as needed
			}
		}
	}
}()
}

