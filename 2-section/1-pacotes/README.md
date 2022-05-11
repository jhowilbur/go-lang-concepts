Inicio projeto

* go mod init {NOME_MODULOS}

Buildar projeto

* go build

Ou se quiser rodar direto

* go run main.go

Se quiser instalar projetos/pacotes/dependencia externas

* go get github.com/badoux/checkmail

Se quiser limpar o go.mod caso nao esteja usando de determinada dependencia

* go mod tidy


---

**Details**

Se method/function comecar com letra maiuscula, permite modo "public"

Se method/function comecar com letra minuscula, so podera ser acessada dentro do pacote que ela esta
