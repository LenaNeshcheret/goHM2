package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Location string

const (
	Start                     Location = "start"
	Forest                    Location = "forest"
	ForestTrail               Location = "forest_trail"
	ForestTrailFollow         Location = "forest_trail_follow"
	ForestTrailAsk            Location = "forest_trail_ask"
	ForestTrailAlone          Location = "forest_trail_alone"
	ForestTree                Location = "forest_tree"
	ForestTreeLight           Location = "forest_tree_light"
	ForestTreeObserve         Location = "forest_tree_observe"
	ForestDeep                Location = "forest_deep"
	River                     Location = "river"
	RiverBridge               Location = "river_bridge"
	RiverBridgeStay           Location = "river_bridge_stay"
	RiverBridgeStayDownstream Location = "river_bridge_stay_downstream"
	RiverBridgeStayUpstream   Location = "river_bridge_stay_upstream"
	RiverPath                 Location = "river_path"
	RiverPathHut              Location = "river_path_hut"
	RiverPathRest             Location = "river_path_rest"
	Cave                      Location = "cave"
	CaveLeft                  Location = "cave_left"
	CaveLeftPool              Location = "cave_left_pool"
	CaveLeftPoolObject        Location = "cave_left_pool_object"
	CaveLeftPoolExits         Location = "cave_left_pool_exits"
	CaveLeftPoolTunnel        Location = "cave_left_pool_tunnel"
	CaveRight                 Location = "cave_right"
	CaveRightCrystals         Location = "cave_right_crystals"
	CaveRightDeepCave         Location = "cave_right_deep_cave"
)

var endGameLocations = []Location{ForestTrailFollow, ForestTrailAsk, ForestTrailAlone, ForestTreeLight,
	ForestDeep, RiverBridgeStayUpstream, RiverBridgeStayDownstream, RiverPathHut, RiverPathRest,
	CaveLeftPoolObject, CaveLeftPoolTunnel, CaveRightCrystals, CaveRightDeepCave}

type Situation struct {
	Description string
	Actions     []string
}

var startOverAction = []string{"If you want to START again, press any button"}

var gameOver = "\nGame over."

var win = "\nCongratulations! You win."

var situations = map[Location]Situation{
	Start: {
		Description: "You wake up in an unknown place with some basic items.\nYou can see three paths ahead of you:",
		Actions:     []string{"1. A forest", "2. A river", "3. A cave"},
	},
	Forest: {
		Description: "You enter the dark forest. It's eerie and quiet.",
		Actions: []string{"1. Follow a faint trail.", "2. Climb a tree to get a better view.",
			"3. Go deeper into the forest.", "4. Go back to the starting point."},
	},
	ForestTrail: {
		Description: "You follow the faint trail and encounter a mysterious figure.",
		Actions: []string{"1. Trust the figure and follow them.", "2. Avoid the figure and continue alone.",
			"3. Ask the figure for help but don't follow.", "4. Go back to the forest entrance."},
	},
	ForestTrailFollow: {
		Description: "You decide to trust the figure and follow them. They lead you to a hidden village where you " +
			"find help and safety. You are saved." + win,
		Actions: startOverAction,
	},
	ForestTrailAsk: {
		Description: "You ask the figure and she ate you." + gameOver,
		Actions:     startOverAction,
	},
	ForestTrailAlone: {
		Description: "You decide to avoid the figure and continue alone. You got lost and starved to death." + gameOver,
		Actions:     startOverAction,
	},
	ForestTree: {
		Description: "You climb a tree and get a better view. You see a light in the distance.",
		Actions: []string{"1. Head towards the light.", "2. Stay in the tree and observe.",
			"3. Climb down and go deeper into the forest.", "4. Climb down and go back to the forest entrance."},
	},
	ForestTreeLight: {
		Description: "You head towards the light and discover an abandoned cabin. Inside, you find supplies and a map " +
			"that leads you out of the forest. You are saved." + win,
		Actions: startOverAction,
	},
	ForestTreeObserve: {
		Description: "You stay in the tree and observe. As night falls, you see glowing eyes approaching. " +
			"You stay quiet and hidden until dawn. You survive the night but remain in the forest.",
		Actions: []string{"1. Climb down and head towards the light.", "2. Climb down and go deeper into the forest.",
			"3. Climb down and go back to the forest entrance."},
	},
	ForestDeep: {
		Description: "You go deeper into the forest and encounter a pack of wild animals." +
			" You are unable to escape and meet a tragic end." + gameOver,
		Actions: startOverAction,
	},
	River: {
		Description: "You arrive at the bank of a rushing river. The water is cold and fast-moving. " +
			"You can see a bridge upstream and a path alongside the riverbank.",
		Actions: []string{"1. Cross the bridge", "2. Follow the riverbank path", "3. Go back to the starting point"},
	},
	RiverBridge: {
		Description: "You decide to cross the bridge. As you step onto the bridge, you notice it sways slightly under" +
			" your weight. Halfway across, you see the river roaring beneath you. You reach the other side safely and" +
			" find yourself at the edge of a dense forest. A narrow path leads deeper into the woods.",
		Actions: []string{"1. Enter the forest", "2. Stay by the river", "3. Go back across the bridge"},
	},
	RiverBridgeStay: {
		Description: "You decide to stay by the river. You sit down by the riverbank and enjoy the sound of the " +
			"rushing water. After a while, you notice a small boat tied to a tree nearby.",
		Actions: []string{"1. Take the boat and row downstream", "2. Take the boat and row upstream",
			"3. Ignore the boat and go back across the bridge"},
	},
	RiverBridgeStayDownstream: {
		Description: "The current washes you into the waterfall and you die." + gameOver,
		Actions:     startOverAction,
	},
	RiverBridgeStayUpstream: {
		Description: "After rowing for several hours, you see friends who help you get out." + win,
		Actions:     startOverAction,
	},
	RiverPath: {
		Description: "You decide to follow the riverbank path. The path is narrow and lined with thick vegetation. " +
			"You hear the sound of the river rushing beside you, and birds chirping in the trees above." +
			" After walking for a while, you come across a small, secluded clearing with a beautiful view of the river. " +
			"In the clearing, you see a small wooden hut.",
		Actions: []string{"1. Investigate the hut", "2. Continue following the path", "3. Rest in the clearing"},
	},
	RiverPathHut: {
		Description: "You went into the hut and found a phone. You called friends and found you." + win,
		Actions:     startOverAction,
	},
	RiverPathRest: {
		Description: "You fell asleep and were eaten by wild animals." + gameOver,
		Actions:     startOverAction,
	},
	Cave: {
		Description: "You arrive at the entrance of a dark, foreboding cave. The air is cool and damp, and you can " +
			"hear the faint sound of dripping water echoing from within. You can see two passages inside the cave.",
		Actions: []string{"1. Enter the left passage", "2. Enter the right passage",
			"3. Leave the cave and go back to the starting point"},
	},
	CaveLeft: {
		Description: "You decide to enter the left passage. The passage is narrow and you have to squeeze through " +
			"in some places. After a while, you reach a small chamber with a pool of water in the center.",
		Actions: []string{"1. Investigate the pool", "2. Search the chamber for other exits",
			"3. Go back to the cave entrance"},
	},
	CaveLeftPool: {
		Description: "You decide to investigate the pool. As you approach the pool, you notice something shiny at the bottom.",
		Actions:     []string{"1. Reach into the pool to grab the shiny object", "2. Ignore the object and leave the chamber"},
	},
	CaveLeftPoolObject: {
		Description: "You found explosives and unfortunately you died." + gameOver,
		Actions:     startOverAction,
	},
	CaveLeftPoolExits: {
		Description: "You decide to search the chamber for other exits. You find a narrow tunnel hidden behind some rocks.",
		Actions:     []string{"1. Enter the tunnel", "2. Go back to the cave entrance"},
	},
	CaveLeftPoolTunnel: {
		Description: "You entered the tunnel and were bitten by a snake. You died." + gameOver,
		Actions:     startOverAction,
	},
	CaveRight: {
		Description: "You decide to enter the right passage. The passage is wider and slopes downward, leading deeper " +
			"into the cave. After a while, you reach a large cavern illuminated by glowing crystals.",
		Actions: []string{"1. Examine the crystals", "2. Explore further into the cavern", "3. Go back to the cave entrance"},
	},
	CaveRightCrystals: {
		Description: "You decide to examine the crystals. The crystals are beautiful and emit a soft, warm light. " +
			"As you touch one, you feel a strange energy coursing through you. And you find yourself at home. " +
			"It was a magic crystal." + win,
		Actions: startOverAction,
	},
	CaveRightDeepCave: {
		Description: "You keep going and can't find a way out." + gameOver,
		Actions:     startOverAction,
	},
}

