package main

import ( "fmt"
		"time"
		"sync"
		"flag"
)

type Cashier struct{
	id int
	status string
}

type Customer struct{
	c_id int
}


type BankSim struct{
	noOfCashiers int
	noOfCustomers int
	timePerCustomer int
	index int
	cashiers []Cashier
	customers []Customer
}

func (c *Cashier) serviceCustomer(bs *BankSim,m *sync.Mutex, c1 chan<- string){

	current := 0
		for {
				m.Lock()	
				current = bs.index
				bs.index+=1
				m.Unlock()
				
				if( current < bs.noOfCustomers){
					c.status = "Occupied"
					str1 := time.Now().Format("2006-01-02 15:04:05")
					str2 := fmt.Sprintf("--> Cashier %d: Customer %d Started",c.id,bs.customers[current].c_id)
					fmt.Println(str1,str2)
					time.Sleep(time.Duration(bs.timePerCustomer) * time.Second)
					str3 := time.Now().Format("2006-01-02 15:04:05")
					str4 := fmt.Sprintf("--> Cashier %d: Customer %d Completed",c.id,bs.customers[current].c_id)
					fmt.Println(str3,str4)
					c.status = "Unoccupied"
				} else {
					c1<-"done"
					break
				}
			}
	}


func main(){


	bs := BankSim{}

	ptr1 := flag.Int("numCashiers",3,"an int")
	ptr2 := flag.Int("numCustomers",100,"an int")
	ptr3 := flag.Int("timePerCustomer",3,"an int")
	flag.Parse()

	bs.noOfCashiers = *ptr1
	bs.noOfCustomers = *ptr2
	bs.timePerCustomer = *ptr3

	c1 := make(chan string,3)
	var m sync.Mutex
	i := 0

	bs.cashiers = make([]Cashier,bs.noOfCashiers)
	for i=0 ; i < bs.noOfCashiers ;i++{
		bs.cashiers[i] = Cashier{i+1,"Unoccupied"} 
	}

	bs.customers = make([]Customer,bs.noOfCustomers)
	for i=0 ; i < bs.noOfCustomers ;i++{
		bs.customers[i] = Customer{i+1} 
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"--> Bank Simulation Started")
	
	for i=0 ; i< bs.noOfCashiers ; i++{
		go bs.cashiers[i].serviceCustomer(&bs,&m,c1)
	}
	
	for i=0 ; i < bs.noOfCashiers  ; i++{
		select{
			case <-c1 :
		}
	}
	
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"--> Bank Simulated Completed")
}
