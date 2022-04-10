package controller

import (
	"fmt"
	"net/http"

	m "../models"
	"github.com/gin-gonic/gin"
)

func DefineVehicles(Weather string) (m.Vehicles) {

	res := m.Vehicles{}

	switch Weather {
	case "Sunny":
		res.Name[0] = "Bike"
		res.Speed[0] = 10
		res.Crater[0] = 20
	case "Tuktuk":
		res.Name[1]= "Tuktuk"
		res.Speed[1] = 10
		res.Crater[1] = 20
	default:
		res.Name[2]= "Car"
		res.Speed[2] = 10
		res.Crater[2] = 20
	}
	return res
}

//
func DefineOrbit( Speed int64, Orbit int64, Crater int64, Weather string)float64 {

	ac_Vehicles := DefineVehicles(Weather)
	MegaMiles := ac_Vehicles.Speed

	var temp_speed int64 
	
	var temp_time float64

	for i := 0; i < len(ac_Vehicles.Name); i++ {
		if Speed <= int64(MegaMiles[i]) {
			temp_speed = Speed
		}else{
			temp_speed = int64(MegaMiles[i])
			temp_time = float64( Orbit ) + ac_Vehicles.Crater[1] * ac_Vehicles.Crater[i] + float64(60/temp_speed)
		}
	}
	return temp_time
}

func DefinePath(c *gin.Context){
	var  l m.Request
	var rb m.Returnblock
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
	}
	Orbit1 := DefineOrbit(l.Orbit1, 18,20,l.Weather )
	Orbit2 := DefineOrbit(l.Orbit2, 20,10,l.Weather )
	fmt.Println("Weather is  ", l.Weather )
	fmt.Println("Orbit1 traffic speed is megamiles/hour",l.Orbit1 )
	fmt.Println("Orbit1 traffic speed is  megamiles/hour",l.Orbit2)

	if (Orbit1 > Orbit2){
		fmt.Println("Vehicle On Orbit 1 Is ")
	}else{
		fmt.Println("Vehicle On Orbit 2 Is ")
	}
}
