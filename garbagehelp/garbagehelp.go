package garbagehelp

import (
	"fmt"

	"github.com/Pinablink/lingue"
	"github.com/Pinablink/lingue/oslabel"
)

var queryHelp string = `         
              !Ajuda! 
Parâmetros de Utilização
-p <DIR> : Informa o local onde se encontra o diretório que será limpo
-bh <number> : Informa a(s) hora(s) anterior ao momento de seu processamento 
-bm <number> : Informa o(s) minuto(s) anterior ao momento de seu processamento                
						 
Exemplo:
Desejo que seja deletado todos os arquivos e diretórios criados a 10 horas atrás 
no meu diretório <CAMINHO_DIRETORIO>
GarbageEnv -p <CAMINHO_DIRETORIO> -bh 10
					
`

var queryHelpWarning string = `         
              !Ajuda! 
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
               Atenção
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
Você não informou corretamente a sua parametrização.
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++			  
Parâmetros de Utilização
-p <DIR> : Informa o local onde se encontra o diretório que será limpo
-bh <number> : Informa a(s) hora(s) anterior ao momento de seu processamento 
-bm <number> : Informa o(s) minuto(s) anterior ao momento de seu processamento. Não obrigatorio                
						 
Exemplo:
Desejo que seja deletado todos os arquivos e diretórios criados a 10 horas atrás 
no meu diretório <CAMINHO_DIRETORIO>
GarbageEnv -p <CAMINHO_DIRETORIO> -bh 10
					
`

//
func Showhelp() {
	lingue.NewLingue().ExecCommand(oslabel.CLEAR_CMD)
	fmt.Println(queryHelp)
}

func WarningShowHelp() {
	lingue.NewLingue().ExecCommand(oslabel.CLEAR_CMD)
	fmt.Println(queryHelpWarning)
}
