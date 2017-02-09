package main

import (
	"cgss/cg"
	"cgss/ipc"
)

var centerClient *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}
	return nil
}

func Help(args []string) int{
    fmt.Println("
        Commands:
            login <username><level><exp>
            logout <username>
            send <message>
            quit(q)
            help(h)
    ")
    return 0
}

func Quit(args []string) int{
    return 1
}

func Logout(args []string) int{
    if len(args) != 2{
        fmt.Println("USAGE: logout <username>")
        return 0
    }
    centerClient.RemovePlayer(args[1])
    return 0
}

func Login(args []string)int{
    if len(args) != 4{
        fmt.Println("USAGE: login <username> <level> <exp>")
        return 0
    }
    level, err := strconv.Atoi(args[2])
    if err != nil{
        fmt.Println("Invalid Parameter: level should be an integer.")
        return 0
    }
    exp, err := strconv.Atoi(args[3])
    if err != nil{
        fmt.Println("Invalid Parameter: <exp> should be an integer")
        return 0
    }
    player := cg.NewPlayer()
    player.Name = args[1]
    player.Level = level
    player.Exp = exp

    err = centerClient.AddPlayer(player)
    if err != nil{
        fmt.Println("Faild adding player", err)
    }
    return 0
}