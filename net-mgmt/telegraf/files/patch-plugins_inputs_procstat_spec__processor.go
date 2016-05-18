--- plugins/inputs/procstat/spec_processor.go.orig	2016-05-18 10:20:15 UTC
+++ plugins/inputs/procstat/spec_processor.go
@@ -65,7 +65,7 @@ func (p *SpecProcessor) pushMetrics() {
 		fields[prefix+"write_bytes"] = io.WriteCount
 	}
 
-	cpu_time, err := p.proc.CPUTimes()
+	cpu_time, err := p.proc.Times()
 	if err == nil {
 		fields[prefix+"cpu_time_user"] = cpu_time.User
 		fields[prefix+"cpu_time_system"] = cpu_time.System
@@ -80,7 +80,7 @@ func (p *SpecProcessor) pushMetrics() {
 		fields[prefix+"cpu_time_guest_nice"] = cpu_time.GuestNice
 	}
 
-	cpu_perc, err := p.proc.CPUPercent(time.Duration(0))
+	cpu_perc, err := p.proc.Percent(time.Duration(0))
 	if err == nil && cpu_perc != 0 {
 		fields[prefix+"cpu_usage"] = cpu_perc
 	}
