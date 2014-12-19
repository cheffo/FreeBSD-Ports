--- src/unix-manager.c.orig
+++ src/unix-manager.c
@@ -102,26 +102,36 @@ int UnixNew(UnixCommand * this)
     TAILQ_INIT(&this->tasks);
     TAILQ_INIT(&this->clients);
 
-    /* Create socket dir */
-    ret = mkdir(SOCKET_PATH, S_IRWXU|S_IXGRP|S_IRGRP);
-    if ( ret != 0 ) {
-        int err = errno;
-        if (err != EEXIST) {
-            SCLogError(SC_ERR_OPENING_FILE,
-                    "Cannot create socket directory %s: %s", SOCKET_PATH, strerror(err));
-            return 0;
-        }
-    }
-
     if (ConfGet("unix-command.filename", &socketname) == 1) {
-        int socketlen = strlen(SOCKET_PATH) + strlen(socketname) + 2;
-        sockettarget = SCMalloc(socketlen);
-        if (unlikely(sockettarget == NULL)) {
-            SCLogError(SC_ERR_MEM_ALLOC, "Unable to allocate socket name");
-            return 0;
+        if (PathIsAbsolute(socketname)) {
+            sockettarget = SCStrdup(socketname);
+            if (unlikely(sockettarget == NULL)) {
+                SCLogError(SC_ERR_MEM_ALLOC, "Unable to allocate socket name");
+                return 0;
+            }
+        } else {
+            int socketlen = strlen(SOCKET_PATH) + strlen(socketname) + 2;
+            sockettarget = SCMalloc(socketlen);
+            if (unlikely(sockettarget == NULL)) {
+                SCLogError(SC_ERR_MEM_ALLOC, "Unable to allocate socket name");
+                return 0;
+            }
+            snprintf(sockettarget, socketlen, "%s/%s", SOCKET_PATH, socketname);
+
+            /* Create socket dir */
+            ret = mkdir(SOCKET_PATH, S_IRWXU|S_IXGRP|S_IRGRP);
+            if ( ret != 0 ) {
+                int err = errno;
+                if (err != EEXIST) {
+                    SCFree(sockettarget);
+                    SCLogError(SC_ERR_OPENING_FILE,
+                            "Cannot create socket directory %s: %s", SOCKET_PATH, strerror(err));
+                    return 0;
+                }
+            }
+
         }
-        snprintf(sockettarget, socketlen, "%s/%s", SOCKET_PATH, socketname);
-        SCLogInfo("Use unix socket file '%s'.", sockettarget);
+        SCLogInfo("Using unix socket file '%s'", sockettarget);
     }
     if (sockettarget == NULL) {
         sockettarget = SCStrdup(SOCKET_TARGET);
