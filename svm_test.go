package svm

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

var writeMax = 100

func benchmarkWithLocks(b *testing.B, baseTickTime time.Duration) {
    mm := &MapMutex{nil, sync.Mutex{}}
    for i := 0; i < b.N; i++ {
        mm.Map = make(Map, writeMax)
        mm.Map[0] = 0 // added to avoid checking nonempty map in read()
        withLocks(mm, writeMax, baseTickTime)
        if len(mm.Map) != writeMax+1 {
            fmt.Printf("Incomplete Write! Only %d\n", len(mm.Map))
        }
    }
}

func benchmarkWithoutLocks(b *testing.B, baseTickTime time.Duration) {
    for i := 0; i < b.N; i++ {
        m := make(Map, writeMax)
        m[0] = 0 // added to avoid checking nonempty map in read()
        withoutLocks(m, writeMax, baseTickTime)
        if len(m) != writeMax+1 {
            fmt.Printf("Incomplete Write! Only %d\n", len(m))
        }
    }
}

func BenchmarkWithLocks10us(b *testing.B) {
    benchmarkWithLocks(b, time.Microsecond*10)
}

func BenchmarkWithoutLocks10us(b *testing.B) {
    benchmarkWithoutLocks(b, time.Microsecond*10)
}

func BenchmarkWithLocks100us(b *testing.B) {
    benchmarkWithLocks(b, time.Microsecond*100)
}

func BenchmarkWithoutLocks100us(b *testing.B) {
    benchmarkWithoutLocks(b, time.Microsecond*100)
}

func BenchmarkWithLocks1000us(b *testing.B) {
    benchmarkWithLocks(b, time.Microsecond*1000)
}

func BenchmarkWithoutLocks1000us(b *testing.B) {
    benchmarkWithoutLocks(b, time.Microsecond*1000)
}

func BenchmarkRead(b *testing.B) {
    m := make(Map, writeMax)
    for i := 0; i < writeMax; i++ {
        write(m)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        read(m)
    }
}

func BenchmarkWrite(b *testing.B) {
    m := make(Map, writeMax)
    for i := 0; i < b.N; i++ {
        write(m)
    }
}
