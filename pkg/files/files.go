// Copyright © 2019 Osiloke Harold Emoekpere <me@osiloke.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package files

import (
	"crypto/sha256"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func HashFilename(fileName string) string {
	extension := filepath.Ext(fileName)
	name := fileName[0 : len(fileName)-len(extension)]
	v := sha256.Sum256([]byte(name))
	// p := sha256.Sum256([]byte(strconv.Itoa(int(time.Now().Unix()))))
	// return fmt.Sprintf("%x_%x.%s", v, p, extension)
	return fmt.Sprintf("%x%s", v, extension)
}
func HashFilenameWithTime(fileName string) string {
	extension := filepath.Ext(fileName)
	name := fileName[0 : len(fileName)-len(extension)]
	v := sha256.Sum256([]byte(name + strconv.Itoa(int(time.Now().Unix()))))
	// p := sha256.Sum256([]byte(strconv.Itoa(int(time.Now().Unix()))))
	// return fmt.Sprintf("%x_%x.%s", v, p, extension)
	return fmt.Sprintf("%x%s", v, extension)
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CopyFile copy file
func CopyFile(filename, destination string) error {
	if ok, _ := Exists(destination); ok {
		return errors.New("CopyFile failed - " + destination + " already exists")
	}
	if ok, _ := Exists(filename); !ok {
		return errors.New("CopyFile failed - " + filename + " does not exist")
	}
	from, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to open filename - %v", err)
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("unable to open destination - %v", err)
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	return nil
}

// WriteFile write data to file
func WriteFile(data []byte, destination string) error {
	if ok, _ := Exists(destination); ok {
		return errors.New("WriteFile failed - " + destination + " already exists")
	}
	to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("unable to open destination - %v", err)
		return err
	}
	defer to.Close()

	_, err = to.Write(data)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	return nil
}

// ReadCsv read a csv file, assumes first line is the header
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines[1:], nil
}
