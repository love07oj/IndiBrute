# IndiBrute
![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)
![Status](https://img.shields.io/badge/status-active-brightgreen)


**IndiBrute** is a fast and efficient CLI tool written in Go to generate wordlists of Indian 10-digit mobile numbers. Tailored for ethical hacking, red teaming, and penetration testing, it allows users to create mobile-number-based wordlists for brute-force simulations.

---

## 📦 Features

- 🔢 Generates valid 10-digit Indian mobile numbers starting with prefixes: `6`, `7`, `8`, or `9`
- 📁 Splits output into ~100MB files (each with ~9.5M lines)
- ⚡ Fast execution using Go’s buffered I/O
- 🧪 Includes a test mode for rapid iteration and debugging

---

## 🧰 Project Structure

| File      | Purpose                                           |
|-----------|---------------------------------------------------|
| `main.go` | Full-scale production generator (up to 1B numbers) |
| `test.go` | Lightweight version (generates only 10K numbers)   |

---

## 🚀 Getting Started

### 🔧 Build Binary

```bash
go build -o indibrute main.go
```

### ⚙️ Run Production Mode
```bash
./indibrute --prefix 7 --output ./output
```

* Generates all numbers from `7000000000` to `7999999999`

* Output is split into ~105 files (≈100MB each)

* Files saved in `./output/`

### 🧪 Run Test Mode
```bash
go run test.go --prefix 7 --output ./test_output
```

* Generates only `10,000` numbers

* Files split into `2,000` lines each (5 files)

* Used for testing functionality and output format

---

## 🔡 Input Format
* Prefix (Required): Starting digit of mobile number — must be one of `6`, `7`, `8`, or `9`

* Output Directory (Optional): Destination to store generated `.txt` files

Example:
```bash
--prefix 8 --output ./generated
```

---

## 📂 Output Format
Each file:

* Named as `{prefix}-{file_number}.txt` (e.g. `7-1.txt`, `8-12.txt`)

* Contains numbers like:

```python-repl
7000000000
7000000001
7000000002
...
```

---

## 📊 Example: Range Breakdown for Prefix `7`
| File        | Start Number | End Number | Count                   |
|-------------|--------------|------------|--------------------------|
| `7-1.txt`   | 7000000000   | 7009532508 | 9,532,509                |
| `7-2.txt`   | 7009532509   | 7019065017 | 9,532,509                |
| `7-3.txt`   | 7019065018   | 7028597526 | 9,532,509                |
| ...         | ...          | ...        | ...                      |
| `7-104.txt` | 7980937483   | 7990479991 | 9,532,509                |
| `7-105.txt` | 7990479992   | 7999999999 | **9,520,008** (shorter)  |


Each number is exactly 10 digits and ends with a newline character (`\n`). File sizes are ~100MB except the last.

---

## 🧠 Use Cases
* Password spraying or brute-force attempts using mobile number-based credentials

* Testing login forms with default/mobile number-based usernames

* Generating wordlists for Wi-Fi WPA2 password cracking (e.g., `aircrack-ng`)

* Social engineering simulations in red team engagements

---

## 🛡️ Legal Notice

This tool is provided strictly for educational and ethical penetration testing purposes.
You must not use it on networks, applications, or systems without explicit legal authorization.

IndiBrute's authors are not responsible for any misuse or damage caused.

---
## 🤝 Contributing
Pull requests, improvements, and issues are welcome!

To contribute:

1. Fork the repo

2. Create your feature branch (`git checkout -b feature/YourFeature`)

3. Commit your changes

4. Push to the branch

5. Open a pull request

---

## 🌟 Thank You

Thank you for checking out **IndiBrute**!  
If you found this tool useful, please consider giving it a ⭐ on GitHub — it helps with visibility and motivates further development.

Happy hacking! 💻🔐
