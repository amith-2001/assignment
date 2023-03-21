package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	
	
)
var db *sql.DB

type seat struct{
	Id int `json:"Id"`
	Seat_identifier string `json:"seat_identifier"`
	Seat_class string `json:"seat_class"`
	Is_booked bool `json:"is_booked"`
	Min_price string `json:"min_price"`
	Normal_price string `json:"normal_price"`
	Max_price int `json:"max_price"`
	Name string `json:"name"`
	Phone_no string `json:phone_no`




}
func getMySQLDB() *sql.DB{
	db,err := sql.Open("mysql","root:root@tcp(localhost:3306)/flurn?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println("error connecting to db")

	}
	return db
}

func getSeats(w http.ResponseWriter,r *http.Request){

	db = getMySQLDB()
	defer db.Close()
	ss := [] seat{}
	s := seat{}

	rows ,err := db.Query("SELECT * FROM flurn.seats ORDER BY seat_class")
	if err != nil{
		fmt.Println("error fetching")

	}
	
		for rows.Next() {
			err := rows.Scan(&s.Id,&s.Seat_identifier,&s.Seat_class)
			if err != nil{
				fmt.Println("error scanning table ")
			}
			ss = append(ss,s)
			
			
			
		}
		json.NewEncoder(w).Encode(ss)
}




func getSeatPrice(w http.ResponseWriter,r *http.Request){

//do this tomorrow

	// db = getMySQLDB()

	// rows ,err := db.Query("SELECT * FROM flurn.seatspricing")
	// if err != nil{
	// 	fmt.Println("error fetching")

	// }

	// for rows.Next() {
	// 	err := rows.Scan(&s.Id,&s.Seat_identifier,&s.Seat_class)
	// 	if err != nil{
	// 		fmt.Println("error scanning table ")
	// 	}

	
	// }
}

type booking_seat struct{
	Id int `json:"Id"`
	Seat_class string `json:"seat_class"`
	Min_price string `json:"min_price`
	Normal_price string `json:"normal_price`
	Max_price string `json:"max_price`
	
}

func createBooking(w http.ResponseWriter,r *http.Request){

	db = getMySQLDB()
	defer db.Close()
	a:=booking_seat{}
	aa := []booking_seat{}


	// result , err := db.Exec("insert into flurn.seats(id,name,phone_no) values(?,?,?)",s.Id,s.Name,s.Phone_no)

	rows , err := db.Query("SELECT * FROM flurn.seatpricing")
	if err != nil{
		panic(err)

	}

	for rows.Next() {
		err := rows.Scan(&a.Id,&a.Seat_class,&a.Min_price,&a.Normal_price,&a.Max_price)
		if err != nil{
			fmt.Println("error scanning table ")
		}
		aa = append(aa,a)

	
	
	// fmt.Println("successful in posting")


	// fmt.Printf("post request")
}
json.NewEncoder(w).Encode(aa)
}
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/seats",getSeats).Methods("GET")
	r.HandleFunc("/seats/{sid}",getSeatPrice).Methods("GET")
	r.HandleFunc("/booking",createBooking).Methods("GET")


	http.ListenAndServe(":8080",r)

}