func main() {
	location := Start
	for {
		printSituation(location)
		choice := getChoice()
		location = handleChoice(location, choice)
		for _, winOrGameOverLocation := range endGameLocations {
			if location == winOrGameOverLocation {
				printSituation(location)
				choice := getChoice()
				location = handleChoice(location, choice)
				location = Start
			}
		}
	}
}

func printSituation(location Location) {
	situation, ok := situations[location]
	if !ok {
		fmt.Println("Invalid location.")
		return
	}
	fmt.Println(situation.Description)
	for _, action := range situation.Actions {
		fmt.Println(action)
	}
}

func getChoice() string {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

var locationMap = map[Location]map[string]Location{
	Start: {
		"1": Forest,
		"2": River,
		"3": Cave,
	},
	Forest: {
		"1": ForestTrail,
		"2": ForestTree,
		"3": ForestDeep,
		"4": Start,
	},
	ForestTrail: {
		"1": ForestTrailFollow,
		"2": ForestTrailAlone,
		"3": ForestTrailAsk,
		"4": Forest,
	},
	ForestTree: {
		"1": ForestTreeLight,
		"2": ForestTreeObserve,
		"3": ForestDeep,
		"4": Forest,
	},
	ForestTreeObserve: {
		"1": ForestTreeLight,
		"2": ForestDeep,
		"3": Forest,
	},
	River: {
		"1": RiverBridge,
		"2": RiverPath,
		"3": Start,
	},
	RiverBridge: {
		"1": Forest,
		"2": RiverBridgeStay,
		"3": River,
	},
	RiverBridgeStay: {
		"1": RiverBridgeStayDownstream,
		"2": RiverBridgeStayUpstream,
		"3": RiverBridge,
	},
	RiverPath: {
		"1": RiverPathHut,
		"2": Forest,
		"3": RiverPathRest,
	},
	Cave: {
		"1": CaveLeft,
		"2": CaveRight,
		"3": Start,
	},
	CaveLeft: {
		"1": CaveLeftPool,
		"2": CaveLeftPoolExits,
		"3": Cave,
	},
	CaveLeftPool: {
		"1": CaveLeftPoolObject,
		"2": CaveLeft,
	},
	CaveLeftPoolExits: {
		"1": CaveLeftPoolTunnel,
		"2": CaveLeft,
	},
	CaveRight: {
		"1": CaveRightCrystals,
		"2": CaveRightDeepCave,
		"3": Cave,
	}}

func handleChoice(location Location, choice string) Location {
	if nextLocation, ok := locationMap[location][choice]; ok {
		return nextLocation
	}
	return location
}
