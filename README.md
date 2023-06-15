# benchmark_pkg_errors_alternatives

「次なる `pkg/errors` を探して」の発表で利用した、benchmark用コードです

## how to run
```bash
go test -bench .
```

## result
```
➜  benchmark_pkg_errors_alternatives git:(master) ✗ go test -bench .
goos: darwin
goarch: arm64
pkg: example
BenchmarkErrors/errors-stack-10-10              37239439                32.27 ns/op           16 B/op          1 allocs/op
BenchmarkErrors/errors-stack-100-10              1603806               748.6 ns/op            16 B/op          1 allocs/op
BenchmarkErrors/errors-stack-1000-10              166596              7258 ns/op              16 B/op          1 allocs/op
BenchmarkErrors/pkg/errors-stack-10-10           1282000               919.0 ns/op           304 B/op          3 allocs/op
BenchmarkErrors/pkg/errors-stack-100-10           666804              1812 ns/op             304 B/op          3 allocs/op
BenchmarkErrors/pkg/errors-stack-1000-10          132570              8269 ns/op             304 B/op          3 allocs/op
BenchmarkErrors/cockroachdb/errors-stack-10-10            833182              1313 ns/op             416 B/op          7 allocs/op
BenchmarkErrors/cockroachdb/errors-stack-100-10           532465              2230 ns/op             416 B/op          7 allocs/op
BenchmarkErrors/cockroachdb/errors-stack-1000-10          138037              8659 ns/op             416 B/op          7 allocs/op
BenchmarkErrors/goark/errs-stack-10-10                   1332050               898.7 ns/op           304 B/op          3 allocs/op
BenchmarkErrors/goark/errs-stack-100-10                   657888              1796 ns/op             304 B/op          3 allocs/op
BenchmarkErrors/goark/errs-stack-1000-10                  145098              8228 ns/op             304 B/op          3 allocs/op
BenchmarkStackFormatting/errors-%s-stack-10-10          16257858                74.15 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%v-stack-10-10          16024240                75.33 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%+v-stack-10-10         15767274                76.35 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%s-stack-30-10          16150242                74.11 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%v-stack-30-10          15828436                74.76 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%+v-stack-30-10         15564033                76.23 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%s-stack-60-10          15989890                73.88 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%v-stack-60-10          16131850                75.36 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/errors-%+v-stack-60-10         15720841                76.04 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%s-stack-10-10      16112121                74.25 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%v-stack-10-10      15811204                75.70 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%+v-stack-10-10       186741              6406 ns/op            1689 B/op         20 allocs/op
BenchmarkStackFormatting/pkg/errors-%s-stack-30-10      16098081                74.39 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%v-stack-30-10      15691231                76.05 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%+v-stack-30-10        93282             12876 ns/op            3724 B/op         34 allocs/op
BenchmarkStackFormatting/pkg/errors-%s-stack-60-10      16025239                74.73 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%v-stack-60-10      15837942                75.98 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/pkg/errors-%+v-stack-60-10        92730             12917 ns/op            3715 B/op         33 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%s-stack-10-10               1260886               944.5 ns/op           995 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%v-stack-10-10               1264054               951.2 ns/op           994 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%+v-stack-10-10               142141              8395 ns/op            8164 B/op         26 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%s-stack-30-10               1240914               969.2 ns/op          1124 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%v-stack-30-10               1237952               967.9 ns/op          1124 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%+v-stack-30-10                81166             14829 ns/op           17228 B/op         23 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%s-stack-60-10               1246021               968.7 ns/op          1123 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%v-stack-60-10               1239973               965.0 ns/op          1124 B/op         16 allocs/op
BenchmarkStackFormatting/cockroachdb/errors-%+v-stack-60-10                80394             14788 ns/op           17205 B/op         22 allocs/op
BenchmarkStackFormatting/goark/errs-%s-stack-10-10                      16901536                72.41 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%v-stack-10-10                      16336003                73.76 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%+v-stack-10-10                       189067              6378 ns/op            1689 B/op         20 allocs/op
BenchmarkStackFormatting/goark/errs-%s-stack-30-10                      16653030                72.16 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%v-stack-30-10                      16387005                73.20 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%+v-stack-30-10                        94442             12750 ns/op            3725 B/op         34 allocs/op
BenchmarkStackFormatting/goark/errs-%s-stack-60-10                      16457293                71.96 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%v-stack-60-10                      16365942                74.21 ns/op            8 B/op          1 allocs/op
BenchmarkStackFormatting/goark/errs-%+v-stack-60-10                        93728             12762 ns/op            3717 B/op         33 allocs/op
PASS
ok      example 70.215s
```
