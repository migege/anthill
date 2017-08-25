nohup ./server/log_server_linux_amd64 --server_address=120.27.122.8:59901 > ./logs/log_server.log &
nohup ./server/time_server_linux_amd64 --server_address=120.27.122.8:59902 > ./logs/time_server.log &
nohup ./server/user_server_linux_amd64 --server_address=120.27.122.8:59903 > ./logs/user_server.log &
