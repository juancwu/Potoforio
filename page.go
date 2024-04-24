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
			Name:        "HawkHacks Website",
			URL:         "https://hawkhacks.ca",
			Description: "HawkHacks is a 36 hour in-person hackathon hosted at Wilfrid Laurier University.",
			Repo:        "https://github.com/LaurierHawkHacks/Landing",
			Classes:     "border-orange-300",
			TechList: []Tech{
				{
					TechName:    "React",
					TechURL:     "https://react.dev/",
					TechClasses: "",
				},
				{
					TechName:    "Firebase",
					TechURL:     "https://firebase.google.com/",
					TechClasses: "",
				},
				{
					TechName:    "TailwindCSS",
					TechURL:     "https://tailwindcss.com/",
					TechClasses: "",
				},
			},
		},
		{
			Name:        "LCS Website",
			URL:         "https://lauriercs.ca",
			Description: "The official Computer Science club at Wilfrid Laurier University website.",
			Repo:        "https://github.com/LaurierCS/Website",
			Classes:     "border-cyan-400",
			TechList: []Tech{
				{
					TechName:    "React",
					TechURL:     "https://react.dev/",
					TechClasses: "",
				},
				{
					TechName:    "Firebase",
					TechURL:     "https://firebase.google.com/",
					TechClasses: "",
				},
				{
					TechName:    "Mantine UI",
					TechURL:     "https://mantine.dev/",
					TechClasses: "",
				},
			},
		},
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
			Name:        "Konbini",
			URL:         "",
			Description: "A \"convenient store\" with API services to manage projects' secrets or environment variables with the ability to separate the variables into their own environment.",
			Repo:        "https://github.com/juancwu/konbini",
			Classes:     "border-sky-500",
			TechList: []Tech{
				{
					TechName:    "Golang",
					TechURL:     "https://go.dev/",
					TechClasses: "transition bg-zinc-950 text-zinc-100",
				},
				{
					TechName:    "Echo",
					TechURL:     "https://echo.labstack.com/",
					TechClasses: "bg-zinc-950 text-zinc-100",
				},
				{
					TechName:    "PostgreSQL",
					TechURL:     "https://www.postgresql.org/",
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
		{
			Name:        "This Site",
			URL:         "https://juancwu.dev",
			Description: "My personal portolio site?",
			Repo:        "https://github.com/juancwu/potoforio",
			Classes:     "border-gray-600",
			TechList: []Tech{
				{
					TechName:    "Golang",
					TechURL:     "https://go.dev/",
					TechClasses: "bg-zinc-950 text-zinc-100",
				},
				{
					TechName:    "HTMX",
					TechURL:     "https://htmx.org/",
					TechClasses: "bg-zinc-950 text-zinc-100",
				},
				{
					TechName:    "Hyperscript",
					TechURL:     "https://hyperscript.org/",
					TechClasses: "bg-zinc-950 text-zinc-100",
				},
			},
		},
	}

	return c.Render(200, "index.html", data)
}
