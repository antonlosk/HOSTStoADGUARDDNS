# HOSTS to AdGuard DNS Converter

A lightweight and fast Go-based command-line utility for converting a standard `hosts` file into DNS rewrite rules format for AdGuard.

## 🚀 How it works

The program reads the source `hosts.txt` file, carefully preserves all comments and blank lines, and converts lines containing IP addresses and domains into AdGuard syntax.

**Conversion example:**
* **Input (`hosts.txt`):**  
  `192.168.1.1 example.com`
* **Output (`hosts_ADGUARDDNS.txt`):**  
  `|example.com^$dnsrewrite=192.168.1.1`

## 📦 Usage (Precompiled file)

1. Place your `hosts.txt` file in the same folder as the program.
2. Run `HOSTStoADGUARDDNS.exe`.
3. Take the finished result from the newly created `hosts_ADGUARDDNS.txt` file.

## 🛠 Building from source

If you want to compile the program yourself, you will need [Go](https://go.dev/) installed.

The project includes a convenient script for Windows:
1. Download the source code.
2. Run `build.bat`.
3. The script will automatically build an optimized `.exe` file without extra debugging information.

**Manual build (Linux / macOS / Windows):**
```bash
go build -ldflags "-s -w" -trimpath -o HOSTStoADGUARDDNS main.go
```
---

## 👨‍💻 Author

**antonlosk**

---

## 🛠️ Tools Used

**Google Gemini 3.1 Pro** — used for code generation, refactoring, and documentation writing.
