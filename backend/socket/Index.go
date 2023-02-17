package socket

import (

	"net/http"
	"os"
	"backend/socket/service"
)

 
func Conn() {
	ip := os.Getenv("SOCKET_IP")
	port := os.Getenv("SOCKET_PORT")
    http.HandleFunc("/workAppSocket/shift", service.ShiftSocketHandler)
	http.ListenAndServe(ip + ":" + port, nil)
}