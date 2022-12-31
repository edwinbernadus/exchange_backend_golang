package main

import (
	"fiff_golang_draft/services"
	//"time"
	//"fiff_golang_draft/module_socket"
)

func main() {
	//database.DebugTestSaveItem()
	//database.DebugTestSaveItem2()
	//database.DebugTestSaveItem3()
	//database.DebugTestLoadItem()
	//database.DebugTestLoadItem2()
	// services.GetEnvTest()

	// var b = module_student.GetTotal(3, 4)
	// fmt.Printf(string(rune(b)))
	//services.TestBuy()
	//database.MainItem3()
	//module_student.Test5()

	//module_socket.WebSocketStart()

	//var addr = flag.String("addr", ":8080", "http service address")
	//err := http.ListenAndServe(*addr, nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	services.HandleServer()
}
