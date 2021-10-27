package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/parser/test_driver"
)

func main() {
	// open sql file
	// TODO:get path from arg
	f, err := os.Open("DDL.sql")
	if err != nil {
		log.Fatal("ERROR:could't open ddl file\n", err)
	}
	defer f.Close()
	src, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("ERROR:", err)
	}

	// TODO: parse DDLnode  https://github.com/pingcap/parser/blob/3a18f1e5dceb/ast/ddl.go
	astNode, err := parse(string(src))
	if err != nil {
		log.Fatal("ERROR: couldn't parse\n", err)
	}
	records := extract(astNode)
	fmt.Printf("RECORD: %v\n", records)

	// create csv and writer
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("ERROR: couldn't create csv file\n", err)
	}
	defer file.Close()
	cw := csv.NewWriter(file)
	defer cw.Flush()
	// write to csv
	cw.WriteAll(records)
	if err := cw.Error(); err != nil {
		log.Fatal("ERROR writing csv:\n", err)
	}
}

func parse(sql string) ([]ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}
	if len(stmtNodes) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	return stmtNodes, nil
}

type ddl struct {
	tableschema [][]string
}

func (d *ddl) Enter(in ast.Node) (ast.Node, bool) {
	// t, b := in.(*ast.TableName)
	fmt.Printf("Entered in:\n %+v\n\n", in)
	if table, ok := in.(*ast.TableName); ok {
		// col := []string{name.Table.O, name.Name.O}
		// d.tableschema = append(d.tableschema, col)
		fmt.Printf("***** GOT TABLE NAME:\n%+v\n\n", table.Name.O)
	}
	if option, ok := in.(*ast.ColumnOption); ok {
		// col := []string{name.Table.O, name.Name.O}
		// d.tableschema = append(d.tableschema, col)
		fmt.Printf("***** GOT OPTION:\n%+v\n\n", option.Expr)
	}
	if colm, ok := in.(*ast.ColumnName); ok {
		col := []string{colm.Table.O, colm.Name.O}
		d.tableschema = append(d.tableschema, col)
	}
	return in, false
}

func (v *ddl) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(nodes []ast.StmtNode) [][]string {
	d := &ddl{}
	for _, n := range nodes {
		n.Accept(d)
	}
	return d.tableschema
}
