package usecase

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func OpenFolder(path string) error {
	var cmd *exec.Cmd

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("erro ao obter o diretório de trabalho: %v", err)
		return err
	}

	path = currentDir + string(os.PathSeparator) + path

	// Verifica o sistema operacional
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", path) // Para Windows
	case "darwin":
		cmd = exec.Command("open", path) // Para macOS
	case "linux":
		cmd = exec.Command("xdg-open", path) // Para Linux
	default:
		return fmt.Errorf("sistema operacional não suportado")
	}

	return cmd.Start()
}
