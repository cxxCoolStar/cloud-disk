package test

import (
	"bytes"
	"cloud-disk/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(data)

	b, err := json.Marshal(data)

	if err != nil {
		t.Fatal(err)
	}

	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dst.String())

}
