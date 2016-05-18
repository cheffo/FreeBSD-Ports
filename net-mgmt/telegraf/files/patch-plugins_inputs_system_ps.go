--- plugins/inputs/system/ps.go.orig	2016-05-18 10:20:15 UTC
+++ plugins/inputs/system/ps.go
@@ -13,14 +13,14 @@ import (
 )
 
 type PS interface {
-	CPUTimes(perCPU, totalCPU bool) ([]cpu.CPUTimesStat, error)
-	DiskUsage(mountPointFilter []string, fstypeExclude []string) ([]*disk.DiskUsageStat, error)
-	NetIO() ([]net.NetIOCountersStat, error)
-	NetProto() ([]net.NetProtoCountersStat, error)
-	DiskIO() (map[string]disk.DiskIOCountersStat, error)
+	CPUTimes(perCPU, totalCPU bool) ([]cpu.TimesStat, error)
+	DiskUsage(mountPointFilter []string, fstypeExclude []string) ([]*disk.UsageStat, error)
+	NetIO() ([]net.IOCountersStat, error)
+	NetProto() ([]net.ProtoCountersStat, error)
+	DiskIO() (map[string]disk.IOCountersStat, error)
 	VMStat() (*mem.VirtualMemoryStat, error)
 	SwapStat() (*mem.SwapMemoryStat, error)
-	NetConnections() ([]net.NetConnectionStat, error)
+	NetConnections() ([]net.ConnectionStat, error)
 }
 
 func add(acc telegraf.Accumulator,
@@ -32,17 +32,17 @@ func add(acc telegraf.Accumulator,
 
 type systemPS struct{}
 
-func (s *systemPS) CPUTimes(perCPU, totalCPU bool) ([]cpu.CPUTimesStat, error) {
-	var cpuTimes []cpu.CPUTimesStat
+func (s *systemPS) CPUTimes(perCPU, totalCPU bool) ([]cpu.TimesStat, error) {
+	var cpuTimes []cpu.TimesStat
 	if perCPU {
-		if perCPUTimes, err := cpu.CPUTimes(true); err == nil {
+		if perCPUTimes, err := cpu.Times(true); err == nil {
 			cpuTimes = append(cpuTimes, perCPUTimes...)
 		} else {
 			return nil, err
 		}
 	}
 	if totalCPU {
-		if totalCPUTimes, err := cpu.CPUTimes(false); err == nil {
+		if totalCPUTimes, err := cpu.Times(false); err == nil {
 			cpuTimes = append(cpuTimes, totalCPUTimes...)
 		} else {
 			return nil, err
@@ -54,8 +54,8 @@ func (s *systemPS) CPUTimes(perCPU, tota
 func (s *systemPS) DiskUsage(
 	mountPointFilter []string,
 	fstypeExclude []string,
-) ([]*disk.DiskUsageStat, error) {
-	parts, err := disk.DiskPartitions(true)
+) ([]*disk.UsageStat, error) {
+	parts, err := disk.Partitions(true)
 	if err != nil {
 		return nil, err
 	}
@@ -70,7 +70,7 @@ func (s *systemPS) DiskUsage(
 		fstypeExcludeSet[filter] = true
 	}
 
-	var usage []*disk.DiskUsageStat
+	var usage []*disk.UsageStat
 
 	for _, p := range parts {
 		if len(mountPointFilter) > 0 {
@@ -83,7 +83,7 @@ func (s *systemPS) DiskUsage(
 		}
 		mountpoint := os.Getenv("HOST_MOUNT_PREFIX") + p.Mountpoint
 		if _, err := os.Stat(mountpoint); err == nil {
-			du, err := disk.DiskUsage(mountpoint)
+			du, err := disk.Usage(mountpoint)
 			du.Path = p.Mountpoint
 			if err != nil {
 				return nil, err
@@ -102,20 +102,20 @@ func (s *systemPS) DiskUsage(
 	return usage, nil
 }
 
-func (s *systemPS) NetProto() ([]net.NetProtoCountersStat, error) {
-	return net.NetProtoCounters(nil)
+func (s *systemPS) NetProto() ([]net.ProtoCountersStat, error) {
+	return net.ProtoCounters(nil)
 }
 
-func (s *systemPS) NetIO() ([]net.NetIOCountersStat, error) {
-	return net.NetIOCounters(true)
+func (s *systemPS) NetIO() ([]net.IOCountersStat, error) {
+	return net.IOCounters(true)
 }
 
-func (s *systemPS) NetConnections() ([]net.NetConnectionStat, error) {
-	return net.NetConnections("all")
+func (s *systemPS) NetConnections() ([]net.ConnectionStat, error) {
+	return net.Connections("all")
 }
 
-func (s *systemPS) DiskIO() (map[string]disk.DiskIOCountersStat, error) {
-	m, err := disk.DiskIOCounters()
+func (s *systemPS) DiskIO() (map[string]disk.IOCountersStat, error) {
+	m, err := disk.IOCounters()
 	if err == internal.NotImplementedError {
 		return nil, nil
 	}
