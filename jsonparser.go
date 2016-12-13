package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	var f interface{}

	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	m := f.(map[string]interface{})
	parseObject("", m)
}

func parseObject(prefix string, m map[string]interface{}) {
	for k, v := range m {
		parse(prefix+"/"+k, v)
	}
}

func parseArray(prefix string, m []interface{}) {
	for i, v := range m {
		parse(prefix+"["+strconv.Itoa(i+1)+"]", v)
	}
}

func parse(key string, value interface{}) {
	switch v := value.(type) {
	case string, float64, bool:
		fmt.Println(key, v)
	case map[string]interface{}:
		parseObject(key, v)
	case []interface{}:
		parseArray(key, v)
	default:
	}
}
