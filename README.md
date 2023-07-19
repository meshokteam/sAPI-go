# sAPI-go
Go библиотека для использования API для продавцов аукциона Мешок

```go
package main

import (
	"fmt"
	"log"

	"github.com/your-username/meshokapi"
)

func main() {
	// Replace "your-token" with your actual Meshok API token
	api := meshokapi.NewMeshokAPI("your-token")

	// Get item list
	itemList, err := api.getItemList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(itemList)

	// Get account info
	accountInfo, err := api.getAccountInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accountInfo)

	// Get item info
	itemInfo, err := api.getItemInfo("item-id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(itemInfo)

	// List an item
	params := url.Values{
		"name":        []string{"Item Title"},
    //More params
	}
	listItemResponse, err := api.listItem(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(listItemResponse)

	// Update an item
	updateParams := url.Values{
		"id":          []string{"item-id"},
    //More params
	}
	updateItemResponse, err := api.updateItem(updateParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateItemResponse)

	// Delete an item
	deleteItemResponse, err := api.deleteItem("item-id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deleteItemResponse)
}
```

Make sure to replace "your-token" with your actual Meshok API token and "your-username" with your GitHub username or any other appropriate package name.
