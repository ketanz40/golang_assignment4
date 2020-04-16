package main

import (
	"fmt"
	"time"
)

var obj1 owner = owner{firstName: "Ketan", lastName: "Joshi", ownersInsurance: insurance{insuranceCompany: "State Farm", 
	phoneNumber: "661-599-3660", priceInsurance: 50000, address: "2702 Eagle Crest Dr."},
	ownerCar: car{price: 30000, manufacturer: "Subaru", model: "Impreza"}}

var carMap map[owner]car = map[owner]car{obj1: obj1.ownerCar}

type car struct{
	price int
	manufacturer string
	model string
}

type owner struct{
	firstName string
	lastName string
	ownersInsurance insurance
	ownerCar car
}

type insurance struct{
	insuranceCompany string
	phoneNumber string
	priceInsurance int
	address string
}

func (i *car) getCarName() string{
    return i.manufacturer + " " + i.model
}

//Print all the data
func query(){
	for i,_ := range carMap{
		fmt.Println(i.firstName, i.lastName,"\n", i.ownerCar.getCarName(),"$",i.ownerCar.price,"\n", i.ownersInsurance.insuranceCompany, 
		"$", i.ownersInsurance.priceInsurance, "\n", i.ownersInsurance.address, i.ownersInsurance.phoneNumber)
		fmt.Println()
	}
}

//Print only cars
func stock() {
	for i,_ := range carMap{
		fmt.Println(i.ownerCar.getCarName())
	}
}

func addNew(fname, lname, oi, pn, ad, manu, mod string, pcar, pi int){
	newObj := owner{firstName: fname, lastName: lname, ownersInsurance: insurance{insuranceCompany: oi, phoneNumber: pn, priceInsurance: pi, address: ad},
	ownerCar: car{price: pcar, manufacturer: manu, model: mod}}
	carMap[newObj] = newObj.ownerCar
}

func deleteCar(fname, lname string){
	for i,_ := range carMap{
		if (fname == i.firstName) && (lname == i.lastName){
			delete(carMap, i)
		}
	}
}

func sell(){
	var fname string
	var lname string
	fmt.Print("Who's car is being sold? ")
	fmt.Scanln(&fname, &lname)
	deleteCar(fname, lname)
}

func buy(){
	var fname, lname, oi, ad, manu, mod, pn string
	var pcar, pi int
	fmt.Print("Who bought a car? ")
	fmt.Scanln(&fname, &lname)
	fmt.Print("What insurance does the owner have? ")
	fmt.Scanln(&oi)
	fmt.Print("What is the owner's address? ")
	fmt.Scanln(&ad)
	fmt.Print("What is the owner's phone number? ")
	fmt.Scanln(&pn)
	fmt.Print("What is the car's name (Manufacturer and Model)? ")
	fmt.Scanln(&manu, &mod)
	fmt.Print("What is the price of the car")
	fmt.Scanln(&pcar) 
	fmt.Print("How much is the owner's insurance? ")
	fmt.Scanln(&pi)
	addNew(fname, lname, oi, pn, ad, manu, mod, pcar, pi)
}

func ticker(){
	ticker := time.NewTicker(1 * time.Second)
	for  _ = range ticker.C {
		fmt.Println("ping")
	}
}


func main(){
	loop:
		for{
			go ticker()   //This can be commented out so that you can the car dealership without the pings interrupting the code and user input
			var userChoice int
			fmt.Print("Press 1 to print a query\nPress 2 to print the stock\nPress 3 to buy a car\n"+
			"Press 4 to sell a car\nPress 5 to quit the program\n")
			fmt.Scanln(&userChoice)
			switch userChoice{
			case 1:
				fmt.Println()
				query()
				fmt.Println()
				break
			case 2:
				fmt.Println()
				stock()
				fmt.Println()
				break
			case 3:
				fmt.Println()
				buy()
				break
			case 4:
				fmt.Println()
				sell()
				break
			case 5:
				fmt.Println("Have a good day!")
				break loop
			default:
				fmt.Println("Invalid Input")
			}
		}
}