package main

import "fmt"

type gasEngine struct{
    mpg uint8
    gallons uint8
}

type electricEngine struct{
    mpkwh uint8
    kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
    return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
    return e.kwh*e.mpkwh
}

type engine interface{
    milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
    if miles<=e.milesLeft(){
        fmt.Println("You can make it there!")
    }else{
        fmt.Println("Need to fuel up first!")
    }
}

func main() {
    // Create some example engines
    gas := gasEngine{mpg: 30, gallons: 10}
    electric := electricEngine{mpkwh: 4, kwh: 100}
    
    // Test both engines
    canMakeIt(gas, 250)      // Test gas engine
    canMakeIt(electric, 255) // Test electric engine
}