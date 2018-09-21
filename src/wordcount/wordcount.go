package wordcount

import	(
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pair struct {
	Key    string
	Value  int
}

// PariList implement the sort interface, and sort it with sort.Sort

type PairList []Pair

func (p PairList) Swap(i, j int) 	{ p[i], p[j] = p[j], p[i] }

func (p PairList) Len() int	{ return len(p)	}

func (p PairList) Less(i, j int) bool {	return p[j].Value < p[i].Value }		
// descending

//catch the word
func SplitOnNonLetters(s string) []string  {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

/*
	Implement the type WordCount base on map, and implement the function:
	Merge(), Report(), SortReport(), UpdateFreq(), WordFreqCounter()
*/

type WordCount map[string]int

// For merge two WordCount
func (source WordCount) Merge(wordcount WordCount) WordCount {
	for k, v := range wordcount {
		source[k] += v
	}

	return source
}

// Print the condition of word frequency staticstic
func (wordcount WordCount) Report() {
	words := make([]string, 0, len(wordcount))
	wordWidth, frequency := range wordcount {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, wordcount[word])
	}
}

// print word frequency from more to fewer
func (wordcount WordCount) SortReport() {
	p := make(PairList, len(wordcount))
	i := 0
	for k, v := range wordcount { // Transform PairList to wordcount map
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	wordWidth, frequencyWidth := 0, 0
	for _, pair :=  range p {
		word, frequency := pair.Key, pair.Value
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")

	for _, pair := range p {
		fmt.Printf("%-*s %*d\n", wordWidth, pair.Key, frequencyWidth, pair.Value)
	}
}

// Read words from a file and update their occurrences
func (wordcount WordCount) UpdateFreq(filename string) {
	var file *os.File
	var err error

	if file, err = os.Open(filename); err != nil {
		log.Printf("filed to open the file: ", err)
		return 
	}
	defer file.Close()	// close file before this function exits

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				wordcount[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
	}
}

// Concurrent statistics of word frequency
func (wordcount WordCount) WordFreqCounter(file []string)  {
	
	result := make(chan Pair, len(files))	// goroutine will send the result to this channel
	done := make(chan struct{}, len(files))	// Each time goroutine finished its work, it will send a empty structure to this channel, which means the work is done

	for i := 0; i < len(files); {	// boost a goroutine while open a file through anomonous way
		go func(done chan<- struct{}, results chan<- Pair, filename string) {
			wordcount := make(WordCount)
			wordcount.UpdateFreq(filename)
			for k, v := range wordcount {
				pair := Pair{k, v}
				result <- pair
			}
			done <- struct{}{}
		}(done, results, files[i])

		i++
	}

	for working := len(files); working > 0; {	// monitor the routine and not exit until all works' goroutine finished its task
		select {
		case pair := <-results:		// receive the statistics result which was sent to the routine
			wordcount[pair.Key] += pair.Value
		
		case <-done:	// check if all the goroutine finished its work
			working--
		}
	}

DONE:	// Reboost for-loop to handle the value which are not done
	for {
		select {
		case pair := <-results:
			wordcount[pair.Key]	+= pair.Value
		default:
			break Done 
		}
	}

	close(results)
	close(done)
}


