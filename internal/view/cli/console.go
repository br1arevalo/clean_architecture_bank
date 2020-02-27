package cli

import (
	"errors"
	"fmt"
	"log"
	"os"

	helper "github.com/maestre3d/bank-account/internal/view/cli/helper"
	presenter "github.com/maestre3d/bank-account/internal/view/cli/presenter"
)

var (
	userPresenter    *presenter.UserPresenter
	productPresenter *presenter.ProductPresenter
)

func Start() {
	userPresenter = presenter.NewUserPresenter()
	productPresenter = presenter.NewProductPresenter()
	showMainMenu()
}

func showMainMenu() {
	var selectedOption string
	fmt.Println("Welcome to Foo Bank")
	fmt.Println("1. User\n2. Product\n3. Store\nPress any key to exit")
	_, err := fmt.Scan(&selectedOption)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	switch selectedOption {
	case "1":
		showUserMenu()
		fmt.Println("User created")
		break
	case "2":
		showProductMenu()
		fmt.Println("Product created")
	case "3":
		showStoreMenu()
	default:
		os.Exit(1)
	}
}

func showUserMenu() {
	var selectedOption string
	fmt.Println("User Operations\n1. Create\n2. Delete\n3. Show all\nPress any key to exit")
	_, err := fmt.Scan(&selectedOption)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	switch selectedOption {
	case "1":
		createUser()
		showUserMenu()
		break
	case "3":
		showAllUsers()
		showUserMenu()
		break
	default:
		showMainMenu()
		break
	}
}

func showProductMenu() {
	var selectedOption string
	fmt.Println("Product Operations\n1. Create\n2. Delete\n3. Show all\nPress any key to exit")
	_, err := fmt.Scan(&selectedOption)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	switch selectedOption {
	case "1":
		createProduct()
		showProductMenu()
		break
	case "3":
		showAllProducts()
		showProductMenu()
		break
	default:
		showMainMenu()
		break
	}
}

func showStoreMenu() {
	var selectedOption string
	fmt.Println("Welcome to Foo Fighters Store\nChoose for")
	fmt.Println("1. View products\n2. Buy products\n3. Go back to main menu\nPress any key to exit")
	_, err := fmt.Scan(&selectedOption)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	switch selectedOption {
	case "1":
		showAllProducts()
		showStoreMenu()
	case "2":
		showBuyProducts()
	case "3":
		showMainMenu()
	default:
		os.Exit(1)
	}

}

func createUser() {
	var name string
	fmt.Println("Enter your full name:")
	_, err := fmt.Scan(&name)
	helper.IsValid(err)

	if name == "" {
		helper.IsValid(errors.New("invalid name"))
	}

	helper.IsValid(userPresenter.CreateUser(name))
}

func showAllUsers() {
	users := userPresenter.GetAll()
	for i := range users {
		fmt.Printf("User %d: \n%v\n", i, users[i])
	}
}

func createProduct() {
	var name string
	var price float64
	fmt.Println("Enter the name of the product")
	fmt.Scan(&name)
	fmt.Println("Enter the price for the new product")
	fmt.Scan(&price)

	helper.IsValid(productPresenter.CreateProduct(name, price))
}

func showAllProducts() {
	products := productPresenter.GetAll()
	for i := range products {
		fmt.Printf("Product %d: \n%v\n", i, products[i])
	}
}

func showBuyProducts() {
	var id string
	showAllProducts()
	fmt.Println("Please select the product you want to buy or press b to go back\nEnter the ID of the product")
	fmt.Scan(&id)
	if id == "b" {
		showStoreMenu()
	}
}
