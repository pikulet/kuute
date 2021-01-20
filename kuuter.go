package main

import (
//    "github.com/ajstarks/svgo"
)

func getKuuteSvg(count int) int {
    numDigits := getNumDigits(count) 
    return numDigits
}

func getNumDigits(count int) (result int) {
    for count != 0 {
        count /= 10
        result += 1
    }
    return result
}
