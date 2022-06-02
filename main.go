package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"bufio"
	"strings"
	"math/rand"
)

func Readln(r *bufio.Reader) (string, error) {
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )
  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Run with the correct args!\nFor example: go run main.go ./earth.txt 100")
		os.Exit(0)
	}

	filepath := os.Args[1]
	N, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Run with the correct args!\nFor example: go run main.go ./earth.txt 100")
		os.Exit(0)
	}

	fmt.Println(filepath)
	fmt.Println(N)

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	s, err := Readln(r)

	// Read from the file and prepare for the processing.
	cityCount := 0
	cityToId := make(map[string]int)
	isDestroyed := make([]bool, 0)
	neighbourCities := make([][]string, 0)
	for err == nil {
		words := strings.Split(s, " ")
		cityToId[words[0]] = cityCount
		isDestroyed = append(isDestroyed, false)
		neighbour := make([]string, 0)
		for j := range words[1:] {
			tkns := strings.Split(words[1 + j], "=")
			neighbour = append(neighbour, tkns[1])
		}
		neighbourCities = append(neighbourCities, neighbour)
		cityCount ++
		s, err = Readln(r)
	}


	// Assign each aliens to the city randomly.
	aliens := make([]int, N)
	isAlienCoveredCity := make([]int, cityCount)
	cityDestroyed := make([]bool, cityCount)
	for i := 0; i < N; i ++ {
		j := rand.Intn(cityCount)
		for isAlienCoveredCity[j] == 1 {
			j = rand.Intn(cityCount)
		}
		
		isAlienCoveredCity[j] = 1
		aliens[i] = j
	}


	// Start moving of aliens.
	stepCount := 0

	for stepCount < 10000 {
		
		isAlienCoveredCity = make([]int, cityCount)

		for i := 0; i < N; i ++ {
			
			if aliens[i] == -1 {
				continue;
			}

			trapped := true;

			for j := 0; j < len(neighbourCities[i]); j ++ {
				if !cityDestroyed[cityToId[neighbourCities[i][j]]] {
					trapped = false;
					break;
				}
			}

			if trapped {
				// trapped
				continue;
			}

			for j := 0; j < len(neighbourCities[aliens[i]]); j ++ {
				nextCityId := cityToId[neighbourCities[aliens[i]][j]]
				if !cityDestroyed[nextCityId] {
					// Alien i moves to cityToId[neighbourCities[i][j]]

					if isAlienCoveredCity[nextCityId] != 0 {
						// Alien i trys to move but already another alien arrived at the city.
						// The city will be destroyed.
						// Two Aliens will be killed.

						fmt.Printf("%s has been destroyed by alien %d and alien %d!\n", neighbourCities[aliens[i]][j], i + 1, isAlienCoveredCity[nextCityId])

						cityDestroyed[nextCityId] = true
						aliens[i] = -1
						aliens[isAlienCoveredCity[nextCityId] - 1] = -1
						break;
					}

					isAlienCoveredCity[nextCityId] = i + 1
					aliens[i] = nextCityId

					break;
				}
			}

		}
		stepCount ++
	}


}