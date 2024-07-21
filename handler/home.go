package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Jucaza1/goth-hotel-web/view/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	server string
}

var items []string
var texts []string

type filejsonSt struct {
	textStruct
}

type textStruct struct {
	Section []SectionSt `json:"section"`
}
type SectionSt struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"paragraph"`
}

func NewHomeHandler(server string) *HomeHandler {
	return &HomeHandler{
		server: server,
	}
}

func (h *HomeHandler) HandleHomeShow(c echo.Context) error {

	return render(c, 200, home.Show(texts, items))
}
func init() {
	for i := 0; i < 12; i++ {
		item := fmt.Sprintf("public/img/landing/land%d.jpg", i+1)
		fmt.Println(item)
		items = append(items, item)

	}
	file1, err := os.Open("./public/text/landing/file1.json")
	if err != nil {
		panic(err)
	}
	fileBytes, err := io.ReadAll(file1)
	if err != nil {
		panic(err)
	}
	var parsedfile1 filejsonSt
	if err := json.NewDecoder(bytes.NewReader(fileBytes)).Decode(&parsedfile1); err != nil {
		panic(err)
	}
	texts = append(texts, parsedfile1.Section[0].Title)
	fmt.Println(texts)
	texts = append(texts, parsedfile1.Section[0].Paragraph...)
	fmt.Printf("title= %s, text= %s", texts[0], texts[1])
}
