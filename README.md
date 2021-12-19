# Subject

This library is common logger library for golang project in BearAtHome

# Environment Setting

Library will use `BATH_LOGGER_LEVEL` to set default log level

# Usage of each log level

Level | Description
--|----
Error| Log about value error or everythine will cause system failure
Warn| If there is something wierd, which may cause system failure or data corrupted. It should be a warning.
Info| System access log or status of running.
Debug| Key value or flow when developer is debugging
Trace| Detail value of each input/output or flow to trace the detail behavior of system
