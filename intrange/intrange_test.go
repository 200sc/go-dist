package intrange

// There was a previous implementation of spread
// which was not a different way of initializing linear
// and this test confirmed that they were equivalent.
//
// func TestAreTheyTheSame(t *testing.T) {
// 	rand.Seed(time.Now().UnixNano())

// 	l := NewLinear(0, 6)
// 	bs := NewSpread(3, 3)

// 	lTot := 0
// 	bsTot := 0

// 	for i := 0; i < 100000; i++ {
// 		lTot += l.Poll()
// 		bsTot += bs.Poll()
// 	}

// 	fmt.Println(lTot / 1000)
// 	fmt.Println(bsTot / 1000)

// 	l = l.Mult(2)
// 	bs = bs.Mult(2)

// 	lTot = 0
// 	bsTot = 0

// 	for i := 0; i < 100000; i++ {
// 		lTot += l.Poll()
// 		bsTot += bs.Poll()
// 	}

// 	fmt.Println(lTot / 1000)
// 	fmt.Println(bsTot / 1000)
// }
