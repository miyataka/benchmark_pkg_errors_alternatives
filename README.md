# benchmark_pkg_errors_alternatives

「[次なる `pkg/errors` を探して](https://speakerdeck.com/always_allokay/errorswotan-site)」，「[ブログのリアクションから始めるGoのパフォーマンス分析入門](https://speakerdeck.com/always_allokay/burogunoriakusiyonkarashi-merugonopahuomansufen-xi-ru-men)」の発表で利用した、benchmark用コードです


## how to run
```bash
go test -bench .
```

## result
```
➜  benchmark_pkg_errors_alternatives git:(main) go test -bench .
goos: darwin
goarch: arm64
pkg: example
BenchmarkErrors/errors-stack-10-10              37280496                33.06 ns/op           16 B/op          1 allocs/op
BenchmarkErrors/errors-stack-100-10              1606914               745.9 ns/op            16 B/op          1 allocs/op
BenchmarkErrors/errors-stack-1000-10              167526              7195 ns/op              16 B/op          1 allocs/op
BenchmarkErrors/pkg/errors-stack-10-10           1256511               916.9 ns/op           304 B/op          3 allocs/op
BenchmarkErrors/pkg/errors-stack-100-10           641211              1816 ns/op             304 B/op          3 allocs/op
BenchmarkErrors/pkg/errors-stack-1000-10          145304              8249 ns/op             304 B/op          3 allocs/op
BenchmarkErrors/cockroachdb/errors-stack-10-10            874520              1338 ns/op             416 B/op          7 allocs/op
BenchmarkErrors/cockroachdb/errors-stack-100-10           537822              2210 ns/op             416 B/op          7 allocs/op
BenchmarkErrors/cockroachdb/errors-stack-1000-10          138954              8640 ns/op             416 B/op          7 allocs/op
BenchmarkErrors/goark/errs-stack-10-10                   1771054               680.5 ns/op           648 B/op          7 allocs/op
BenchmarkErrors/goark/errs-stack-100-10                   851818              1399 ns/op             648 B/op          7 allocs/op
BenchmarkErrors/goark/errs-stack-1000-10                  152260              7885 ns/op             648 B/op          7 allocs/op
BenchmarkStackFormatting/errors-%s-stack-10-10          16073302                74.19 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%v-stack-10-10          16066155                75.49 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%+v-stack-10-10         15782058                76.25 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%s-stack-30-10          16099944                74.91 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%v-stack-30-10          16081317                75.46 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%+v-stack-30-10         15475874                76.05 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%s-stack-60-10          16071885                74.00 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%v-stack-60-10          15950331                74.89 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%+v-stack-60-10         15637929                76.25 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%s-stack-10-10      15922783                74.14 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%v-stack-10-10      15663325                75.78 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%+v-stack-10-10       186388              6378 ns/op            1689 B/op         20 allocs/op
BenchmarkStackFormatting/pkg/errors-%s-stack-30-10      15865809                74.45 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%v-stack-30-10      15776956                75.97 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%+v-stack-30-10        93548             12825 ns/op            3724 B/op         34 allocs/op
BenchmarkStackFormatting/pkg/errors-%s-stack-60-10      16115764                74.51 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%v-stack-60-10      15763244                75.55 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%+v-stack-60-10        93242             12849 ns/op            3716 B/op         33 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%s-stack-10-10               1284194               934.8 ns/op           994 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%v-stack-10-10               1278844               938.5 ns/op           994 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%+v-stack-10-10               142902              8307 ns/op            8161 B/op         26 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%s-stack-30-10               1244330               955.0 ns/op          1124 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%v-stack-30-10               1246329               962.1 ns/op          1124 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%+v-stack-30-10                76638             14675 ns/op           17220 B/op         23 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%s-stack-60-10               1249936               951.6 ns/op          1124 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%v-stack-60-10               1245757               961.4 ns/op          1123 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%+v-stack-60-10                80919             14867 ns/op           17222 B/op         22 allocs/op
BenchmarkStackFormatting/goark/errs-%s-stack-10-10                      15291566                78.53 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%v-stack-10-10                      14251922                82.78 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%+v-stack-10-10                       621520              1887 ns/op            1401 B/op         33 allocs/op
BenchmarkStackFormatting/goark/errs-%s-stack-30-10                      15319588                78.71 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%v-stack-30-10                      14493854                82.72 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%+v-stack-30-10                       615178              1884 ns/op            1401 B/op         33 allocs/op
BenchmarkStackFormatting/goark/errs-%s-stack-60-10                      15143666                78.09 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%v-stack-60-10                      14326141                82.58 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%+v-stack-60-10                       613336              1896 ns/op            1401 B/op         33 allocs/op
PASS
ok      example 69.441s
```

## how to run for cpuprofile
```bash
# profile
go test -bench XXX -cpuprofile=c.out

# view on browser profile
go tool pprof -http :8080 -timeout 1 c.out
```
