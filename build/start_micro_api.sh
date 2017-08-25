#!/bin/bash
nohup micro --server_address=127.0.0.1:59800 --broker_address=127.0.0.1:59700 api --handler=proxy --namespace=com.mayibot.ah.api > ./logs/micro_api.log &
