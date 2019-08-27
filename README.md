# Alien_Invasion
Simulating Alien Invasion on the world, written in Golang.

# To Run the Progarm:
Run the command 

go run main.go "n"

where n is the number of aliens you want to simulate the invasion of.

Example:

go run main.go 8


# Solution Approach:
There are many ways/approaches to the above problem statement of Alien Invasion, like using go-routines+channels, go-routines+timeSleep etc.. In my approach to the solution, I wanted to explore more of the basic golang features, and have mainly used the features of golang maps as a way to simulate the above problem statement.

# Files contained:
1. a.txt - A text file containing world map, which will be read as input for the simulation of alien routing/travel to different places.

2. main.go - main go execution file, which simulates the alien Invasion.

3. aliens/aliens.go - Go file cotaining logic for the alien invasion simulation



