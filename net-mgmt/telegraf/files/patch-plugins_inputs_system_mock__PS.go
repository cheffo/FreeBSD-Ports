--- plugins/inputs/system/mock_PS.go.orig	2016-05-18 10:20:15 UTC
+++ plugins/inputs/system/mock_PS.go
@@ -15,55 +15,55 @@ type MockPS struct {
 	mock.Mock
 }
 
-func (m *MockPS) LoadAvg() (*load.LoadAvgStat, error) {
+func (m *MockPS) LoadAvg() (*load.AvgStat, error) {
 	ret := m.Called()
 
-	r0 := ret.Get(0).(*load.LoadAvgStat)
+	r0 := ret.Get(0).(*load.AvgStat)
 	r1 := ret.Error(1)
 
 	return r0, r1
 }
 
-func (m *MockPS) CPUTimes(perCPU, totalCPU bool) ([]cpu.CPUTimesStat, error) {
+func (m *MockPS) CPUTimes(perCPU, totalCPU bool) ([]cpu.TimesStat, error) {
 	ret := m.Called()
 
-	r0 := ret.Get(0).([]cpu.CPUTimesStat)
+	r0 := ret.Get(0).([]cpu.TimesStat)
 	r1 := ret.Error(1)
 
 	return r0, r1
 }
 
-func (m *MockPS) DiskUsage(mountPointFilter []string, fstypeExclude []string) ([]*disk.DiskUsageStat, error) {
+func (m *MockPS) DiskUsage(mountPointFilter []string, fstypeExclude []string) ([]*disk.UsageStat, error) {
 	ret := m.Called(mountPointFilter, fstypeExclude)
 
-	r0 := ret.Get(0).([]*disk.DiskUsageStat)
+	r0 := ret.Get(0).([]*disk.UsageStat)
 	r1 := ret.Error(1)
 
 	return r0, r1
 }
 
-func (m *MockPS) NetIO() ([]net.NetIOCountersStat, error) {
+func (m *MockPS) NetIO() ([]net.IOCountersStat, error) {
 	ret := m.Called()
 
-	r0 := ret.Get(0).([]net.NetIOCountersStat)
+	r0 := ret.Get(0).([]net.IOCountersStat)
 	r1 := ret.Error(1)
 
 	return r0, r1
 }
 
-func (m *MockPS) NetProto() ([]net.NetProtoCountersStat, error) {
+func (m *MockPS) NetProto() ([]net.ProtoCountersStat, error) {
 	ret := m.Called()
 
-	r0 := ret.Get(0).([]net.NetProtoCountersStat)
+	r0 := ret.Get(0).([]net.ProtoCountersStat)
 	r1 := ret.Error(1)
 
 	return r0, r1
 }
 
-func (m *MockPS) DiskIO() (map[string]disk.DiskIOCountersStat, error) {
+func (m *MockPS) DiskIO() (map[string]disk.IOCountersStat, error) {
 	ret := m.Called()
 
-	r0 := ret.Get(0).(map[string]disk.DiskIOCountersStat)
+	r0 := ret.Get(0).(map[string]disk.IOCountersStat)
 	r1 := ret.Error(1)
 
 	return r0, r1
@@ -87,10 +87,10 @@ func (m *MockPS) SwapStat() (*mem.SwapMe
 	return r0, r1
 }
 
-func (m *MockPS) NetConnections() ([]net.NetConnectionStat, error) {
+func (m *MockPS) NetConnections() ([]net.ConnectionStat, error) {
 	ret := m.Called()
 
-	r0 := ret.Get(0).([]net.NetConnectionStat)
+	r0 := ret.Get(0).([]net.ConnectionStat)
 	r1 := ret.Error(1)
 
 	return r0, r1
