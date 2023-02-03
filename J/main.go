package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var dictionarySize uint16
	fmt.Fscan(in, &dictionarySize)

	var dictionary = make([]string, dictionarySize)
	for i := uint16(0); i < dictionarySize; i++ {
		var word string
		fmt.Fscan(in, &word)
		dictionary[i] = word
	}

	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j])
	})

	var requests uint16
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

	}

}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
