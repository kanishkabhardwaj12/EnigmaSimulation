package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type rotor struct {
	Position     int               `json:"position"`
	SlotPosition int               `json:"slot_position"`
	RotorMapIn   map[string]string `json:"rotor_map_in"`
	RotorMapOut  map[string]string `json:"rotor_map_out"`
}

type plugboard struct {
	PlugMap map[string]string `json:"plug_map"`
}

type reflector struct {
	Reflector map[string]string `json:"reflector_map"`
}

type disk struct {
	DiskMapIn  map[string]string `json:"disk_map_in"`
	DiskMapOut map[string]string `json:"disk_map_out"`
}

type enigmaMachine struct {
	Rotor1    rotor     `json:"rotor1"`
	Rotor2    rotor     `json:"rotor2"`
	Rotor3    rotor     `json:"rotor3"`
	Disk      disk      `json:"disk"`
	Plugboard plugboard `json:"plugboard"`
	Reflector reflector `json:"reflector"`
}

func AtoI(current string) int {
	convertedInt, err := strconv.Atoi(current)
	if err != nil {
		fmt.Println("Error converting string value of 'current' from rotor:", err)
	}
	return convertedInt
}

func outOfRangeCheck(check int) int {

	if check > 26 {
		return check % 26
	}
	if check < 0 {
		return (check % 26) + 26
	}
	return check
}

func processRight(r rotor, current string) string {

	convertedInt := AtoI(current)
	convertedInt += r.Position
	convertedInt = outOfRangeCheck(convertedInt)
	current = strconv.Itoa(convertedInt)
	current = r.RotorMapIn[current]
	convertedInt = AtoI(current)
	convertedInt -= r.Position
	convertedInt = outOfRangeCheck(convertedInt)
	current = strconv.Itoa(convertedInt)
	return current
}

func processLeft(r rotor, current string) string {
	convertedInt := AtoI(current)
	convertedInt += r.Position
	convertedInt = outOfRangeCheck(convertedInt)
	current = strconv.Itoa(convertedInt)
	current = r.RotorMapOut[current]
	convertedInt = AtoI(current)
	convertedInt -= r.Position
	convertedInt = outOfRangeCheck(convertedInt)
	current = strconv.Itoa(convertedInt)
	return current
}

func encrypt(input string, e enigmaMachine) string {
	parts := strings.Split(input, "")
	var encrypted strings.Builder

	for _, part := range parts {
		current := part

		// Plugboard in
		// if valuePbIn, foundPbIn := e.Plugboard.PlugMap[current]; foundPbIn {
		// 	current = valuePbIn
		// }
		// Disk
		current = e.Disk.DiskMapIn[current]
		// Rotor1
		current = processRight(e.Rotor1, current)
		// Rotor2
		current = processRight(e.Rotor2, current)
		// // Rotor3
		current = processRight(e.Rotor3, current)
		// // Reflector
		current = e.Reflector.Reflector[current]
		// Rotor3
		current = processLeft(e.Rotor3, current)
		// // Rotor2
		current = processLeft(e.Rotor2, current)
		// Rotor1
		current = processLeft(e.Rotor1, current)
		//diskmap
		current = e.Disk.DiskMapOut[current]
		// Plugboard
		// if valuePbOut, foundPbOut := pb.PlugMap[current]; foundPbOut {
		// 	current = valuePbOut
		// }

		e.Rotor1.Position += 1
		if e.Rotor1.Position > 25 {
			e.Rotor1.Position = 0
			e.Rotor2.Position += 1
			if e.Rotor2.Position > 25 {
				e.Rotor2.Position = 0
				e.Rotor3.Position += 1
			}
		}
		encrypted.WriteString(current)
	}

	return encrypted.String()
}

func main() {
	// Open the JSON file
	file, err := os.Open("enigma.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var enigma enigmaMachine

	err = json.Unmarshal(byteValue, &enigma)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	input := "OGNAXRREOY"

	encrypted := encrypt(input, enigma)
	fmt.Println("Encrypted text:", encrypted)
}
