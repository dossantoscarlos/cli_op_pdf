package usecase

import (
	"log"
	"strings"

	"github.com/dossantoscarlos/genpdf_cli/internal/util"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePDFs(inputFiles []string, outputFile string) (string, error) {

	nameFiles := inputFiles[0]
	log.Default().Println(inputFiles[1])
	separator := util.SeparatorPath()
	log.Default().Println(separator)

	name := strings.Split(nameFiles, separator)

	log.Default().Println(name)

	outputFile = generatedName(outputFile, name[2])

	err := api.MergeCreateFile(inputFiles, outputFile, false, nil)

	if err != nil {
		return "", err
	}

	return outputFile, nil
}
