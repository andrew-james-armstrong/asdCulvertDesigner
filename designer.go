package main

import (
	"flag"
	"fmt"
	"github.com/andrew-james-armstrong/asd/designers"
	"github.com/andrew-james-armstrong/asd/utils"
	"log"
	"net/http"
	"os"
	"text/template"
)

var addr = flag.String("addr", ":9001", "http service address")

func main() {
	flag.Parse()
	http.HandleFunc("/", DesignSelector)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func DesignSelector(w http.ResponseWriter, req *http.Request) {
	var pat string
	pat = req.URL.Path
	switch pat {
	case "/SingleCellCulvert":
		{
			println("In single cell")
			w.Write([]byte(designers.One_Cell(req)))
		}
	case "/TwoCellCulvert":
		{
			pageTemplate, err := template.ParseFiles("result2.tmpl")
			if err != nil {
				fmt.Printf(err.Error())
			}
			pageTemplate.Execute(w, utils.TwoCellData)
		}
	case "/SingleBoxUnderpass":
		{
			pageTemplate, err := template.ParseFiles("result3.tmpl")
			if err != nil {
				fmt.Printf(err.Error())
			}
			pageTemplate.Execute(w, utils.UnderpassData)
		}
	default:
		{
			_, err := os.Stat("." + pat)
			if err != nil {
				log.Fatal(err)
			} else {
				http.ServeFile(w, req, "."+pat)
			}
		}
	}
}
