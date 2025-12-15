package usecase

import (
	"fmt"
	"log"
	"os"

	"github.com/dossantoscarlos/genpdf_cli/internal/model"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func SplitPDF(params model.Pdf) ([]string, error) {
	var err error

	log.Default().Println("SplitPDF")

	if len(params.PageRange) == 0 {
		return nil, fmt.Errorf("error ao extrair pagina")
	}

	if len(params.PageRange) == 1 && params.PageRange[0] == "all" {
		params.PageRange = nil
	}

	err = api.ExtractPagesFile(
		params.InputValueName, params.OutputValueName,
		params.PageRange, nil)
	if err != nil {
		return nil, fmt.Errorf("1%v", err.Error())
	}

	// 2. Fazer o merge dos arquivos no diretório
	filesToMerge := []string{}
	directory := params.OutputValueName

	// Listar todos os arquivos no diretório
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o diretório: %v", err)
	}

	// Adicionar arquivos ao slice filesToMerge na ordem correta
	for _, file := range files {
		if !file.IsDir() {
			filesToMerge = append(filesToMerge, fmt.Sprintf("%s/%s", directory, file.Name()))
		}
	}

	return filesToMerge, nil
}
