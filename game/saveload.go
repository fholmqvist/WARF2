package game

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"os"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/entity"
	j "projects/games/warf2/jobsystem"
	"projects/games/warf2/mouse"
	u "projects/games/warf2/ui"
	m "projects/games/warf2/worldmap"
)

// SaveGame defines the equivalent
// struct of game that is safe for
// marshaling to JSON.
type SaveGame struct {
	WorldMap  m.Map         `json:"w"`
	Dwarves   []dwarf.Dwarf `json:"dw"`
	JobSystem j.JobSystem   `json:"j"`
	Data      entity.Data   `json:"da"`
}

func (g Game) saveGame() {
	sg := SaveGame{
		WorldMap:  g.WorldMap,
		Dwarves:   g.Dwarves,
		JobSystem: g.JobSystem,
		Data:      g.Data,
	}

	sg.saveToDisk()
}

// Saves the current
// game to disk.
func (sg SaveGame) saveToDisk() {
	filename := "testing.json"

	file, err := os.Create(fmt.Sprintf("./saves/%s", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m, err := json.Marshal(sg)
	if err != nil {
		log.Fatal(m, err)
	}

	_, err = io.WriteString(file, string(m))
	if err != nil {
		log.Fatal(err)
	}
}

// Loads a game
// from disk.
func loadGame() Game {
	filename := "./saves/testing.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Unable to load file:", filename, err)
	}

	sg := SaveGame{}
	err = json.Unmarshal(file, &sg)
	if err != nil {
		log.Fatal(err)
	}

	sg.JobSystem.Map = &sg.WorldMap

	var workers []j.Worker
	for i := range sg.Dwarves {
		workers = append(workers, &sg.Dwarves[i])
	}
	sg.JobSystem.Workers = workers

	for i := range sg.WorldMap.Tiles {
		sg.WorldMap.Tiles[i].Map = &sg.WorldMap
		sg.WorldMap.SelectedTiles[i].Map = &sg.WorldMap
		sg.WorldMap.Items[i].Map = &sg.WorldMap
	}

	return Game{
		WorldMap:  sg.WorldMap,
		Dwarves:   sg.Dwarves,
		JobSystem: sg.JobSystem,
		Data:      sg.Data,

		time:        Time{Frame: 1},
		mouseSystem: mouse.System{},
		ui: u.UI{
			MouseMode: u.Element{
				X:     m.TileSize,
				Y:     m.TileSize*m.TilesH - m.TileSize,
				Color: color.White,
			},
		},
	}
}
