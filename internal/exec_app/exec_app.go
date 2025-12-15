package execapp

import (
	"log"
	"strings"

	"github.com/dossantoscarlos/genpdf_cli/internal/model"
	"github.com/dossantoscarlos/genpdf_cli/internal/usecase"
	"github.com/dossantoscarlos/genpdf_cli/internal/util"
)

func MergePDF(pdf model.Pdf) {
	var path string

	pdf.OutputValueName = "merge_pdf" + util.SeparatorPath()

	if err := util.VerificaDirectory(pdf.OutputValueName); err != nil {
		log.Default().Fatalf("error em verificar diretorio: %v\n", err)
		return
	}

	if util.IsDirectory(pdf.InputValueName) {
		var err error

		path = pdf.InputValueName

		pdf.InputValueName, err = util.Files(path)
		log.Default().Println(pdf.InputValueName)

		if err != nil {
			log.Default().Fatalf("lista de arquivo falhou: %v", err)
			return
		}
	}

	inputValueName := strings.Split(pdf.InputValueName, ",")

	log.Default().Printf("\n %s \n", inputValueName[1])

	nameOutput, err := usecase.MergePDFs(inputValueName, pdf.OutputValueName)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	usecase.OpenFileInBrowser(nameOutput)
}

func ExtractPDF(pdf model.Pdf) {
	pdf.OutputValueName = "pdf_extract" + util.SeparatorPath()
	util.VerificaDirectory(pdf.OutputValueName)
	nameOutput, err := usecase.ExtractPages(pdf)
	if err != nil {
		log.Fatalf("Erro ao extrair páginas: %v", err)
		return
	}
	usecase.OpenFileInBrowser(nameOutput)
}

func CompressPDF(pdf model.Pdf) {
	pdf.OutputValueName = "compress_pdf" + util.SeparatorPath()
	util.VerificaDirectory(pdf.OutputValueName)
	nameOutput, err := usecase.CompressPDF(pdf.InputValueName, pdf.OutputValueName)
	if err != nil {
		log.Fatalf("Erro ao comprimir o PDF: %v", err)
		return
	}
	usecase.OpenFileInBrowser(nameOutput)
}

func SplitPDF(pdf model.Pdf) {
	pdf.OutputValueName = "split_pdf" + util.SeparatorPath()
	util.VerificaDirectory(pdf.OutputValueName)
	_, err := usecase.SplitPDF(pdf)
	if err != nil {
		log.Fatalf("Erro ao dividir PDF: %v", err)
		return
	}
	usecase.OpenFolder(pdf.OutputValueName)
}
