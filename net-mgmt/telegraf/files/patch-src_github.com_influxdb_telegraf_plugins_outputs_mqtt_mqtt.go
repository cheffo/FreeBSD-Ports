--- plugins/outputs/mqtt/mqtt.go.orig	2016-01-19 14:54:35 UTC
+++ plugins/outputs/mqtt/mqtt.go
@@ -9,7 +9,7 @@ import (
 	"strings"
 	"sync"
 
-	paho "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
+	paho "github.com/go-mqtt/mqtt"
 	"github.com/influxdb/influxdb/client/v2"
 	"github.com/influxdb/telegraf/internal"
 	"github.com/influxdb/telegraf/plugins/outputs"
