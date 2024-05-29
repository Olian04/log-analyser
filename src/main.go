package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/sourcegraph/conc/pool"
	"golang.org/x/exp/maps"
)

func main() {
	file, err := os.Open("./data/out")
	if err != nil {
		panic("Unable to open data file")
	}
	defer file.Close()

	p := pool.NewWithResults[[]string]().WithMaxGoroutines(1000)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		p.Go(func() []string {
			json, err := gabs.ParseJSON(line)
			if err != nil {
				panic(err)
			}
			return maps.Keys(json.Data().(map[string]interface{}))
		})
	}

	occurrences := map[string]int{}
	total_consumed := 0

	resp := p.Wait()
	for _, keys := range resp {
		if len(keys) == 0 {
			continue
		}
		total_consumed += 1
		for _, key := range keys {
			occurrences[key] += 1
		}
	}

	for k, v := range occurrences {
		percentage := (float32(v) / float32(total_consumed))
		fmt.Printf("%v %v\n", percentage, k)
	}
}
