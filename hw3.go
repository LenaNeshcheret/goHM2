package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	Location := "start"

	for {
		printSituation(Location)
		choice := getChoice()
		Location = handleChoice(Location, choice)
	}

}

func printSituation(location string) {
	switch location {
	case "start":
		fmt.Println("You wake up in an unknown place with some basic items.")
		fmt.Println("You can see three paths ahead of you:")
		fmt.Println("1. A forest")
		fmt.Println("2. A river")
		fmt.Println("3. A cave")
		fmt.Println("Which path will you take? (1/2/3)")
	case "forest":
		fmt.Println("You enter the dark forest. It's eerie and quiet.")
		fmt.Println("1. Follow a faint trail.")
		fmt.Println("2. Climb a tree to get a better view.")
		fmt.Println("3. Go deeper into the forest.")
		fmt.Println("4. Go back to the starting point.")
		fmt.Println("Which action will you take? (1/2/3/4)")
	case "forest_trail":
		fmt.Println("You follow the faint trail and encounter a mysterious figure.")
		fmt.Println("1. Trust the figure and follow them.")
		fmt.Println("2. Avoid the figure and continue alone.")
		fmt.Println("3. Ask the figure for help but don't follow.")
		fmt.Println("4. Go back to the forest entrance.")
		fmt.Println("Which action will you take? (1/2/3/4)")
	case "forest_trail_follow":
		fmt.Println("You decide to trust the figure and follow them. " +
			"They lead you to a hidden village where you find help and safety. You are saved.")
		printSituation("win")
	case "forest_trail_ask":
		fmt.Println("You ask the figure and she ate you.")
		printSituation("game_over")
	case "forest_trail_alone":
		fmt.Println("You decide to avoid the figure and continue alone." +
			"You got lost and starved to death.")
		printSituation("game_over")
	case "forest_tree":
		fmt.Println("You climb a tree and get a better view. You see a light in the distance.")
		fmt.Println("1. Head towards the light.")
		fmt.Println("2. Stay in the tree and observe.")
		fmt.Println("3. Climb down and go deeper into the forest.")
		fmt.Println("4. Climb down and go back to the forest entrance.")
		fmt.Println("Which action will you take? (1/2/3/4)")
	case "forest_tree_light":
		fmt.Println("You head towards the light and discover an abandoned cabin. Inside, you find supplies and a map that leads you out of the forest. You are saved.")
		printSituation("win")
	case "forest_tree_observe":
		fmt.Println("You stay in the tree and observe. As night falls, you see glowing eyes approaching. You stay quiet and hidden until dawn. You survive the night but remain in the forest.")
		fmt.Println("1. Climb down and head towards the light.")
		fmt.Println("2. Climb down and go deeper into the forest.")
		fmt.Println("3. Climb down and go back to the forest entrance.")
		fmt.Println("Which action will you take? (1/2/3)")
	case "forest_tree_observe_light":
		fmt.Println("You followed the light and found a way out.")
		printSituation("win")
	case "forest_deep":
		fmt.Println("You go deeper into the forest and encounter a pack of wild animals. " +
			"You are unable to escape and meet a tragic end.")
		printSituation("game_over")
	case "river":
		fmt.Println("You arrive at the bank of a rushing river.")
		fmt.Println("The water is cold and fast-moving.")
		fmt.Println("You can see a bridge upstream and a path alongside the riverbank.")
		fmt.Println("1. Cross the bridge")
		fmt.Println("2. Follow the riverbank path")
		fmt.Println("3. Go back to the starting point")
		fmt.Println("What will you do? (1/2/3)")
	case "river_bridge":
		fmt.Println("You decide to cross the bridge.")
		fmt.Println("As you step onto the bridge, you notice it sways slightly under your weight.")
		fmt.Println("Halfway across, you see the river roaring beneath you.")
		fmt.Println("You reach the other side safely and find yourself at the edge of a dense forest.")
		fmt.Println("A narrow path leads deeper into the woods.")
		fmt.Println("Do you want to:")
		fmt.Println("1. Enter the forest")
		fmt.Println("2. Stay by the river")
		fmt.Println("3. Go back across the bridge")
		fmt.Println("What will you do? (1/2/3)")
	case "river_bridge_stay":
		fmt.Println("You decide to stay by the river.")
		fmt.Println("You sit down by the riverbank and enjoy the sound of the rushing water.")
		fmt.Println("After a while, you notice a small boat tied to a tree nearby.")
		fmt.Println("Do you want to:")
		fmt.Println("1. Take the boat and row downstream")
		fmt.Println("2. Take the boat and row upstream")
		fmt.Println("3. Ignore the boat and  go back across the bridge")
		fmt.Println("What will you do? (1/2/3)")
	case "river_bridge_stay_downstream":
		fmt.Println("The current washes you into the waterfall and you die")
		printSituation("game_over")
	case "river_bridge_stay_upstream":
		fmt.Println("After swimming for several hours, you see friends who help you get out.")
		printSituation("win")
	case "river_path":
		fmt.Println("You decide to follow the riverbank path.")
		fmt.Println("The path is narrow and lined with thick vegetation.")
		fmt.Println("You hear the sound of the river rushing beside you, and birds chirping in the trees above.")
		fmt.Println("After walking for a while, you come across a small, secluded clearing with a beautiful view of the river.")
		fmt.Println("In the clearing, you see a small wooden hut.")
		fmt.Println("Do you want to:")
		fmt.Println("1. Investigate the hut")
		fmt.Println("2. Continue following the path")
		fmt.Println("3. Rest in the clearing")
		fmt.Println("What will you do? (1/2/3)")
	case "river_path_hut":
		fmt.Println("They went into the hut and found a phone. They called friends and found you.")
		printSituation("win")
	case "river_path_rest":
		fmt.Println("You fell asleep and were eaten by wild animals")
		printSituation("win")
	case "cave":
		fmt.Println("You arrive at the entrance of a dark, foreboding cave.")
		fmt.Println("The air is cool and damp, and you can hear the faint sound of dripping water echoing from within.")
		fmt.Println("You can see two passages inside the cave.")
		fmt.Println("1. Enter the left passage")
		fmt.Println("2. Enter the right passage")
		fmt.Println("3. Leave the cave and go back to the starting point")
		fmt.Println("What will you do? (1/2/3)")
	case "cave_left":
		fmt.Println("You decide to enter the left passage.")
		fmt.Println("The passage is narrow and you have to squeeze through in some places.")
		fmt.Println("After a while, you reach a small chamber with a pool of water in the center.")
		fmt.Println("Do you want to:")
		fmt.Println("1. Investigate the pool")
		fmt.Println("2. Search the chamber for other exits")
		fmt.Println("3. Go back to the cave entrance")
		fmt.Println("What will you do? (1/2/3)")
	case "cave_left_pool":
		fmt.Println("You decide to investigate the pool.")
		fmt.Println("As you approach the pool, you notice something shiny at the bottom.")
		fmt.Println("Do you want to:")
		fmt.Println("1. Reach into the pool to grab the shiny object")
		fmt.Println("2. Ignore the object and leave the chamber")
		fmt.Println("What will you do? (1/2)")
	case "cave_left_pool_object":
		fmt.Println("You found explosives and unfortunately you died.1")
		printSituation("game_over")
	case "cave_left_pool_otherExits":
		fmt.Println("You decide to search the chamber for other exits.")
		fmt.Println("You find a narrow tunnel hidden behind some rocks.")
		fmt.Println("Do you want to:")
		fmt.Println("1. Enter the tunnel")
		fmt.Println("2. Go back to the cave entrance")
		fmt.Println("What will you do? (1/2)")
	case "cave_left_pool_otherExits_tunnel":
		fmt.Println("You entered the tunnel and were bitten by a snake. You died.")
		printSituation("game_over")
	case "cave_right":
		fmt.Println("You decide to enter the right passage.")
		fmt.Println("The passage is wider and slopes downward, leading deeper into the cave.")
		fmt.Println("After a while, you reach a large cavern illuminated by glowing crystals.")
		fmt.Println("Do you want to:")
		fmt.Println("1. Examine the crystals")
		fmt.Println("2. Explore further into the cavern")
		fmt.Println("3. Go back to the cave entrance")
		fmt.Println("What will you do? (1/2/3)")
	case "cave_right_crystals":
		fmt.Println("You decide to examine the crystals.")
		fmt.Println("The crystals are beautiful and emit a soft, warm light.")
		fmt.Println("As you touch one, you feel a strange energy coursing through you.")
		fmt.Println("And you find yourself at home. It was a magic crystal")
		printSituation("win")
	case "cave_right_deepCave":
		fmt.Println("You keep going and can't find a way out.")
		printSituation("game_over")
	case "win":
		fmt.Println("Congratulations! You are win.")
		fmt.Println("Please press 1 and start the game over")
	case "game_over":
		fmt.Println("Game over.")
		fmt.Println("Please press 1 and start the game over")
	default:
		fmt.Println("Invalid choice. Please enter a valid option.")
	}
}

