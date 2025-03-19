package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Rotor structure
type Rotor struct {
	Position     int               `json:"position"`
	SlotPosition int               `json:"slot_position"`
	RotorMapIn   map[string]string `json:"rotor_map_in"`
	RotorMapOut  map[string]string `json:"rotor_map_out"`
}

// Plugboard structure
type Plugboard struct {
	PlugMap map[string]string `json:"plug_map"`
}

// Reflector structure
type Reflector struct {
	ReflectorMap map[string]string `json:"reflector_map"`
}

// Disk structure
type Disk struct {
	DiskMapIn  map[string]string `json:"disk_map_in"`
	DiskMapOut map[string]string `json:"disk_map_out"`
}

// EnigmaMachine structure
type EnigmaMachine struct {
	Rotor1    Rotor     `json:"rotor1"`
	Rotor2    Rotor     `json:"rotor2"`
	Rotor3    Rotor     `json:"rotor3"`
	Disk      Disk      `json:"disk"`
	Plugboard Plugboard `json:"plugboard"`
	Reflector Reflector `json:"reflector"`
}

func AtoI(current string) (int, error) {
	convertedInt, err := strconv.Atoi(current)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int: %v", err)
	}
	return convertedInt, nil
}

func outOfRangeCheck(check int) int {
	check %= 26
	if check < 0 {
		check += 26
	}
	return check
}

func processRight(r Rotor, current string) (string, error) {
	i, err := AtoI(current)
	if err != nil {
		return "", err
	}
	i = outOfRangeCheck(i + r.Position)
	key := strconv.Itoa(i)
	mappedStr, ok := r.RotorMapIn[key]
	if !ok {
		return "", fmt.Errorf("key %s not found in RotorMapIn", key)
	}
	mapped, err := AtoI(mappedStr)
	if err != nil {
		return "", err
	}
	result := outOfRangeCheck(mapped - r.Position)
	return strconv.Itoa(result), nil
}

func processLeft(r Rotor, current string) (string, error) {
	i, err := AtoI(current)
	if err != nil {
		return "", err
	}
	i = outOfRangeCheck(i + r.Position)
	key := strconv.Itoa(i)
	mappedStr, ok := r.RotorMapOut[key]
	if !ok {
		return "", fmt.Errorf("key %s not found in RotorMapOut", key)
	}
	mapped, err := AtoI(mappedStr)
	if err != nil {
		return "", err
	}
	result := outOfRangeCheck(mapped - r.Position)
	return strconv.Itoa(result), nil
}
func stepRotors(e *EnigmaMachine) {
	e.Rotor1.Position = (e.Rotor1.Position + 1) % 26
	if e.Rotor1.Position == e.Rotor1.SlotPosition {
		e.Rotor2.Position = (e.Rotor2.Position + 1) % 26
		if e.Rotor2.Position == e.Rotor2.SlotPosition {
			e.Rotor3.Position = (e.Rotor3.Position + 1) % 26
		}
	}
}

func plugboardSubstitute(pb Plugboard, char string) string {
	if val, ok := pb.PlugMap[char]; ok {
		return val
	}
	for k, v := range pb.PlugMap {
		if v == char {
			return k
		}
	}
	return char
}

func encrypt(input string, e *EnigmaMachine) (string, error) {
	input = strings.ToUpper(input)
	parts := strings.Split(input, "")
	var encrypted strings.Builder

	for _, part := range parts {
		if part < "A" || part > "Z" {
			continue
		}

		current := plugboardSubstitute(e.Plugboard, part)

		diskIn, ok := e.Disk.DiskMapIn[current]
		if !ok {
			return "", fmt.Errorf("character %s not found in DiskMapIn", current)
		}
		current = diskIn

		var err error
		if current, err = processRight(e.Rotor1, current); err != nil {
			return "", err
		}
		if current, err = processRight(e.Rotor2, current); err != nil {
			return "", err
		}
		if current, err = processRight(e.Rotor3, current); err != nil {
			return "", err
		}

		if refl, ok := e.Reflector.ReflectorMap[current]; ok {
			current = refl
		} else {
			return "", fmt.Errorf("character %s not found in ReflectorMap", current)
		}

		if current, err = processLeft(e.Rotor3, current); err != nil {
			return "", err
		}
		if current, err = processLeft(e.Rotor2, current); err != nil {
			return "", err
		}
		if current, err = processLeft(e.Rotor1, current); err != nil {
			return "", err
		}

		diskOut, ok := e.Disk.DiskMapOut[current]
		if !ok {
			return "", fmt.Errorf("character %s not found in DiskMapOut", current)
		}
		current = diskOut

		current = plugboardSubstitute(e.Plugboard, current)

		encrypted.WriteString(current)
		stepRotors(e)
	}

	return encrypted.String(), nil
}

func LoadEnigmaConfig(filePath string) (*EnigmaMachine, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var enigma EnigmaMachine
	err = json.Unmarshal(byteValue, &enigma)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	return &enigma, nil
}

func EncryptMessage(input string, config *EnigmaMachine) string {
	encrypted, err := encrypt(input, config)
	if err != nil {
		return fmt.Sprintf("Encryption error: %v", err)
	}
	return encrypted
}
