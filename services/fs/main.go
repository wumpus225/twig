package fs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

func GetDataDirectory() string {
	dirname, err :=  homedir.Dir()

	if err != nil {
		log.Fatal(err)
	}

	return path.Join(dirname, ".twig-cli")
}

func CreateDataDirectory()  {
	pathname := GetDataDirectory()

	fmt.Println("constructing pot....")

	if _, err := os.Stat(pathname); os.IsNotExist(err) {
		os.Mkdir(path.Join(pathname), os.ModeDir|0755)
	}

	// Create JSON Datafile
	
	dataFileTarget := path.Join(GetDataDirectory(), "user")
	
	if _, err := os.Stat(path.Join(dataFileTarget, "data.json")); os.IsNotExist(err) {
		os.Mkdir(dataFileTarget, os.ModeDir|0755)

		file, _ := json.MarshalIndent(PlantDataFile{ OwnedPlants: []Plant{}}, "", " ")
		_ = ioutil.WriteFile(path.Join(dataFileTarget,"data.json" ), file, 0644)
	} else {
		 fmt.Println("pot is already created!!!!")
		 return
	}

	fmt.Println("succesfully created your brand new pot!")
}


