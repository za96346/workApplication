#!/bin/bash
expect << EOF
set timeout 600
spawn certbot certonly --nginx --email za96346@gmail.com --agree-tos -d workhard.app
expect {
    "*(Y)es/(N)o:*" {
        send "n\r"
        send "exit\r"
    }
}
expect eof
EOF