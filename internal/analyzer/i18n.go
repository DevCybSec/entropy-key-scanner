package analyzer

type Messages struct {
	StartScan          string
	FoundIn            string
	HighEntropy        string
	ScanComplete       string
	FilesFound         string
	ErrorPath          string
	ErrorWhileScanning string
}

var Languages = map[string]Messages{
	"en": {
		StartScan:          "Starting scan in: %s with %d workers...",
		FoundIn:            "\n[!] Findings in: %s",
		HighEntropy:        "[SEC-ENT] High entropy detected (%.2f). Possible unidentified secret.",
		ScanComplete:       "\nScan completed in %v. Total findings: %d",
		FilesFound:         "Analyzing: %s",
		ErrorPath:          "Error: Should specify a path to scan\n",
		ErrorWhileScanning: "Error whilst scanning: %v\n",
	},
	"es": {
		StartScan:          "Iniciando escaneo en: %s con %d workers...",
		FoundIn:            "\n[!] Hallazgos en: %s",
		HighEntropy:        "[SEC-ENT] Alta entropía detectada (%.2f). Posible secreto no identificado.",
		ScanComplete:       "\nEscaneo completado en %v. Total hallazgos: %d",
		FilesFound:         "Analizando: %s",
		ErrorPath:          "Error: Debes especificar un directorio para escanear.\n",
		ErrorWhileScanning: "Error durante el escaneo: %v\n",
	},
}
