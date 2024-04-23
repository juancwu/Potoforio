package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
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

func renderPage(c echo.Context, conn *pgx.Conn) error {
	data := Page{}

	// query a project
	project := Project{}
	row := conn.QueryRow(context.Background(), "SELECT name, repository_url, deployment_url, description FROM projects WHERE id = 1;")
	if err := row.Scan(&project.Name, &project.Repo, &project.URL, &project.Description); err != nil {
		log.Printf("Error querying project: %v\n", err)
	} else {
		log.Printf("Project: %v\n", project)
	}

	// query technologies
	techs := []Tech{}
	rows, err := conn.Query(context.Background(), "SELECT technologies.name, technologies.url FROM technologies INNER JOIN projects_technologies ON technologies.id = projects_technologies.technology_id WHERE projects_technologies.project_id = 1;")
	if err != nil {
		log.Printf("Error querying technologies: %v\n", err)
	} else {
		for rows.Next() {
			tech := Tech{}
			err = rows.Scan(&tech.TechName, &tech.TechURL)
			if err != nil {
				log.Printf("Error scanning technology: %v\n", err)
			} else {
				techs = append(techs, tech)
			}
		}
	}
	defer rows.Close()
	log.Printf("All technologies: %v\n", techs)

	project.TechList = techs
	project.Classes = "border-indigo-500"

	data.Projects = []Project{
		project,
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
