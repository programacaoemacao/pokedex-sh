package filehelper

import (
	"encoding/json"
	"fmt"
	"os"
)

func createJSONFile(filename string) (*os.File, error) {
	return os.Create(filename)
}

func saveJSONFileContent(file *os.File, content interface{}) error {
	jsonData, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("marshaling json error: %v", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("writing error: %v", err)
	}

	return nil
}

func CreateFileWithContent(filename string, content interface{}) error {
	jsonFile, err := createJSONFile(filename)
	if err != nil {
		return fmt.Errorf("can't create file %q: %+v", filename, err)
	}

	defer jsonFile.Close()

	err = saveJSONFileContent(jsonFile, content)
	if err != nil {
		return fmt.Errorf("can't save file content to JSON : %v", err)
	}

	return nil
}
