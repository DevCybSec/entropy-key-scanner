package analyzer

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/DevCybSec/entropy-key-scanner/internal/entropy"
)

// Listado de carpetas que ignoraremos por defecto
var ignoredDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	"dist":         true,
	"bin":          true,
	"obj":          true,
}

// Listado de extensiones que NO son de texto/código
var ignoredExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
	".pdf": true, ".exe": true, ".dll": true, ".so": true,
	".zip": true, ".tar": true, ".gz": true, ".7z": true,
}

func (e *Engine) DiscoverFiles(root string) error {
	// Cerramos el canal de jobs cuando terminemos de caminar
	defer close(e.Jobs)

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 1. Nivel de Directorio: Si es una carpeta ignorada, saltamos todo su contenido
		if d.IsDir() {
			if ignoredDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// 2. Nivel de Archivo: Validar extensión
		ext := strings.ToLower(filepath.Ext(path))
		if ignoredExts[ext] {
			return nil
		}

		// 3. Enviar el trabajo al worker pool
		e.Jobs <- Job{Path: path}
		return nil
	})
}

func (e *Engine) analyzeFile(path string, data []byte) {
	// 1. Análisis de Precisión (Regex en Go)
	for _, rule := range ScannerRules {
		if rule.Pattern.Match(data) {
			e.Results <- Result{
				FilePath: path,
				Findings: []string{fmt.Sprintf("[%s] %s", rule.ID, rule.Description)},
			}
			// Si encontramos un match exacto, podemos optar por saltar la entropía
			// o reportar ambos.
		}
	}

	// 2. Análisis Probabilístico (Entropía en C)
	// Solo lo hacemos si el archivo es pequeño para no degradar performance
	if len(data) < 1024*1024 { // 1MB limit
		score := entropy.GetEntropy(data)

		// Umbral de 7.0 suele indicar datos altamente aleatorios (cifrados o llaves)
		if score > 5.0 {
			e.Results <- Result{
				FilePath: path,
				Findings: []string{fmt.Sprintf("[SEC-ENT] Alta entropía detectada (%.2f). Posible secreto no identificado.", score)},
			}
		}
	}
}
