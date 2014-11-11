FreeBSD Ports
==================

hyperleveldb - FreeBSD Port for https://github.com/rescrv/HyperLevelDB

* Additional this port renames all "leveldb" stuff to "hyperleveldb", so it can be used by influxdb.

==================

suricata - https://bugs.freebsd.org/bugzilla/show_bug.cgi?id=193220

* security/suricata: Change summary

- Added JSON knob - this allows Suricata to be compiled with JSON output support
- Added GEOIP knob - this allows Suricata to support rules with geoip word
- Added HTP_PORT knob - this make the use of www/libhtp-suricata optional. E.g. user can choose between build-in and port version.

Default behavior is not changed.

==================
