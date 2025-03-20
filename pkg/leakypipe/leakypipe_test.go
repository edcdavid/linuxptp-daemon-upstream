package leakypipe

import (
	"bufio"
	"fmt"
	"testing"
	"time"
)

func TestLeakyPipe(t *testing.T) {
	pipe := New(5, 5)
	defer pipe.Close()
	go func() {
		for i := 0; i < 9; i++ {
			data := []byte(fmt.Sprintf("%d%d%d%d%d", i, i, i, i, i))
			_, _ = pipe.Write(data)
		}
	}()

	allData := ""
	go func() {
		time.Sleep(time.Second)
		for {
			data := make([]byte, 5)
			length, _ := pipe.Read(data)
			if length != 0 {
				allData += string(data)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(allData)
	if len(allData) != 5*5 {
		t.Fail()
	}
}

func TestLeakyPipeScan(t *testing.T) {
	pipe := New(5, 5)
	defer pipe.Close()
	go func() {
		for i := 0; i < 9; i++ {
			data := []byte(fmt.Sprintf("%d%d%d%d%d", i, i, i, i, i))
			_, _ = pipe.Write(data)
		}
	}()

	allData := ""
	go func() {
		time.Sleep(time.Second)
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			output := scanner.Text()

			allData += string(output)

		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(allData)
	if len(allData) != 5*5 {
		t.Fail()
	}

	for i := 0; i < 1; i++ {
		data := []byte(fmt.Sprintf("%d%d%d%d%d", i, i, i, i, i))
		_, _ = pipe.Write(data)
	}
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		output := scanner.Text()

		allData += string(output)

	}

	fmt.Println(allData)
	if len(allData) != 5*6 {
		t.Fail()
	}
}
