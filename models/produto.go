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

	selectTodosProdutos, err := dababase.Query("select * from produtos order by id asc")
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

func EditaProduto(id string) Produto {
	database := db.ConnectWithDB()

	produtoOriginal, err := database.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoAtualizado := Produto{}

	for produtoOriginal.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoOriginal.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoAtualizado.Id = id
		produtoAtualizado.Nome = nome
		produtoAtualizado.Descricao = descricao
		produtoAtualizado.Preco = preco
		produtoAtualizado.Quantidade = quantidade
	}
	defer database.Close()
	return produtoAtualizado
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	database := db.ConnectWithDB()
	produto, err := database.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	produto.Exec(nome, descricao, preco, quantidade, id)
	defer database.Close()
}
