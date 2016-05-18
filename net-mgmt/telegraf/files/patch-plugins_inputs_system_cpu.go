--- plugins/inputs/system/cpu.go.orig	2016-05-18 10:20:15 UTC
+++ plugins/inputs/system/cpu.go
@@ -11,7 +11,7 @@ import (
 
 type CPUStats struct {
 	ps        PS
-	lastStats []cpu.CPUTimesStat
+	lastStats []cpu.TimesStat
 
 	PerCPU   bool `toml:"percpu"`
 	TotalCPU bool `toml:"totalcpu"`
@@ -105,7 +105,7 @@ func (s *CPUStats) Gather(acc telegraf.A
 	return nil
 }
 
-func totalCpuTime(t cpu.CPUTimesStat) float64 {
+func totalCpuTime(t cpu.TimesStat) float64 {
 	total := t.User + t.System + t.Nice + t.Iowait + t.Irq + t.Softirq + t.Steal +
 		t.Guest + t.GuestNice + t.Idle
 	return total
