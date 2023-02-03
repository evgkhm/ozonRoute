//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//
//	in := bufio.NewReader(os.Stdin)
//	rymMap := make(map[string][]string)
//	var anyWord [2]string
//	var libSize int
//	fmt.Fscan(in, &libSize)
//	for i := 0; i < libSize; i++ {
//		var word string
//		fmt.Fscan(in, &word)
//		if i < 2 {
//			anyWord[i] = word
//		}
//		origWord := word
//		for j := len(word) - 1; j >= 0; j-- {
//			_, ok := rymMap[word[j:]]
//			if !ok {
//				var libArr []string
//				libArr = append(libArr, origWord)
//				rymMap[word[j:]] = libArr
//			} else {
//				tempLib := rymMap[word[j:]]
//				tempLib = append(tempLib, origWord)
//				rymMap[word[j:]] = tempLib
//			}
//		}
//	}
//	var checkSize int
//	fmt.Fscan(in, &checkSize)
//	var out string
//	for k := 0; k < checkSize; k++ {
//		var word string
//		fmt.Fscan(in, &word)
//		out = anyWord[1]
//		for l := len(word) - 1; l >= 0; l-- {
//			val, ok := rymMap[word[l:]]
//			if ok {
//				for m := 0; m < len(val); m++ {
//					if word != val[m] {
//						out = val[m]
//						break
//					}
//				}
//			}
//		}
//
//		fmt.Println(out)
//	}
//}

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var dictionarySize uint16
	fmt.Fscan(in, &dictionarySize)

	dictionary := make(map[string]string)
	var firstSecondWords []string
	for i := uint16(0); i < dictionarySize; i++ {
		var word, tmpWord string
		fmt.Fscan(in, &word)
		tmpWord = word
		//сохранение первых двух слов
		if i == 0 {
			firstSecondWords = append(firstSecondWords, word)
		} else if i == 1 {
			firstSecondWords = append(firstSecondWords, word)
		}
		//запись в мапу суффиксов (ключ), слов (значение)
		for len(tmpWord) >= 2 {
			tmpWord = trimFirstRune(tmpWord)
			dictionary[tmpWord] = word
		}
	}
	var requests uint16
	fmt.Fscan(in, &requests)
	for i := uint16(0); i < requests; i++ {
		var word, tmpWord, res string
		fmt.Fscan(in, &word)
		tmpWord = word
		if word != firstSecondWords[0] {
			res = firstSecondWords[0]
		} else {
			res = firstSecondWords[1]
		}

		for len(tmpWord) >= 1 {
			tmpWord = trimFirstRune(tmpWord)
			if _, ok := dictionary[tmpWord]; ok &&  {
				res = dictionary[tmpWord]
			}
		}
		fmt.Fprintln(out, res)
	}
	/*var requests uint16
	fmt.Fscan(in, &requests)

	for i := uint16(0); i < requests; i++ {
		m := make(map[string]string)

		var word, newWord string
		fmt.Fscan(in, &word)
		newWord = word
	L1:
		for len(newWord) >= 1 {
			for _, val := range dictionary {
				if strings.HasSuffix(val, newWord) && word != val {
					m[val] = newWord
					break L1
				}
			}
			newWord = trimFirstRune(newWord)
		}

		if len(m) == 0 {
			fmt.Fprintln(out, dictionary[0]) //вывод любого сова
		} else {
			for key, _ := range m {
				fmt.Fprintln(out, key)
			}
		}

	}*/

}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
