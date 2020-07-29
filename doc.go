// Package catsort implements a sorting algorithm that divides a string into
// chunks of runes of the same category and compares them.
//
// 1. Divide the string into chunks of runes of the same category of the three.
//     - Number Category : consecutive runes where unicode.IsDigit(rune) is true.
//     - Letter Category : consecutive runes where unicode.IsLetter(rune) is true.
//     - Symbol Category : other runes.
// 2. Compare categories head to tail.
//     - Symbol < Number < Letter
//         - Place Symbol before Number
//         - Place Number before Letter
//     - Comparing Number to Number is comparing with its numeric value.
//     - Comparing Letter to Letter is comparing as a string value.
//     - Comparing Symbol to Symbol is comparing as a string value.
//
// For example comparing "2020-2-2" to "2020-01-02"
//
//     "2020-2-2"   => Number(2020), Symbol(-), Number(2), Symbol(-), Number(2)
//     "2020-01-02" => Number(2020), Symbol(-), Number(1), Symbol(-), Number(2)
//
//  - The first chunk, Number(2020) is same. The second chunk, Symbol(-) is also same.
//  - The next chunk, comparing Number(2) to Number(1). 1 is placed before 2.
//  - The result is that "2020-01-02" is placed before "2020-2-2"
package catsort