func getChoice() string {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func handleChoice(location, choice string) string {
	switch location {
	case "start":
		switch choice {
		case "1":
			location = "forest"
		case "2":
			location = "river"
		case "3":
			location = "cave"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "forest":
		switch choice {
		case "1":
			location = "forest_trail"
		case "2":
			location = "forest_tree"
		case "3":
			location = "forest_deep"
		case "4":
			location = "start"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "forest_trail":
		switch choice {
		case "1":
			location = "forest_trail_follow"
		case "2":
			location = "forest_trail_alone"
		case "3":
			location = "forest_trail_ask"
		case "4":
			location = "forest"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "forest_tree":
		switch choice {
		case "1":
			location = "forest_tree_light"
		case "2":
			location = "forest_tree_observe"
		case "3":
			location = "forest_deep"
		case "4":
			location = "forest"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "forest_tree_observe":
		switch choice {
		case "1":
			location = "forest_tree_observe_light"
		case "2":
			location = "forest_deep"
		case "3":
			location = "forest"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "river":
		switch choice {
		case "1":
			location = "river_bridge"
		case "2":
			location = "river_path"
		case "3":
			location = "start"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "river_bridge":
		switch choice {
		case "1":
			location = "forest"
		case "2":
			location = "river_bridge_stay"
		case "3":
			location = "river"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "river_bridge_stay":
		switch choice {
		case "1":
			location = "river_bridge_stay_downstream"
		case "2":
			location = "river_bridge_stay_upstream"
		case "3":
			location = "river_bridge"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "river_path":
		switch choice {
		case "1":
			location = "river_path_hut"
		case "2":
			location = "forest"
		case "3":
			location = "river_path_rest"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "cave":
		switch choice {
		case "1":
			location = "cave_left"
		case "2":
			location = "cave_right"
		case "3":
			location = "start"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "cave_left":
		switch choice {
		case "1":
			location = "cave_left_pool"
		case "2":
			location = "cave_left_pool_otherExits"
		case "3":
			location = "cave"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "cave_left_pool_otherExits":
		switch choice {
		case "1":
			location = "cave_left_pool_otherExits_tunnel"
		case "2":
			location = "cave_left"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "cave_left_pool":
		switch choice {
		case "1":
			location = "cave_left_pool_object"
		case "2":
			location = "cave_left"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "cave_right":
		switch choice {
		case "1":
			location = "cave_right_crystals"
		case "2":
			location = "cave_right_deepCave"
		case "3":
			location = "cave"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	case "forest_trail_follow", "forest_trail_alone", "forest_trail_ask", "forest_tree_light", "forest_deep",
		"forest_tree_observe_light", "river_bridge_stay_upstream", "river_bridge_stay_downstream",
		"river_path_hut", "river_path_rest", "cave_left_pool_object", "cave_left_pool_otherExits_tunnel",
		"cave_right_crystals", "cave_right_deepCave":
		switch choice {
		case "1":
			location = "start"
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
	return location
}
