package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Response struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}

// Order represents the model for an order
type Order struct {
	OrderID      int       `json:"orderId" gorm:"primaryKey" `
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}

// Item represents the model for an order
type Item struct {
	ItemID      int    `json:"itemId" gorm:"primaryKey" `
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"-"`
}

var db *gorm.DB
var orders []Order
var prevOrderID = 0
var err error

func dbInit() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Order{}, &Item{})
	fmt.Println("Success Open DB")
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description This api for get all orders from db_product
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {

	var order Order
	w.Header().Set("Content-Type", "application/json")

	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
	}

	db.Create(&order)
	json.NewEncoder(w).Encode(order)
}

// GetOrder godoc
// @Summary Get details for a given orderId
// @Description Get details of order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order"
// @Success 200 {object} Order
// @Router /orders/{orderId} [get]
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID, error := strconv.Atoi(params["orderId"])
	if error != nil {
		log.Fatal(error.Error())
		return
	}
	for _, order := range orders {
		if order.OrderID == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

// UpdateOrder godoc
// @Summary Update order identified by the given orderId
// @Description Update the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be updated"
// @Success 200 {object} Order
// @Router /orders/{orderId} [put]
func updateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder Order
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	params := mux.Vars(r)
	inputOrderID, error := strconv.Atoi(params["orderId"])
	if error != nil {
		log.Fatal(error.Error())
		return
	}
	updatedOrder.OrderID = inputOrderID
	db.Where("order_id = ?", params["orderId"]).Delete(&Item{})
	db.Save(&updatedOrder)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedOrder)

}

// DeleteOrder godoc
// @Summary Delete order identified by the given orderId
// @Description Delete the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be deleted"
// @Success 204 "No Content"
// @Router /orders/{orderId} [delete]
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db.Where("order_id = ?", params["orderId"]).Delete(&Item{})
	db.Where("order_id = ?", params["orderId"]).Delete(&Order{})

	response := Response{
		Code:   "200",
		Status: "Success",
	}
	json.NewEncoder(w).Encode(response)
}

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email testing@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
	dbInit()
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/orders", CreateOrder).Methods("POST")
	// // Read
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	// Read-all
	router.HandleFunc("/orders", getOrders).Methods("GET")
	// Update
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	// Delete
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")
	fmt.Println("STARTING SERVER AT localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
