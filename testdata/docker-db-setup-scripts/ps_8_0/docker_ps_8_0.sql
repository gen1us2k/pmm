#!/bin/bash

SET GLOBAL slow_query_log='ON';
GRANT SELECT, PROCESS, SUPER, REPLICATION CLIENT, RELOAD ON *.* TO 'pmm-agent'@'%';
