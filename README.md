# Alien_Invasion
Simulating Alien Invasion on the world, written in Golang.

# To Run the Progarm:
Run the command 

go run main.go "n"

where n is the number of aliens you want to simulate the invasion of.

Example:

go run main.go 8

# Problem Statement
You should create a program that reads in the world map, creates N aliens, and unleashes them. The program should run until all the aliens have been destroyed, or each alien has moved at least 10,000 times.

And Once the program has finished, it should print out whatever is left of the world in the same format as the input file.

# Solution Approach:
There are many ways/approaches to the above problem statement of Alien Invasion, like using go-routines+channels, go-routines+timeSleep etc.. In my approach to the solution, I wanted to explore more of the basic golang features, and have mainly used the features of golang maps as a way to simulate the above problem statement.

# Files contained:
1. a.txt - A text file containing world map, which will be read as input for the simulation of alien routing/travel to different places.

2. main.go - main go execution file, which simulates the alien Invasion.

3. aliens/aliens.go - Go file cotaining logic for the alien invasion simulation

# Solution Outcomes:
According to problem statement when two aliens end up in the same place, they fight, and in the process kill each other and destroy the city. When a city is destroyed, it is removed from the map, and so are any roads that lead into or out of it.

Some of the outcomes expected out of the solution:

### Scenario 1:
When two aliens end up in the same city/palce they fight and destroy each other,then will print out a message like:

#### "City Bar has been destroyed by alien 3 and alien 7"

### Scenario 2:
If all the cities are destroyed and there still exists some aliens, then will print out a message like:

#### "All the cities have been destroyed by aliens and still some aliens survive"

### Scenario 3: 
If all the cities and aliens are destroyed, then will print out a message like:

#### "All cities and aliens have been destroyed"

### Scenario 4:
When all the aliens are destroyed but still some cities reamain safe, then will print out a message like:

#### "All the aliens sucessfully destroyed"

Additionally will print the remaining world map:

Example: If "Foo north=Bar west=Baz south=Qu-ux" is the initial map of the city Foo, and if cities Bar and Qu-ux have been destroyed along with aliens, then the map for the city Foo would look something like:

#### Foo north= west=Baz south=

### Scenario 5:
When aliens dont end up in the same city, i.e when they cannot fight and destroy each other, continue routing the aliens to new city for >10000 times, and if alien has moved at more than 10,000 times, then print out a message like:

#### "Aliens Moved More than 10000 times"

Additioanlly print the reamining world map:

#### Note: More than one scenario can occur for a single simulation run.






