func score(input []float64) []float64 {
    var var0 []float64
    if (input[2]) <= (2.6) {
        var0 = []float64{1.0, 0.0, 0.0}
    } else {
        if (input[2]) <= (4.8500004) {
            if (input[3]) <= (1.6500001) {
                var0 = []float64{0.0, 1.0, 0.0}
            } else {
                var0 = []float64{0.0, 0.3333333333333333, 0.6666666666666666}
            }
        } else {
            if (input[3]) <= (1.75) {
                var0 = []float64{0.0, 0.42857142857142855, 0.5714285714285714}
            } else {
                var0 = []float64{0.0, 0.0, 1.0}
            }
        }
    }
    return var0
}
