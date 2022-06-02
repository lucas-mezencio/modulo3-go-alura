package models

import "web-crud/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	dababase := db.ConnectWithDB()

	selectTodosProdutos, err := dababase.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	var produtos []Produto

	for selectTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer dababase.Close()
	return produtos
}

func CriarProduto(nome, descricao string, preco float64, quantidade int) {
	database := db.ConnectWithDB()

	scriptInsercao, err :=
		database.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	scriptInsercao.Exec(nome, descricao, preco, quantidade)
	defer database.Close()
}

func DeletaProduto(id string) {
	database := db.ConnectWithDB()

	scriptDelecao, err := database.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}
	scriptDelecao.Exec(id)

	defer database.Close()
}
