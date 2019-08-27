package aliens

import(
	"os"
	"fmt"
	"math/rand"
	"bufio"
	"strings"	
)

var CityMap map[string][]string // Map of City as key to slice of possible route to other cities from it
var Cities map[string][]string // Map of all existing Cities 
var AliensInCity map[string][]int // Map of cities to aliens present at that City
var PresentAlienCity map[int] string //Map of aliens to their present city they moved to
var DestroyedCity []string // List of all destroyed city
var Count int // Count of Alien Movement
var Map map[string][]string //Map of Original input of City Map to its corresponding routes from it

var Result []string //Array representing map read from the text file


//Function to Read the World map from the text file and map it to corresponding CityMap, Cities, Map
func AlienMap(){

	CityMap = make(map[string][]string)
	Cities = make(map[string][]string)
	Map = make(map[string][]string)

	n := rand.Intn(5)
	text := atLine(n)
	for _, a := range text {
	words := strings.Fields(a)
		for i := 1; i< len(words); i++ {
			Map[words[0]] = append(Map[words[0]], words[i])
			res1 := contains(Cities["Cities"], words[0])
			if res1 != true{
			Cities["Cities"] = append(Cities["Cities"], words[0])
			}			
			words1 := strings.Split(words[i], "=")
			res := contains(CityMap[words[0]], words1[1])
			if res != true{
			CityMap[words[0]] = append(CityMap[words[0]],words1[1])
			}
		}
	}	
}


//Function for alien Movement
func AlienRoam(alien int){
	if Count < 1{ // For the first entry into a city, i.e first landing of aliens
		if (len(Cities["Cities"]) == 0 && len(PresentAlienCity)!=0){ //check if all the cities are destroyed and if some aliens still exist && if yes exit()
			fmt.Println("\nAll the cities have been destroyed by aliens and still some aliens survive")
			os.Exit(1)
		}else{		
			n := alien % len(Cities["Cities"]) // For random selection of Cities for Alien "n" to travel
			//fmt.Print("\nAlien", alien, " has entered city ", Cities["Cities"][n])
			PresentAlienCity[alien] = Cities["Cities"][n] // Map the current City of arrival for Alien "n"
			if len(AliensInCity[Cities["Cities"][n]]) < 2{ //Check's if 2 aliens are in the same city
				AliensInCity[Cities["Cities"][n]] = append(AliensInCity[Cities["Cities"][n]],alien) 
					if len(AliensInCity[Cities["Cities"][n]]) == 2{ //If 2 alien are in the same city, i.e they are routed to the same city, then destroy aliens and the city
						fmt.Println("\nCity", Cities["Cities"][n]," has been destroyed by alien", AliensInCity[Cities["Cities"][n]][0]," and", AliensInCity[Cities["Cities"][n]][1])
						DestroyedCity = append(DestroyedCity, Cities["Cities"][n]) //Update the map of destroyed City
						deleteEntry(Cities["Cities"][n], AliensInCity[Cities["Cities"][n]][0] , AliensInCity[Cities["Cities"][n]][1], n) //Call for a function to delete the city and aliens from their corresponding mapping
						for k := range AliensInCity{
							AliensInCity[k] = nil //reset the map of AliensInCity(which shows which alien is in a particular city), because in next iteration alien would move to different city
						}
				}
			}
		}
		
	}else if Count >=1{ //For subsequent tarvel of aliens from their first city

		k := PresentAlienCity[alien]
		if (len(CityMap) == 0 && len(PresentAlienCity)!=0){ //check if all the cities are destroyed and if some aliens still exist && if yes exit()
			fmt.Println("\nAll the cities have been destroyed by aliens and still some aliens survive")
			os.Exit(2)
			
		}else if len(CityMap) == 0 && len(PresentAlienCity) == 0{ //check if there is any city for the alien to travel && if yes exit()
			fmt.Println("\nAll cities and aliens have been destroyed")
			os.Exit(1)				
		}else{
		if len(CityMap[k])== 0 {
			k = PresentAlienCity[alien]
		}else{
		n:= rand.Intn(9) % len(CityMap[k]) // For random selection of Cities for Alien "n" to travel
		//fmt.Println("\n Alien ", alien," has now entered the city ", CityMap[k][n])
		PresentAlienCity[alien] = CityMap[k][n] // Map the current City of arrival for Alien "n"
		con := contain(AliensInCity[CityMap[k][n]], alien) //check if already an entry of alien "n" for the city travelled exists
		if len(AliensInCity[CityMap[k][n]]) < 2 && con == false{ //Check's if 2 aliens are in the same city
			AliensInCity[CityMap[k][n]] = append(AliensInCity[CityMap[k][n]],alien) //Add to AliensInCity
				if len(AliensInCity[CityMap[k][n]]) == 2{ //If 2 alien are in the same city, i.e they are routed to the same city, then destroy aliens and the city
					fmt.Println("\nCity", CityMap[k][n]," has been destroyed by alien", AliensInCity[CityMap[k][n]][0]," and", AliensInCity[CityMap[k][n]][1])
					DestroyedCity = append(DestroyedCity, CityMap[k][n])  //Update the map of destroyed City
					deleteEntry(CityMap[k][n], AliensInCity[CityMap[k][n]][0] , AliensInCity[CityMap[k][n]][1], n)  //Call for a function to delete the city and aliens from their corresponding mapping
					for k1 := range AliensInCity {
					AliensInCity[k1] = nil //reset the map of AliensInCity(which shows which alien is in a particular city), because in next iteration alien would move to different city
				}	
				//check if all the aliens have been successfully destroyed
				//If yes then print out the reamining map of the world and exit
				if len(PresentAlienCity) == 0{ 
						fmt.Println("All the aliens sucessfully destroyed")
						RemainingMap() //Prints out the remaining map of the world after cites have benn destroyed
						os.Exit(0)
					}
				}
			}
		}	

	}
}

}

