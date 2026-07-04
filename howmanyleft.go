package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: howmanyleft (%start <> %goal) %inc")
		os.Exit(0)
	}
	var numbers [3]int
	var err error
	for i := range numbers {
		numbers[i] ,err = strconv.Atoi(os.Args[i+1])
		if err != nil {
			fmt.Println("err reading arg",i,":",err)
			os.Exit(-1);
		}
	}
	start,end,inc := numbers[0],numbers[1],numbers[2]
	if end == start {
		fmt.Println(end, "to", start, " == 0")
	}
	if end < start {
		end, start = start, end
	}
	var j int
	for i := start; i < end; i += inc {
		j++
	}
	fmt.Println(start, "to", end, "by", inc, "==", j)

}
