package main

import "fmt"

type device struct {
    frequencies []int
    frequency int
    found []int
}

func (d *device) Calibrate() {
    var f int = 0
    
    //for index, element := range someSlice {
    for _, element := range d.frequencies {
        f += element
    }
    
    d.SetFrequency(f)
}

func (d *device) Contains(f int) {
    for _, v := range d.found {
        if v.Equals(f) {
            return true
        }
    }
    
    return false
}

func (d *device) CheckRepeats() {
    var f int = 0
    var floop int = 0
    d.found = append(d.found, f)
    
    for floop < 1 {
        for _, element := range d.frequencies {
            f += element
            if (d.Contains(f)) {
                floop = 1
            } else {
                d.found = append(d.found, f)
            }
        }
    }
    
    return f
}

func (d *device) SetFrequency(frequency int) {
    d.frequency = frequency
}

func (d device) Frequency() int {
    return d.frequency
}

func main() {
    d := device{frequencies: []int{-15, -17, -16, -12, -4, 10, 10, -19, -18, -6, 4, 4, 5, -11, -4, -12, 1, -8, -10, -16, -18, 17, -16, -7, -5, 2, -4, -1, 13, -19, 12, 19, 5, 11, 1, -16, -8, 19, 15, -14, 1, -12, -17, 10, -3, 5, 7, 4, -10, 5, -7, 15, 11, 10, 19, -17, -3, 14, 8, -14, 7, 11, 16, -7, -1, 6, 4, -1, 15, 17, -10, 4, 16, -18, -10, -2, -18, -13, 15, -7, -14, -12, -2, 1, 7, 5, -1, -19, 1, 9, -20, -9, 1, 16, -7, 11, 14, -24, 4, -3, 21, -17, -10, -44, 13, 2, -9, -15, -17, -8, 13, 4, -1, 4, 1, 2, 6, -18, 1, 19, 4, 7, -15, -12, 4, 17, -3, -15, -19, -19, -17, 16, 17, -19, -3, -9, 4, 12, 12, 12, -21, 17, -19, -4, 16, -4, -18, -16, 15, 14, 12, -6, 16, -1, 16, 6, -4, 6, -4, 11, 20, -7, 12, 51, 19, 15, -2, 12, -9, -6, 1, 10, 1, 19, 8, 3, -1, -19, -2, 4, 11, 8, 2, 10, 14, 4, -8, 9, -13, 18, 5, -8, 15, -19, -13, -13, 10, -4, -16, 17, -3, 7, 20, 18, -12, 13, -10, 14, 19, 17, 7, 7, 3, 17, -8, -4, 1, -8, -12, -1, -15, 9, -7, 6, 10, 4, 4, -14, -9, -19, 13, -5, 19, -17, -1, 5, -13, 2, 4, 15, 8, 10, 8, 17, 17, 1, 6, -15, 5, 12, 11, -3, 12, -6, -17, -17, 10, -15, -5, -1, -4, -3, 5, -21, -2, -18, 3, 19, 10, 15, -5, 17, -14, 6, 7, 3, -11, 6, 20, -9, 6, 6, -10, -9, -8, 16, 19, -9, 5, 6, -3, 12, -8, 7, -1, 16, 17, -19, 3, 11, 18, 19, -16, 11, 11, 5, 11, -25, 13, -7, 9, -18, -16, 18, 22, 35, 82, 8, 11, 19, 17, -11, 19, -16, 1, 21, 17, 18, -9, -10, 8, -15, 4, 20, -19, -10, 2, 22, 11, -5, 16, -7, -7, -15, -11, 28, 13, -7, 1, 3, 11, -7, 3, -15, 34, 5, -22, 7, -15, -2, 29, -45, -51, -6, 10, 6, 22, -17, -24, -12, 38, -71, 46, -10, -19, -15, -115, -26, -17, -2, 4, -15, 9, -4, -10, -14, -4, -18, -18, 21, 8, 13, 15, -13, -6, -6, 7, 4, 12, -7, -20, -14, -22, -2, -16, -7, -2, 12, -24, -18, 35, -45, 128, -2, 22, -10, -17, 14, -4, -11, 13, -6, 29, -40, -11, 7, -25, 5, -10, 58, -171, -169, -93, -106, -61611, -19, -3, -3, 14, 1, -6, 11, -7, -16, -13, -7, -14, 12, 4, -8, -12, -2, -3, 14, -18, -10, 3, -7, -15, -3, -13, 3, 1, -15, -5, -11, 15, 14, -9, 6, 2, 13, 5, 15, -9, 5, 2, 15, 18, 16, -17, -5, -18, 1, 6, -25, -17, 5, 13, -14, 7, 9, -11, -19, -13, -15, 9, -10, 5, 7, 6, -16, -11, 1, 1, 15, -18, -9, -1, 6, -17, -5, -19, 15, 7, -15, 7, 19, -7, -10, -5, -1, 17, 15, -3, 14, 5, -7, -8, -14, -8, -23, -18, -14, 11, 1, 7, -18, 1, -5, 6, -9, -7, -19, -2, 17, -2, 8, -10, 7, -22, -14, -10, -8, 17, -19, -12, 15, 9, -16, 9, 16, -5, -18, -4, -7, -4, -17, 14, 1, 8, -10, 3, -19, 7, -11, 14, 12, 12, -4, -19, -10, -8, -1, -16, -2, -7, -15, -5, -2, -15, 11, -17, -1, -5, 1, -9, -3, -13, -14, 8, -17, -4, -4, -12, 3, 1, 19, 2, 2, 5, -2, -9, 19, 17, -2, -19, 13, 12, 16, 10, -11, 12, -5, 10, 5, -7, 18, 11, 1, -15, 6, -10, 17, -19, 18, 8, 11, -2, 19, -21, -18, -18, -15, 4, 18, 13, -4, 14, 19, -7, 9, 7, 21, 18, -4, 18, 15, -6, 16, 8, -12, 3, -18, 4, 20, -8, -9, 15, 12, 1, 17, 6, -9, 13, -12, -13, 10, -6, -1, 17, 12, -18, 4, 5, 11, -6, 1, 14, -17, 4, 17, 9, -15, -8, 7, 5, -7, 14, 5, 13, 7, -1, -7, -10, 14, -6, 26, 15, 1, 13, 13, -11, -6, 16, -2, -19, 8, -14, 7, -16, 21, 17, 3, 21, 4, -8, -2, 28, 1, -17, 14, -9, 2, -18, 12, 10, 18, -17, -15, -18, 2, 1, -31, 8, -20, -2, -20, -21, 9, 8, 10, -45, 20, 27, 10, -65, 12, -17, -9, -11, 24, 7, -46, -13, 12, -19, 1, -25, 13, -5, -19, 9, -14, -6, 14, 4, -22, -5, -2, -7, 34, 13, -17, -8, 19, 22, -13, 6, 6, 6, -16, 25, 12, 8, -16, -32, -12, -36, -21, 1, 13, -17, -11, 9, 13, 7, -11, -13, -10, -14, 18, 4, -25, 4, -12, 6, -20, -10, 17, -11, 3, -10, -4, -10, 9, -15, -6, 17, 13, -28, -8, 4, -20, 13, -6, 12, -17, -9, -14, 13, -4, 57, -5, 37, -14, 6, -15, -16, -2, -31, 3, 88, 8, -22, 6, 22, 17, 9, -31, -121, -45, 21, -82, -37, -93, -31, -17, 117, -61928, 11, -17, -10, 9, 10, -18, 10, -7, -6, -5, -4, 19, -12, 17, 9, -16, 1, 10, -12, -6, 19, -4, -11, -15, -7, 13, -10, 2, -15, 2, -8, 17, 9, 4, -8, -19, 7, -8, -12, -1, -5, 2, 6, -15, 8, 11, 13, -18, -8, -11, -16, 13, 6, -18, -10, 19, 2, -8, -10, -13, -19, 15, 8, -5, 19, 3, 12, -16, -5, -5, -18, -11, -7, 4, -5, 15, -1, 10, -6, -5, -16, 9, 10, 16, 14, -15, 10, -18, -11, -13, -11, -2, -9, -17, 11, 8, 125143}}
    d.Calibrate()
    fmt.Println("resulting frequency: ", d.Frequency())
}
