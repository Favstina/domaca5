package main

import (
	"context"
	"log"
	"os"

	"github.com/Favstina/domaca5/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "Redovalnica",
		Usage: "Upravljanje z ocenami",

		// stikala
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša možna ocena",
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja možna ocena",
				Value: 10,
			},
		},

		Commands: []*cli.Command{
			{
				Name:  "izpis",
				Usage: "Izpiši vse ocene",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					studenti := loadTestData()
					redovalnica.IzpisVsehOcen(studenti)
					return nil
				},
			},

			{
				Name:  "uspeh",
				Usage: "Izpiši končni uspeh vseh študentov",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					studenti := loadTestData()
					redovalnica.IzpisiKoncniUspeh(studenti)
					return nil
				},
			},

			{
				Name:  "dodaj",
				Usage: "Dodaj oceno študentu",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "vpisna", Usage: "Vpisna številka"},
					&cli.IntFlag{Name: "ocena", Usage: "Nova ocena"},
				},
				Action: func(ctx context.Context, cmd *cliCommand) error {
					vpisna := cmd.String("vpisna")
					ocena := cmd.Int("ocena")

					// parametri iz root level flags
					min := cmd.Parent().Int("minOcena")
					max := cmd.Parent().Int("maxOcena")

					studenti := loadTestData()
					redovalnica.DodajOceno(studenti, vpisna, ocena)

					return nil

				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func primerRedovalnice() {
	studenti := make(map[string]redovalnica.Student)

	prviStu := Student{"Anja", "Blasko", []int{2, 4, 8, 6, 9, 3, 7, 10}}
	drugiStu := Student{"Rok", "Neki", []int{5, 7, 6, 8, 5}}
	tretjiStu := Student{"Tomaz", "Priimek", []int{10, 9, 10, 8, 9, 8, 9, 10, 9}}

	studenti["1"] = prviStu
	studenti["2"] = drugiStu
	studenti["3"] = tretjiStu
}
