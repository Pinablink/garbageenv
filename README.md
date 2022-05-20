
# 💻 Garbageenv

</br>
👉 É um serviço interno que tem como função deletar recursos que estão fora da validade, diretórios e arquivos. Essa é uma das ferramentas que o GGIZ emprega como controle de Saneamento Ambiental!!
</br>
</br>
 
## ⚠️👀 Importante 👀⚠️
</br>
Este serviço esta sendo disponibilizado para consulta didática. Os testes realizados atendem a utilização na plataforma GGIZ e talvez precise de ajuste para sua realidade. A evolução ocorrerá conforme a necessidade da plataforma.
</br>
</br>

### Dependências
</br>
As dependências do Garbageenv são de pacotes do GGIZ e de terceiros.
</br>
- github.com/satori/go.uuid
</br>
- g2ssms/send
</br>
- github.com/Pinablink/mTools/mtime
</br>
- github.com/Pinablink/warning
</br>
- github.com/Pinablink/lingue
</br>
</br>


### Modo básico de Uso:
_____

A interface é por linha de comando. Abaixo um exemplo no Prompt Linux. Segue os parâmetros de inicialização. 
</br>
</br>


```
./garbageenv -p /root/sibilussolicdown/arq -bh 4

```

### 🤔 Uma Ajudinha Rápida
_____

### 👉 Parâmetros de Utilização
-p  : Informa o local onde se encontra o diretório que será limpo.
</br>
-bh  : Informa a(s) hora(s) anterior ao momento de seu processamento.
</br> 
-bm  : Informa o(s) minuto(s) anterior ao momento de seu processamento                
</br>	

### 👉 Exemplo:
Desejo que sejam deletados todos os arquivos e diretórios criados a 10 horas atrás 
no meu diretório 
</br>
garbageenv -p <CAMINHO_DIRETORIO> -bh 10



