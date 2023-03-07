package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func GetSystemInformation() error {
	host, err := host.Info()
	cpu, err := cpu.Info()
	ram, err := mem.VirtualMemory()
	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Fatalf("[ERR] unable to detect PC Stats: %v", err)
	}

	//TODO experimental

	// test, _ := net.LookupIP(host.Hostname)
	// log.Println(test)

	sys := struct {
		Hostname        string
		OS              string
		PlatformVersion string
		CPU             string
		RAM             uint64
		Disk            uint64
	}{
		Hostname:        host.Hostname,
		OS:              host.OS,
		PlatformVersion: host.PlatformVersion,
		CPU:             cpu[0].ModelName,
		RAM:             ram.Total / 1024 / 1024,
		Disk:            diskStat.Total / 1024 / 1024,
	}

	data, err := json.Marshal(sys)
	if err != nil {
		log.Fatalf("Failed to Marshal: %v", err)
	}

	_, err = http.Post("http://127.0.0.1:3333/api/sysinfo", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed while sending system information to server: %v", err)
	}

	log.Println(sys)

	return nil
}
