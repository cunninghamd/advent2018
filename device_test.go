package main

import "testing"

func TestDeviceDay1(t *testing.T) {
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

func TestDeviceDay2(t *testing.T) {
    var d device
    var expect int
    
    // abcdef contains no letters that appear exactly two or three times.
    // bababc contains two a and three b, so it counts for both.
    // abbcde contains two b, but no letter appears exactly three times.
    // abcccd contains three c, but no letter appears exactly two times.
    // aabcdd contains two a and two d, but it only counts once.
    // abcdee contains two e.
    // ababab contains three a and three b, but it only counts once.
    // checksum of 4 * 3 = 12
    expect = 12
    d = device{boxIds: []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}}
    var checksum = d.GetChecksum()
    if checksum != expect {
        t.Errorf("Expected %d, got %d", expect, checksum)
    }
    var expects = "fgij"
    d = device{boxIds: []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}}
    var common = d.GetCommon()
    if common != expects {
        t.Errorf("Expected %s, got %s", expects, common)
    }
}

func TestDeviceDay3(t *testing.T) {
    var d device
    var expect int
    
    // #1 @ 1,3: 4x4
    // #2 @ 3,1: 4x4
    // #3 @ 5,5: 2x2
    // 4 overlapping claims
    expect = 4
    d = device{claims: []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}}
    var claims = d.GetOverlappingClaims()
    if claims != expect {
        t.Errorf("Expected %d, got %d", expect, claims)
    }
}

func TestDeviceDay4(t *testing.T) {
    var d device
    var expect int
    
    //[1518-11-01 00:00] Guard #10 begins shift
    //[1518-11-01 00:05] falls asleep
    //[1518-11-01 00:25] wakes up
    //[1518-11-01 00:30] falls asleep
    //[1518-11-01 00:55] wakes up
    //[1518-11-01 23:58] Guard #99 begins shift
    //[1518-11-02 00:40] falls asleep
    //[1518-11-02 00:50] wakes up
    //[1518-11-03 00:05] Guard #10 begins shift
    //[1518-11-03 00:24] falls asleep
    //[1518-11-03 00:29] wakes up
    //[1518-11-04 00:02] Guard #99 begins shift
    //[1518-11-04 00:36] falls asleep
    //[1518-11-04 00:46] wakes up
    //[1518-11-05 00:03] Guard #99 begins shift
    //[1518-11-05 00:45] falls asleep
    //[1518-11-05 00:55] wakes up
    
    // setting  10 5 25
    // setting  10 30 55
    // setting  99 40 50
    // setting  10 24 29
    // setting  99 36 46
    // setting  99 45 55    
    
    expect = 240
    d = device{logEntries: []string{"[1518-11-01 00:00] Guard #10 begins shift", "[1518-11-01 00:05] falls asleep", "[1518-11-01 00:25] wakes up", "[1518-11-01 00:30] falls asleep", "[1518-11-01 00:55] wakes up", "[1518-11-01 23:58] Guard #99 begins shift", "[1518-11-02 00:40] falls asleep", "[1518-11-02 00:50] wakes up", "[1518-11-03 00:05] Guard #10 begins shift", "[1518-11-03 00:24] falls asleep", "[1518-11-03 00:29] wakes up", "[1518-11-04 00:02] Guard #99 begins shift", "[1518-11-04 00:36] falls asleep", "[1518-11-04 00:46] wakes up", "[1518-11-05 00:03] Guard #99 begins shift", "[1518-11-05 00:45] falls asleep", "[1518-11-05 00:55] wakes up"}}
    var guardMinute = d.GetGuardAndMinute()
    if guardMinute != expect {
        t.Errorf("Expected %d, got %d", expect, guardMinute)
    }
    expect = 4455
    guardMinute = d.GetSleepiestMinute()
    if guardMinute != expect {
        t.Errorf("Expected %d, got %d", expect, guardMinute)
    }
}

// dabAcCaCBAcCcaDA
// dabA--aCBA--caDA
// dab----CBA--caDA

func TestDeviceDay5(t *testing.T) {
    var d device
    var pInputs puzzleInputs
    var expect int
    
    expect = 10
    d = device{polymerLength: 16, polymers: pInputs.Day5Test()}
    var reactions = d.GetReactions("")
    if reactions != expect {
        t.Errorf("Expected %d, got %d", expect, reactions)
    }
    
    expect = 4
    var optimized = d.GetOptimizedReactions()
    if optimized != expect {
        t.Errorf("Expected %d, got %d", expect, optimized)
    }
}
