# Cli para operações com PDF

## Objetivo

Com a nescessidade de operar com compreensão, corte e mesclagem pra tornar isso mais rapido e livre de erros de execução nasce o CLI_OP_PDF.

## Tecnologias

- Linguagens presentes:
  - Go
  - Lote(cmd do win)

- Arquitetura:
  - Diretorios:
    - cmd
    - internal
    - scripts

- Bibliotecas:
  - github.com/pdfcpu/pdfcpu v0.9.1

## Build

Para executar a build basta rodar o script `build.cmd` presente no diretorio __Scripts__.

## Modo de Uso do CLI

Sintaxe basica:
`cli_po_pdf.exe -input document.pdf -option 1 -pages 2,3`

- Parametros:
  - `-input` : responsavel por receber o arquivo original do pdf
  - `-option`: escolha das funções disponiveis.
    > 1 - Extract PDF
      2 - Merge PDF
      3 - Compress PDF
      4 - Split PDF
  - `-pages`: use os intervalos que deseja imprimir no novo pdf.
  
  __Obs:__ Somente o merge aceita uma lista separada por virgula para o `-input` : 
    >`cli_po_pdf.exe -input primeiro.pdf,segundo.pdf -option 1 -pages 2,3`
