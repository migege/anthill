#!/bin/bash
nohup ./api/log_api_rest_linux_amd64 --server_address=127.0.0.1:59801 > ./logs/log_api_rest.log &
