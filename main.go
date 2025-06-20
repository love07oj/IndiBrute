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
        linesPerFile = 9532509 // ~100MB per file
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
        end, _ := strconv.ParseInt(*prefix+"999999999", 10, 64)
        totalNumbers := end - start + 1
        totalFiles := totalNumbers / linesPerFile
        if totalNumbers%linesPerFile != 0 {
                totalFiles++
        }

        fmt.Printf("Generating numbers from %d to %d in %d files...\n", start, end, totalFiles)

        current := start
        for i := int64(1); i <= totalFiles; i++ {
                fileName := filepath.Join(*outputDir, fmt.Sprintf("%s-%d.txt", *prefix, i))
                file, err := os.Create(fileName)
                if err != nil {
                        log.Fatalf("Failed to create file: %v", err)
                }
                writer := bufio.NewWriter(file)

                count := int64(0)
                for count < linesPerFile && current <= end {
                        _, err := fmt.Fprintf(writer, "%010d\n", current)
                        if err != nil {
                                log.Fatalf("Failed to write to file: %v", err)
                        }
                        current++
                        count++
                }

                writer.Flush()
                file.Close()
                fmt.Printf("Created: %s [%d numbers]\n", fileName, count)
        }

        fmt.Println("✅ Done!")
}

func isValidPrefix(p string) bool {
        return p == "6" || p == "7" || p == "8" || p == "9"
}
