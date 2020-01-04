package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"../common/tools"
	"./service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Start Program Without Arguments !")
		return
	}

	if !tools.AllowExec() {
		fmt.Println("Program Is Already Start !")
		return
	}

	filename := os.Args[1]
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Println("LoginServer Arguments Unmarshal Error: ", err.Error())
		return
	}

	fmt.Printf("Program Arguments :\n%s\n", string(bytes))

	success := service.Start(config)

	if success {
		tools.BuildStopScript()

		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs

		service.Stop()

		tools.RemoveStopScript()

		fmt.Println("Got signal:", sig)
	} else {
		fmt.Println("LoginServer Start Failed !")
	}
}
