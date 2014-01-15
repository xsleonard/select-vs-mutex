select-vs-mutex
===============

Golang comparison of a lock-free select to a mutex

This is a companion repo to [Avoiding Locks in Golang](https://xsleonard.github.io/2014/01/15/avoiding-locks-golang/)

Benchmarks
----------

```
go test -bench=With
```

will run the benchmarks comparing mutex (WithLock) and select (WithoutLock).

```
go test -bench="(Read)|(Write)"
```

will run the benchmarks for the lock-free read and write methods.
