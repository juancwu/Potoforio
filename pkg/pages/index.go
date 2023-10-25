package pages

import (
    "github.com/labstack/echo/v4"
)

type Tech struct {
    TechName string
    TechURL string
    TechClasses string
}

type Project struct {
    Name string
    URL string
    Description string
    TechList []Tech
    Repo string
    Classes string
}

type Page struct {
    Projects []Project
}

func Index(c echo.Context) error {

    data := Page{}
    data.Projects = []Project{
        {
            Name: "Shoto",
            URL: "https://www.shoto.at",
            Description: "Simple and effective URL shortoner.",
            Repo: "https://github.com/juancwu/shoto",
            Classes: "border-indigo-500 hover:bg-indigo-950 hover:shadow-indigo-500/50",
            TechList: []Tech{
                {
                    TechName: "Next.js",
                    TechURL: "https://nextjs.org/",
                    TechClasses: "bg-zinc-950 text-zinc-100",
                },
                {
                    TechName: "Drizzle ORM",
                    TechURL: "https://orm.drizzle.team/",
                    TechClasses: "bg-zinc-950 text-zinc-100",
                },
                {
                    TechName: "Turso",
                    TechURL: "https://turso.tech/",
                    TechClasses: "bg-zinc-950 text-zinc-100",
                },
            },
        },
        {
            Name: "Bento",
            URL: "",
            Description: "API service to manage projects' secrets or environment variables with the ability to separate the variables into their own environment.",
            Repo: "https://github.com/juancwu/bento",
            Classes: "border-sky-500 hover:bg-sky-950 hover:shadow-sky-500/50",
            TechList: []Tech{
                {
                    TechName: "Golang",
                    TechURL: "https://go.dev/",
                    TechClasses: "bg-zinc-950 text-zinc-100",
                },
                {
                    TechName: "Chi",
                    TechURL: "https://go-chi.io/#/",
                    TechClasses: "bg-zinc-950 text-zinc-100",
                },
                {
                    TechName: "PlanetScale",
                    TechURL: "https://planetscale.com/",
                    TechClasses: "bg-zinc-950 text-zinc-100",
                },
            },
        },
        {
            Name: "Interpreter In Go",
            URL: "",
            Description: "Custom programming language interpreter written with go.",
            Repo: "https://github.com/juancwu/interpretor-in-go",
            Classes: "border-pink-500 hover:bg-pink-950 hover:shadow-pink-500/50",
            TechList: []Tech{
                {
                    TechName: "Golang",
                    TechURL: "https://go.dev/",
                    TechClasses: "bg-zinc-950 text-zinc-100",
                },
            },
        },
    }

    return c.Render(200, "index.html", data)
}
