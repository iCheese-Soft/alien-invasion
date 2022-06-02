package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

	
func check(e error) {
	if e != nil {
		panic(e)
	}
}


func RandCityName(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	city := make([]rune, n)
  for i := range city {
    city[i] = letterRunes[rand.Intn(len(letterRunes))]
  }
  return string(city)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cityNames := make([]string, 0)
	cityCount := 10000
	directions := make([]string, 4)
	directions[0] = "north"
	directions[1] = "east"
	directions[2] = "south"
	directions[3] = "west"

	filename := "earth10000.txt"
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	for i := 0; i < cityCount; i ++ {
		cityNames = append(cityNames, RandCityName(5));
	}

	for i := 0; i < cityCount; i ++ {
		s := fmt.Sprint(cityNames[i]);
		for j := 0; j < 4; j ++ {
			if rand.Intn(2) % 5 == 0 {
				continue;
			}
			ss := fmt.Sprint(" " + directions[j] + "=" + cityNames[rand.Intn(cityCount)]);
			s += ss;
		}
		s += "\n";
		f.WriteString(s);
	}
}