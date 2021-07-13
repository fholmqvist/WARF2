package game

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	j "github.com/Holmqvist1990/WARF2/jobservice"
	"github.com/Holmqvist1990/WARF2/mouse"
	u "github.com/Holmqvist1990/WARF2/ui"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

// SaveGame defines the equivalent
// struct of game that is safe for
// marshaling to JSON.
type SaveGame struct {
	WorldMap   m.Map        `json:"w"`
	JobService j.JobService `json:"j"`
}

func (g Game) SaveGame() {
	sg := SaveGame{
		WorldMap:   *g.WorldMap,
		JobService: *g.JobService,
	}
	sg.saveToDisk()
}

// Saves the current
// game to disk.
func (sg SaveGame) saveToDisk() {
	_, err := os.Stat("./saves/")
	if os.IsNotExist(err) {
		err := os.MkdirAll("saves", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
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
	sg.JobService.Map = &sg.WorldMap
	for i := range sg.WorldMap.Tiles {
		sg.WorldMap.Tiles[i].Map = &sg.WorldMap
		sg.WorldMap.SelectedTiles[i].Map = &sg.WorldMap
		sg.WorldMap.Items[i].Map = &sg.WorldMap
	}
	return Game{
		WorldMap:    &sg.WorldMap,
		JobService:  &sg.JobService,
		time:        Time{Frame: 1},
		mouseSystem: &mouse.System{},
		ui:          u.GenerateUI(),
	}
}
