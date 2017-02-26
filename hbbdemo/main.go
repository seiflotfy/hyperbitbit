package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"log"

	"github.com/seiflotfy/hyperbitbit"
)

func main() {
	files, err := filepath.Glob(fmt.Sprintf("%s/*", os.Args[1]))
	if err != nil {
		log.Fatalln(err)
		return
	}

	totalUnique := map[string]bool{}
	totalHBB := hyperbitbit.New()

	for _, f := range files {
		f, err := os.Open(f)
		if err != nil {
			log.Fatalln(err)
			return
		}
		reader := bufio.NewReader(f)
		unique := map[string]bool{}

		hbb := hyperbitbit.New()

		for {
			text, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			unique[string(text)] = true
			totalUnique[string(text)] = true
			hbb.Add([]byte(text))
			totalHBB.Add([]byte(text))
		}

		est := hbb.Cardinality()
		ratio := fmt.Sprintf("%.2f%%", 100*(1-float64(len(unique))/float64(est)))
		log.Println("\n\tfile: ", f.Name(), "\n\texact:", len(unique), "\n\testimate:", est, "\n\tratio:", ratio)
	}

	est := totalHBB.Cardinality()
	ratio := fmt.Sprintf("%.2f%%", 100*(1-float64(len(totalUnique))/float64(est)))
	log.Println("\n\ttotal\n\texact:", len(totalUnique), "\n\testimate:", est, "\n\tratio:", ratio)
}
