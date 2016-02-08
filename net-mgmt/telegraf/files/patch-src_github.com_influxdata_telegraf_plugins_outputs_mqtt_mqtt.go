--- plugins/outputs/mqtt/mqtt.go.orig	2016-02-01 14:31:07 UTC
+++ plugins/outputs/mqtt/mqtt.go
@@ -9,7 +9,7 @@ import (
 	"strings"
 	"sync"
 
-	paho "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
+	paho "github.com/go-mqtt/mqtt"
 	"github.com/influxdata/telegraf"
 	"github.com/influxdata/telegraf/internal"
 	"github.com/influxdata/telegraf/plugins/outputs"
