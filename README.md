# The Enigma Machine Simulator

This project simulates an Enigma Machine using the Go programming language. It includes a complete implementation of the Enigma encryption logic along with a graphical user interface (GUI) built using the [Fyne](https://fyne.io/) library.

The encryption algorithm is reciprocal—encrypting the plaintext produces ciphertext, and encrypting the ciphertext (with the same initial configuration) recovers the original plaintext.

---

## Project Structure

```
TheEnigmaMachine/
│
├── enigma.go         # Contains the core Enigma machine logic:
│                     #  - Rotor, plugboard, reflector, and disk processing.
│                     #  - Encryption function and configuration loader.
│
├── enigma_ui.go      # Implements the GUI using Fyne.
│                     #  - Loads the configuration from enigma.json.
│                     #  - Provides an interface for encryption/decryption.
│
├── main.go           # Entry point for the application.
│                     #  - Simply calls RunUI() from enigma_ui.go.
│
├── enigma.json       # JSON configuration file for the Enigma machine.
│                     #  - Contains settings for rotors, plugboard, reflector, and disk mappings.
│
├── go.mod            # Go module file.
└── go.sum            # Go dependencies.
```

---

## Prerequisites

- **Go** (version 1.16 or later recommended)
- [Fyne](https://fyne.io/) GUI library

Install the Fyne library with:

```sh
go get fyne.io/fyne/v2
```

---

## How to Run

1. **Clone the Repository**

   ```sh
   git clone <repository_url>
   cd TheEnigmaMachine
   ```

2. **Build and Run the Application**

   From the project root directory, run:

   ```sh
   go run .
   ```

   This command compiles and runs all Go files in the project.

3. **Using the GUI**

   - The GUI window will open with a field to enter text.
   - Enter a message (using only A–Z characters) and click the **"Encrypt/Decrypt"** button.
   - The machine's reciprocal encryption means that if you encrypt a plaintext message and then input the ciphertext (with the same starting configuration), you will recover the original plaintext.

---

## Configuration Details

The `enigma.json` file defines the settings for the Enigma machine components:

- **Rotors:**  
  Each rotor has:
  - A starting position (`position`).
  - A notch position (`slot_position`) that defines when the next rotor should step.
  - A wiring mapping from right-to-left (`rotor_map_in`) and the reverse mapping (`rotor_map_out`).  
  **Note:** The wiring keys use a 1-indexed system (from "1" to "26").

- **Plugboard:**  
  Defines letter substitutions (currently optional, as plugboard logic is commented out).

- **Reflector:**  
  Provides the fixed reciprocal mapping (keys and values in the reflector map).

- **Disk:**  
  Converts between letters and their numeric representation (1–26).  
  - `disk_map_in` converts a letter (e.g., "A") to a number (e.g., "1").
  - `disk_map_out` converts the number back to the letter.

Ensure that your JSON keys exactly match the struct tags. For example, all rotors should use `"slot_position"` (with underscore) for the notch position.

---

## Testing the Reciprocal Behavior

In the console version (in `main()` of `enigma.go`), after encrypting a sample input, the program re-encrypts the ciphertext with the same initial configuration. Because the encryption function is reciprocal, the decrypted text should match the original plaintext.

You can also test this via the GUI by entering your ciphertext back into the input field.

---
