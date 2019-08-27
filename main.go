package main

import(
	"os"
	"fmt"
	"strconv"
	"./aliens"
	"time"
)


func main() {

	aliens.AliensInCity = make(map[string][]int) //Initialize mapping
	aliens.PresentAlienCity = make(map[int]string) //Initialize mapping

	s :=os.Args[1]
	alien, _ := strconv.Atoi(s) // Command line input for the number of aliens
	aliens.AlienMap()	//Call to the function to read world map from the text file and create corresponding mappings
	
	 // For each alien, Call for the function for alien first landing into a city
	for i:=1 ; i <= alien; i++{
		aliens.PresentAlienCity[i] = ""
		aliens.AlienRoam(i)
	}
	
	// For aliens surviving, call for a function to subsequent routing of aliens from city to city till all the cities have been destroyed or all aliens are dead
	
	for len(aliens.PresentAlienCity) != 0 && aliens.Count <10000{
		for k := range aliens.PresentAlienCity {
			aliens.Count++
			aliens.AlienRoam(k)
		}
	}

	// Check if the alien have travelled more than 10000 iterations
	//If yes, Exit the progarm and print out the remaining map of the world
	if aliens.Count>=10000{
		fmt.Println("\n", len(aliens.PresentAlienCity), " alien/alien's still exist and have moved more than 10000 times")
		//fmt.Println("\nAliens Moved More than 10000 times")
		aliens.RemainingMap()
		os.Exit(0)
	}

	time.Sleep(time.Millisecond*30000)
}	

