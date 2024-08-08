package main

import "fmt"

func DoSomeThing(body interface{}) (interface{}, error) {
	var ret interface{}
	var err error
	// ...
	return ret, err
}

func SimpleRetry(body interface{}) (resp interface{}, err error) {
	for i := 0; i < 3; i++ {
		resp, err = DoSomeThing(body)
		if err == nil {
			break
		}
	}
	return resp, err
}

func main() {
	resp, err := SimpleRetry(nil)
	if err != nil {
		fmt.Println("simple retry failed")
	}
	fmt.Printf("simple retry: %v", resp)
}

// 3 Problems
// 1) no time duration between retry
// 2) when error happens, more retry is not needed
// 3) Thundering herd problem, see alse https://en.wikipedia.org/wiki/Thundering_herd_problem