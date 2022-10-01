package controllers

import(
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

type Car struct{
  CarId string `json:"car_id"`
  Brand string `json:"brand"`
  Model string `json:"model"`
  Price int `json:"price"`
}

var CarData = []Car{}

func CreateCar(ctx *gin.Context){
  var newCar Car

  if err := ctx.ShouldBindJSON(&newCar); err != nil{
    ctx.AbortWithError(http.StatusBadRequest, err)
    return
  }
  newCar.CarId = fmt.Sprintf("c%d", len(CarData)+1)
  CarData = append(CarData, newCar)

  ctx.JSON(http.StatusCreated, gin.H{
    "car": newCar,
  })
}


func UpdateCar(ctx *gin.Context){
  carId := ctx.Param("carId")
  condition := false
  var updateCar Car

  if err := ctx.ShouldBindJSON(&updateCar); err != nil{
    ctx.AbortWithError(http.StatusBadRequest, err)
    return
  }

  for i, car := range CarData{
    if carId == car.CarId{
     condition = true
     CarData[i] = updateCar
     CarData[i].CarId = carId
     break
    }
  }

  if !condition{
    ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
      "error_status": "Data Not Found",
      "error_message": fmt.Sprintf("car with id %v not found", carId),
    })
    return
  }
  ctx.JSON(http.StatusOK, gin.H{
    "message": fmt.Sprintf("car with id %v has been successfully update", carId),
  })
}

func GetCar(ctx *gin.Context){
  carId := ctx.Param("carId")
  condition := false
  var carData  Car

  for i, car := range CarData{
    if carId == car.CarId{
      condition = true
      carData = CarData[i]
      break
    }
  }

  if !condition{
    ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
      "error_status": "Data Not found",
      "error_message": fmt.Sprintf("car with id %v not found", carId),
    })
    return
  }

  ctx.JSON(http.StatusOK, gin.H{
    "car": carData,
  })
}


func DeleteCar(ctx *gin.Context){
  carId := ctx.Param("carId")
  condition := false
  var carIndex int

  for i, car := range CarData{
    if carId == car.CarId{
      condition = true
      carIndex = i
      break
    }
  }

  if !condition{
    ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
      "error_status": "Data Not Found",
      "error_message": fmt.Sprintf("car with id %v not found", carId),
    })
    return
  }

  copy(CarData[carIndex:], CarData[carIndex+1:])
  CarData[len(CarData)-1] = Car{}
  CarData = CarData[:len(CarData)-1]

  ctx.JSON(http.StatusOK, gin.H{
    "message": fmt.Sprintf("car with id %v has been successfully deleted", carId),
  })
}
