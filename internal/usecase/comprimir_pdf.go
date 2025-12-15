package usecase

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func CompressPDF(inputPath string, outputFile string) (string, error) {

	nameOutput := generatedName(outputFile, inputPath)

	err := api.OptimizeFile(inputPath, nameOutput, nil)
	if err != nil {
		return "", err
	}

	return nameOutput, nil
}
