package shopping

import (
	"errors"
	"strconv"
)

var userList []User = make([]User, 0);
var orderList []Order = make([]Order, 0)
var products map[string]uint = map[string]uint{
	"Milk": 20,
	"Coffee": 30,
}

func GetAllUsers() []User {
	return userList;
}

func GetAllOrders() []Order {
	return orderList;
}

func AddUser(user User) (newUser User){
	var newId uint = uint(len(userList))
	user.Id = &newId
	userList = append(userList, user)
	return user
}

func addNewOrder(order Order) Order {
	var newId uint = uint(len(orderList))
	order.Id = &newId
	orderList = append(orderList, order)
	return order
}

func GetUserById(id uint) (user *User) {
	if len(userList) == 0 || uint(len(userList)) <= id{
		return nil
	}
	if *userList[id].Id == id {
		return &userList[id]
	}
	return nil
}

func AssignOrderToUserById(userId uint, name string, quantity uint) (user *User, err error) {

	var userFound *User = GetUserById(userId)
	
	if userFound == nil {
		return nil, errors.New("User with Id" + strconv.FormatUint(uint64(userId), 10) + "doesn't exist")
	}


	var price uint = products[name]
	if price == 0 {
		return nil, errors.New("Product not found")
	}
	var order = Order{
	}
	order.Name = name
	order.Amount = uint64(price * quantity)

	newOrder := addNewOrder(order)
	if userFound.Orders != nil && len(*userFound.Orders) > 0 {
		*userFound.Orders = append(*userFound.Orders, newOrder)
	} else {
		newOrderList := (append([]Order{}, newOrder))
		userFound.Orders = &newOrderList
	}
	return userFound, nil
}

