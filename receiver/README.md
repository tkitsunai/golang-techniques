# Pointer receiver VS Value receiver

### Guidelines

Ref: https://golang.org/doc/faq#Pointers

#### ポインターレシーバーを使う

```go
type PointerReceiver struct {
}

func (p *PointerReceiver) Address() string {
return fmt.Sprintf("%p", p)
}
```

値レシーバー呼び出し時に構造体はコピーされてしまう。

以下のコードでは `NAME1` となりSetNameを呼んだ先で自分自身の構造体を参照した場合は呼び出し時点でコピーされる。

```go
type ValueReceiver struct {
	Name string
}

func (v ValueReceiver) SetName(name string) {
	v.Name = name
}
```

```go
    valueReceiver := receiver.ValueReceiver{
        Name: "NAME1",
    }

    valueReceiver.SetName("test")

    assert.Equal(t, "test", valueReceiver.Name) <- FAILED
    assert.Equal(t, "NAME1", valueReceiver.Name) <- PASS
```

レシーバ呼び出し時にImmutableに敢えてしたいという以外、値レシーバーを使う必要性は薄い。
