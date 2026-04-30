package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/DevCybSec/entropy-key-scanner/internal/analyzer"
)

type FinalReport struct {
	ScannedPath string            `json:"scanned_path"`
	TotalFound  int               `json:"total_found"`
	Duration    string            `json:"duration"`
	Results     []analyzer.Result `json:"results"`
}

func main() {
	// 1. Definición de Parámetros (Flags)
	langPtr := flag.String("lang", "en", "Language for output (en, es)")
	workerPtr := flag.Int("workers", runtime.NumCPU(), "Parallel workers")
	jsonPtr := flag.Bool("json", false, "JSON format output")
	helpPtr := flag.Bool("help", false, "Show help")

	flag.Parse()

	msg, ok := analyzer.Languages[*langPtr]
	if !ok {
		msg = analyzer.Languages["en"]
	}

	if *helpPtr {
		fmt.Println("High-Perf Security Engine - v1.0.0")
		fmt.Println("Uso: high-perf-scanner [options] <directory-path>")
		flag.PrintDefaults()
		return
	}

	// 2. Validación de argumentos posicionales
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, msg.ErrorPath)
		os.Exit(1)
	}
	searchPath := args[0]

	// 3. Inicialización del Motor
	startTime := time.Now()
	engine := analyzer.NewEngine(*workerPtr)
	engine.Start()

	// 4. Recolector de Resultados
	var allResults []analyzer.Result
	var wgResults sync.WaitGroup
	wgResults.Add(1)

	go func() {
		defer wgResults.Done()
		for res := range engine.Results {
			if len(res.Findings) > 0 {
				allResults = append(allResults, res)

				// Si no es JSON, imprimimos en tiempo real para UX
				if !*jsonPtr {
					fmt.Printf(msg.FoundIn, res.FilePath)
					for _, finding := range res.Findings {
						fmt.Printf("    - %s\n", finding)
					}
				}
			}
		}
	}()

	// 5. Inicio del Descubrimiento de Archivos
	if !*jsonPtr {
		fmt.Printf(msg.StartScan, searchPath, *workerPtr)
	}

	err := engine.DiscoverFiles(searchPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, msg.ErrorWhileScanning, err)
		os.Exit(1)
	}

	// 6. Cierre Ordenado
	// Wait bloquea hasta que todos los workers terminen de procesar lo que hay en el canal Jobs
	engine.Wait.Wait()

	// Una vez que los workers terminan, cerramos el canal de resultados
	close(engine.Results)

	// Esperamos a que la goroutine de recolección termine de procesar el último resultado
	wgResults.Wait()

	duration := time.Since(startTime)

	// 7. Salida Final
	if *jsonPtr {
		report := FinalReport{
			ScannedPath: searchPath,
			TotalFound:  len(allResults),
			Duration:    duration.String(),
			Results:     allResults,
		}
		jsonData, _ := json.MarshalIndent(report, "", "  ")
		fmt.Println(string(jsonData))
	} else {
		fmt.Printf(msg.ScanComplete, duration, len(allResults))
	}
}
