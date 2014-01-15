package svm

import (
    "math/rand"
    "sync"
    "time"
)

type Map map[int]int

type MapMutex struct {
    Map
    sync.Mutex
}

func getRandomKey(m Map) int {
    keys := make([]int, 0, len(m))
    for k, _ := range m {
        keys = append(keys, k)
    }
    return keys[rand.Int()%len(keys)]
}

func write(m Map) {
    k := rand.Int()
    m[k] = rand.Int()
}

func read(m Map) int {
    return m[getRandomKey(m)]
}

func writeLocked(m *MapMutex) {
    m.Lock()
    defer m.Unlock()
    k := rand.Int()
    m.Map[k] = rand.Int()
}

func readLocked(m *MapMutex) int {
    m.Lock()
    defer m.Unlock()
    return m.Map[getRandomKey(m.Map)]
}

func withoutLocks(m Map, writeMax int, baseTickTime time.Duration) {
    r := time.Tick(baseTickTime)
    w := time.Tick(baseTickTime * 2)
    n := 0
loop:
    for {
        select {
        case <-r:
            read(m)
        case <-w:
            write(m)
            n += 1
        }
        if n >= writeMax {
            break loop
        }
    }
}

func withLocks(m *MapMutex, writeMax int, baseTickTime time.Duration) {
    quit := make(chan int)
    r := time.NewTicker(baseTickTime)
    w := time.NewTicker(baseTickTime * 2)
    go func() {
        for _ = range r.C {
            readLocked(m)
        }
    }()
    go func() {
        n := 0
        for _ = range w.C {
            writeLocked(m)
            n += 1
            if n >= writeMax {
                break
            }
        }
        quit <- 1
    }()
    <-quit
    w.Stop()
    r.Stop()
}
