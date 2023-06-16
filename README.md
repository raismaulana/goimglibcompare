# goimglibcompare

## run:
```
go run . -lib flag
```
flag: stdlib, vips

## test:
```
go test -bench=. -v -benchmem -benchtime 10x
```

## test specific:
```
go test -benchmem -bench ^Benchmark_Stdlib$ goimglibcompare -memprofile stdlib.out -benchtime 1x
go test -benchmem -bench ^Benchmark_Vips$ goimglibcompare -memprofile vips.out -benchtime 1x
```

## profiling:
```
go tool pprof -png stdlib.out > stdlib-profile.png
go tool pprof -png vips.out > vips-profile.png
```