//function to display the remaining map of the world
func RemainingMap(){ 
	fmt.Println("\nRemaining map of the city is:")
	for k := range Map{
		fmt.Print("\n",k)
			for k1 := range Map[k]{
				fmt.Print("\t",Map[k][k1])
				}
		}
}

// function to check if a particular element exists in a sting slice
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	
	_, ok := set[item] 
	return ok
}

// function to check if a particular element exists in a integer slice
func contain(slice []int, item int) bool {
	set := make(map[int]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	
	_, ok := set[item] 
	return ok
}

// function to delete the entry of aliens and city from all the mappings once they are destroyed 
func deleteEntry(city string, alien1 int, alien2 int, n int){

	for i := 0; i < len(Cities["Cities"]); i++ { //delete the entry of "city" destroyed from Cities
		url := Cities["Cities"][i]
			if url == city {
				Cities["Cities"] = append(Cities["Cities"][:i], Cities["Cities"][i+1:]...)
				i-- // Important: decrease index
				break
				}
		}
	//fmt.Println(Cities)	
	delete(AliensInCity , city) //Delete the mapping of the city from AliensInCity
	delete(CityMap, city)	//Delete the mapping of the city from CityMap
	delete(Map,city) //Delete the mapping of the city from Map
	for k2 := range Map {
		for i := range Map[k2]{
			s := Map[k2][i]
			s = strings.Replace(s, city, "", 1)
			Map[k2][i] = s
    		// fmt.Println(s)
		}
	}
	for k := range CityMap {
	for i := 0; i < len(CityMap[k]); i++ {
		item := CityMap[k][i]
			if item == city {
				CityMap[k] = append(CityMap[k][:i], CityMap[k][i+1:]...)
				i--
				break
			}
		}	
	}
	delete(PresentAlienCity,alien1) //Delete the mapping/entry of alien1
	delete(PresentAlienCity,alien2) //Delete the mapping/entry of alien2
	
}

//function to read the text file of world map
func atLine(n int) (s []string) { 
   file, _ := os.Open("map.txt")

defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
	line := scanner.Text()
	Result = append( Result, line)
	
}
//fmt.Println("Result",Result)

if err := scanner.Err(); err != nil {
    fmt.Println(err)
}
 return Result
}
	
