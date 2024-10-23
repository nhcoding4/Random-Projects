package FileActions

import (
	"os"
	"strconv"
)

// --------------------------------------------------------------------------------------------------------------------

// Loads a high score from a file

func LoadScore(fileLocation string) (int, error) {
	presentFile, err := exists(fileLocation)
	if err != nil {
		return 0, err
	}

	if !presentFile {
		return 0, nil
	}

	data, err := os.ReadFile(fileLocation)
	if err != nil {
		return 0, err
	}

	score, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return score, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Writes the new high score to file.

func WriteScore(fileLocation, directoryLocation string, score int) error {
	directoryExists, err := exists(directoryLocation)
	if err != nil {
		return err
	}

	if !directoryExists {
		err = os.Mkdir(directoryLocation, os.ModePerm)
		if err != nil {
			return err
		}
	}

	highScoreExists, err := exists(fileLocation)
	if err != nil {
		return err
	}

	if highScoreExists {
		err = os.Remove(fileLocation)
	}

	fileData := []byte(strconv.Itoa(score))
	err = os.WriteFile(fileLocation, fileData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// Checks if a file or directory exists

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// --------------------------------------------------------------------------------------------------------------------
