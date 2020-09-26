package main

import "fmt"

type (
    // Values defines interface to be used in mocking
    Values interface {
        area() int
        perimeter() int
    }

    // Measurements defines the measurements used to calculate
    // the  perimeter and area of a rectangle
    Measurements struct {
        length int
        breadth  int        
    }
)

// area calculates and returns the area of a rectangle
func (config *Measurements) area() int {
    return config.length * config.breadth
}

// perimeter calculates and returns the perimeter of a rectangle
func (config *Measurements) perimeter() int {
    return 2*(config.length+config.breadth)
}

// GetareaAndperimeter fetches and returns the area and the perimeter
func GetareaAndperimeter(val Values) (int, int) {
    return val.area(), val.perimeter()
}

func main() {
    var (
        perimeter, area int

        data = &Measurements{
            length: 1,
            breadth: 5,
            
        }
    )

    area, perimeter = GetareaAndperimeter(data)

    fmt.Printf("area : %d , perimeter : %d \n", area, perimeter)
}
