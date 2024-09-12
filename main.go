package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("halo")

	e := echo.New()

	e.GET("/", indexHandler)

	if err := e.Start(":5000"); err != nil {
		log.Fatal(err)
	}
}

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate(htmlFilesPath []string) *Templates {
	return &Templates{
		templates: template.Must(template.ParseFiles(htmlFilesPath...)),
	}
}

type AnObject struct {
	Name        string
	Description string
}

type IndexVM struct {
	Objects  []AnObject
	NextPage int
}

var dataVM = []AnObject{
	{Name: "Object 1", Description: "Object 1 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 2", Description: "Object 2 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 3", Description: "Object 3 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 4", Description: "Object 4 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 5", Description: "Object 5 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 6", Description: "Object 6 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 7", Description: "Object 7 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 8", Description: "Object 8 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 9", Description: "Object 9 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 10", Description: "Object 10 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 11", Description: "Object 11 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 12", Description: "Object 12 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 13", Description: "Object 13 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 14", Description: "Object 14 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 15", Description: "Object 15 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 16", Description: "Object 16 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
	{Name: "Object 17", Description: "Object 17 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil."},
}

func getPagedObjects(page int, max_count int) ([]AnObject, error) {
	output := []AnObject{}

	start_offset := 0

	if page > 1 {
		start_offset = (page - 1) * max_count
	}

	end_offset := start_offset + max_count

	if len(dataVM) < end_offset {
		end_offset = len(dataVM)
	}

	output = append(output, dataVM[start_offset:end_offset]...)

	return output, nil
}

func indexHandler(c echo.Context) error {
	fmt.Printf("Index page called")
	param_page := c.QueryParam("page")
	page := 1
	var err error
	if param_page != "" {
		page, err = strconv.Atoi(param_page)
		if err != nil {
			fmt.Printf("Error converint page from query params")
			page = 1
		}
	}

	listOfObjects, err := getPagedObjects(page, 5)

	if err != nil {
		panic(err)
	}

	indexVM := IndexVM{
		Objects:  listOfObjects,
		NextPage: page + 1,
	}

	templates := []string{"./base.html", "./list.html"}
	c.Echo().Renderer = NewTemplate(templates)

	return c.Render(http.StatusOK, "base.html", indexVM)
}
