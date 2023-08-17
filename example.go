package main

import (
	"context"
	"fmt"
	"os"

	jschan "jschan_go_api/app"

	"github.com/joho/godotenv"
)

func main() {


	_ = godotenv.Load()	// load .env file locally
	board_password := os.Getenv("PASSWORD") // get value from .env
	b_endpoint := os.Getenv("ENDPOINT")
	
	// test
	println("endpoint: ",b_endpoint)
	//println("password: ",board_password)


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

	//GetBanners(*client,ctx)
	//GetBoards(*client,ctx,"b")
	//GetBoardIndexPages(*client,ctx,"b")

	GetRecentPosts(*client,ctx,"b")


}

func GetReports(ip string, board_name string,client jschan.Client, ctx context.Context) {
	manageReportsOptions := &jschan.GetManageReportsOptions{
		Page:  0,
		IP:    ip,
		Board: board_name,
	}
	reports, err2 := client.GetManageReports(ctx, manageReportsOptions)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Printf("Fetched %d reports\n", len(reports.Reports))

}


// stuff only
func GetRecentPosts(client jschan.Client,  ctx context.Context, board_name string) {

	getRecentOptions := &jschan.GetManageRecentOptions{
		Board: board_name,
	}

	recent, err := client.GetManageRecent(ctx, getRecentOptions)
	if err != nil {
		fmt.Println(err)
		return
	  }
	  fmt.Printf("%+v", recent)


}


// Get /b/
func GetBoards(client jschan.Client,  ctx context.Context, search_request string) {

	getBoardsPublicOptions := &jschan.GetBoardsPublicOptions{
		Search:    search_request,
	  }
	  boards, err := client.GetBoardsPublic(ctx, getBoardsPublicOptions)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  fmt.Printf("%+v", boards)
}


func GetBoardIndexPages(client jschan.Client, ctx context.Context, board_name string) {


	getIndexOptions := &jschan.GetIndexOptions{
		Board: board_name,
		Page:  1,
	  }
	  index, err := client.GetIndex(ctx, getIndexOptions)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  fmt.Printf("%+v", index)
}

func GetAllThreadsBoard(client jschan.Client, ctx context.Context, board_name string) {

	getCatalogOptions := &jschan.GetCatalogOptions{
		Board: board_name,
	  }
	  catalog, err := client.GetCatalog(ctx, getCatalogOptions)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  fmt.Printf("%+v", catalog)
}


func GetCustomPage(client jschan.Client, ctx context.Context,board_name string,page_name string) {

	getCustomPageOptions := &jschan.GetCustomPageOptions{
		Board: board_name,
		CustomPage: page_name,
	  }
	  custompage, err := client.GetCustomPage(ctx, getCustomPageOptions)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  fmt.Printf("%+v", custompage)
}


func GetBanners(client jschan.Client, ctx context.Context, board_name string) {

	getBannersOptions := &jschan.GetBannersOptions{
		Board: board_name,
	  }
	  banners, err := client.GetBanners(ctx, getBannersOptions)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  fmt.Printf("%+v", banners)
}


