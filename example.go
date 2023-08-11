package main

import (
	"context"
	"fmt"
	jschan "jschan/app"
	"os"

	"github.com/joho/godotenv"
)

func main() {


	_ = godotenv.Load()
	board_password := os.Getenv("PASSWORD")
	b_endpoint := os.Getenv("ENDPOINT")
	
	// test
	println("endpoint: ",b_endpoint)
	println("password: ",board_password)


	//client := jschan.NewClient("http://dev-jschan.lan")
	client := jschan.NewClient(b_endpoint)
	ctx := context.Background()
	//client.BaseURL

	loginOptions := &jschan.PostLoginOptions{
		//Credentials for a private dev jschan
		Username: "admin",
		Password: board_password,
	}
	err := client.Login(ctx, loginOptions)
	if err != nil {
		fmt.Println(err)
		return
	}
	if client.SessionCookie != "" {
		fmt.Printf("Logged in as user %s\n", loginOptions.Username)
		if _, err := client.GetCSRFToken(ctx); err != nil {
			fmt.Println(err)
			return
		}
	}

	/*
	manageReportsOptions := &jschan.GetManageReportsOptions{
		Page:  0,
		IP:    "10.0.0.192",
		Board: "",
	}
	reports, err2 := client.GetManageReports(ctx, manageReportsOptions)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Printf("Fetched %d reports\n", len(reports.Reports))

	*/

}
