# First Class Collection (FCC)

First Class Collection(以下FCC)は「ThoughtWorksアンソロジー」で紹介されたオブジェクト指向言語におけるテクニック

言語に用意されたArray、Listといった型に対する操作・処理をまとめ、名前付けを行えたり、処理を制限させるリストを作ることができる

### Guidelines

#### Type defineを使ったFirst Class Collectionを使う

とある構造体に対して、配列 `[]User` を Type define機能を使うことで、名前付けと型を制限することができる

```go
type User struct {
	ID   string
	Name string
}

type Users []User
```

User配列に対する操作は `Users` を経由する。

フィルター機能を作りたい場合は、以下のようなレシーバーを定義することで、

```go
func (u Users) Filter(condition func(u User) bool) Users {
	result := u[:0]
	for _, user := range u {
		if condition(user) {
			result = append(result, user)
		}
	}
	return result
}
```

ポイントは値レシーバにすることで、元の配列状態をImmutableに扱えるようにする。オリジナルの配列自体へ操作することはせずに新しい `Users` を返すように作ることでChainしながら呼び出すこともできる。

Sliceを作成する場合はメモリアロケーションに気をつけなければいけないが、殆どのケースにおいてはこれで対応できる。
