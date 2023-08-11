package main

import (
	"context"
	"fmt"
	"jschan/app"
)

func main() {

	client := jschan.NewClient("http://dev-jschan.lan")
	ctx := context.Background()

	loginOptions := &jschan.PostLoginOptions{
		//Credentials for a private dev jschan
		Username: "admin",
		Password: "WjPQLl7mOGsjpRjPezY8FsCdnXI=",
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

}
