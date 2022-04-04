package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/kagestonedragon/server/pkg/user"
	"google.golang.org/grpc"
)

//func addUsers(filePath string) error {
//	f, err := os.Open(filePath)
//	if err != nil {
//		return nil
//	}
//	defer f.Close()
//
//	csvReader := csv.NewReader(f)
//	records, err := csvReader.ReadAll()
//	if err != nil {
//		log.Fatal(" for "+filePath, err)
//	}
//}

func getPath() (string, error) {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		return "", errors.New("Input file is missing")
	}

	return args[0], nil
}

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial(":9091", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	client := user.NewGRPCClient(conn, nil, nil)

	request := user.AddUserRequest{
		Name:   "Andrew_Hype55DAGESTAN",
		Active: true,
	}

	u, err := client.AddUser(ctx, &request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("added, id - %d\n", u.Id)
	}

	//path, err := getPath()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Print(path)
	//
	//for {
	//	if err := addUsers(path); err != nil {
	//		log.Print(err)
	//	}
	//
	//	time.Sleep(time.Duration(5) * time.Second)
	//}
}
