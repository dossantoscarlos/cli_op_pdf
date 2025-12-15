package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	execapp "github.com/dossantoscarlos/genpdf_cli/internal/exec_app"
	"github.com/dossantoscarlos/genpdf_cli/internal/model"
)

func insertData(input *string, options *string, pages *string) {
	flag.StringVar(input, "input", "", "Arquivo PDF de entrada (obrigatório)")
	flag.StringVar(options, "option", "", "1:Extract PDF 2:Merge PDF 3:Compress PDF 4:Split PDF")
	flag.StringVar(pages, "pages", "", "Páginas a extrair (ex: 1,2,5)")
	flag.Parse()
}

func main() {
	var err error
	var input string
	var options string
	var pages string

	insertData(&input, &options, &pages)

	if input == "" || pages == "" {
		flag.PrintDefaults()
		log.Default().Fatalln("input ou pages vazio")
		return
	}

	log.Default().Println(input)
	log.Default().Println(pages)
	log.Default().Println(options)

	option, err := strconv.Atoi(options)
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("error option: %v", err.Error())
		return
	}

	if option < 1 || option > 4 {
		flag.PrintDefaults()
		log.Fatalln("Opcao invalida")
		return
	}

	pageRanges := strings.Split(pages, ",")

	pdf := model.Pdf{
		InputValueName: input,
		PageRange:      pageRanges,
	}

	switch option {
	case 1:
		execapp.ExtractPDF(pdf)
	case 3:
		execapp.CompressPDF(pdf)
	case 2:
		execapp.MergePDF(pdf)
	case 4:
		execapp.SplitPDF(pdf)
	}
}
