--- plugins/inputs/system/system.go.orig	2016-05-18 10:20:15 UTC
+++ plugins/inputs/system/system.go
@@ -22,12 +22,12 @@ func (_ *SystemStats) Description() stri
 func (_ *SystemStats) SampleConfig() string { return "" }
 
 func (_ *SystemStats) Gather(acc telegraf.Accumulator) error {
-	loadavg, err := load.LoadAvg()
+	loadavg, err := load.Avg()
 	if err != nil {
 		return err
 	}
 
-	hostinfo, err := host.HostInfo()
+	hostinfo, err := host.Info()
 	if err != nil {
 		return err
 	}
