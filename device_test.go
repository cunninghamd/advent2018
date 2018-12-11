package main

import "testing"

func TestDevice(t *testing.T) {
    var d device
    var expect int
    //+1, -2, +3, +1 = 3
    //+1, +1, +1 results in  3
    //+1, +1, -2 results in  0
    //-1, -2, -3 results in -6
    expect = 3
    d = device{frequencies: []int{1, -2, 3, 1}}
    d.Calibrate()
    if d.Frequency() != expect {
        t.Errorf("Expected %d, got %d", expect, d.Frequency())
    }
    expect = 3
    d = device{frequencies: []int{1, 1, 1}}
    d.Calibrate()
    if d.Frequency() != expect {
        t.Errorf("Expected %d, got %d", expect, d.Frequency())
    }
    expect = 0
    d = device{frequencies: []int{1, 1, -2}}
    d.Calibrate()
    if d.Frequency() != expect {
        t.Errorf("Expected %d, got %d", expect, d.Frequency())
    }
    expect = -6
    d = device{frequencies: []int{-1, -2, -3}}
    d.Calibrate()
    if d.Frequency() != expect {
        t.Errorf("Expected %d, got %d", expect, d.Frequency())
    }
    
    //+1, -1 first reaches 0 twice.
    //+3, +3, +4, -2, -4 first reaches 10 twice.
    //-6, +3, +8, +5, -6 first reaches 5 twice.
    //+7, +7, -2, -7, -4 first reaches 14 twice.
    var repeated int 
    expect = 0
    d = device{frequencies: []int{1, -1}}
    repeated = d.CheckRepeats()
    if repeated != expect {
        t.Errorf("Expected %d, got %d", expect, repeated)
    }
    expect = 10
    d = device{frequencies: []int{3, 3, 4, -2, -4}}
    repeated = d.CheckRepeats()
    if repeated != expect {
        t.Errorf("Expected %d, got %d", expect, repeated)
    }
    expect = 5
    d = device{frequencies: []int{-6, 3, 8, 5, -6}}
    repeated = d.CheckRepeats()
    if repeated != expect {
        t.Errorf("Expected %d, got %d", expect, repeated)
    }
    expect = 14
    d = device{frequencies: []int{7, 7, -2, -7, -4}}
    repeated = d.CheckRepeats()
    if repeated != expect {
        t.Errorf("Expected %d, got %d", expect, repeated)
    }
}
