package main

import (
        "bufio"
        "flag"
        "fmt"
        "log"
        "os"
        "path/filepath"
        "strconv"
)

const (
        linesPerFile = 2000        // For testing: 2,000 lines per file
        totalLines   = 10000       // For testing: generate 10,000 numbers total
)

func main() {
        prefix := flag.String("prefix", "", "Starting digit of mobile number (6, 7, 8, or 9)")
        outputDir := flag.String("output", "output", "Directory to save generated files")
        flag.Parse()

        if *prefix == "" || len(*prefix) != 1 || !isValidPrefix(*prefix) {
                log.Fatal("Please provide a valid prefix: 6, 7, 8, or 9")
        }

        err := os.MkdirAll(*outputDir, os.ModePerm)
        if err != nil {
                log.Fatalf("Failed to create output directory: %v", err)
        }

        start, _ := strconv.ParseInt(*prefix+"000000000", 10, 64)
        current := start

        totalFiles := totalLines / linesPerFile
        if totalLines%linesPerFile != 0 {
                totalFiles++
        }

        fmt.Printf("🔧 Test mode: generating %d numbers in %d files...\n", totalLines, totalFiles)

        for i := 1; i <= totalFiles; i++ {
                filename := filepath.Join(*outputDir, fmt.Sprintf("%s-%d.txt", *prefix, i))
                file, err := os.Create(filename)
                if err != nil {
                        log.Fatalf("Failed to create file: %v", err)
                }
                writer := bufio.NewWriter(file)

                count := 0
                for count < linesPerFile && current < start+totalLines {
                        fmt.Fprintf(writer, "%010d\n", current)
                        current++
                        count++
                }

                writer.Flush()
                file.Close()
                fmt.Printf("✅ Created: %s [%d lines]\n", filename, count)
        }

        fmt.Println("🎉 Test generation complete!")
}

func isValidPrefix(p string) bool {
        return p == "6" || p == "7" || p == "8" || p == "9"
}
