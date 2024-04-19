package main

import (
	"github.com/labstack/echo/v4"
)

type Tech struct {
	TechName    string
	TechURL     string
	TechClasses string
}

type Project struct {
	Name        string
	URL         string
	Description string
	TechList    []Tech
	Repo        string
	Classes     string
}

type Page struct {
	Projects []Project
}

func renderPage(c echo.Context) error {
	data := Page{}
	data.Projects = []Project{
		{
			Name:        "Shoto",
			URL:         "https://www.shoto.at",
			Description: "Simple and effective URL shortoner.",
			Repo:        "https://github.com/juancwu/shoto",
			Classes:     "border-indigo-500",
			TechList: []Tech{
				{
					TechName:    "Next.js",
					TechURL:     "https://nextjs.org/",
					TechClasses: "transition bg-zinc-950 text-zinc-100 hover:bg-zinc-100 hover:text-gray-950",
				},
				{
					TechName:    "Drizzle ORM",
					TechURL:     "https://orm.drizzle.team/",
					TechClasses: "transition bg-zinc-950 text-zinc-100 hover:bg-green-600",
				},
				{
					TechName:    "Turso",
					TechURL:     "https://turso.tech/",
					TechClasses: "transition bg-zinc-950 text-zinc-100 hover:bg-teal-600",
				},
			},
		},
		{
			Name:        "Bento",
			URL:         "",
			Description: "API service to manage projects' secrets or environment variables with the ability to separate the variables into their own environment.",
			Repo:        "https://github.com/juancwu/bento",
			Classes:     "border-sky-500",
			TechList: []Tech{
				{
					TechName:    "Golang",
					TechURL:     "https://go.dev/",
					TechClasses: "transition bg-zinc-950 text-zinc-100",
				},
				{
					TechName:    "Chi",
					TechURL:     "https://go-chi.io/#/",
					TechClasses: "bg-zinc-950 text-zinc-100",
				},
				{
					TechName:    "PlanetScale",
					TechURL:     "https://planetscale.com/",
					TechClasses: "bg-zinc-950 text-zinc-100",
				},
			},
		},
		{
			Name:        "Hachi-Bitto",
			URL:         "",
			Description: "A general proramming language created just to satify YAPL, which stands for \"yet another programming language\".",
			Repo:        "https://github.com/juancwu/hachi-bitto",
			Classes:     "border-emerald-500",
			TechList: []Tech{
				{
					TechName:    "Golang",
					TechURL:     "https://go.dev/",
					TechClasses: "bg-zinc-950 text-zinc-100",
				},
			},
		},
	}

	return c.Render(200, "index.html", data)
}
