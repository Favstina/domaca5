// program, ki omogoča dodajanje ocen studentov za posamezne dn in
// izracun koncne ocene nalog

package redovalnica

import "fmt"

// struktura student
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// doda studentu novo oceno v seznam
// pred dodajanjem preveri, ali je ocena v intervalu [0, 10]
// ce student ne obstaja, naj funkcija izpise obvestilo, da studenta ni na seznamu
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	e, ok := studenti[vpisnaStevilka]
	if ok {
		if ocena >= 0 && ocena <= 10 {
			e.Ocene = append(e.Ocene, ocena)
			// treba posodobit vnos v slovarju
			studenti[vpisnaStevilka] = e
		}
	} else {
		fmt.Println("Tega studenta ni v seznamu.")
	}
	return
}

// vrne povprecno oceno studenta
// ce student ne obstaja, naj funkcija vrne -1.0
// ce je stevilo ocen < 6 vrne povprecje 0.0
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	e, ok := studenti[vpisnaStevilka]
	var povprecje float64 = 0.0
	if ok {
		var dolzina int = len(e.Ocene)
		if dolzina >= 6 {
			for i := 0; i < dolzina; i++ {
				povprecje += float64(e.Ocene[i])
			}
			povprecje /= float64(dolzina)
		} else {
			povprecje = 0.0
		}
	} else {
		povprecje = -1.0
	}
	return povprecje
}

// izpise imena in priimke vseh studentov in njihove ocene
func IzpisVsehOcen(studenti map[string]Student) {
	for _, stu := range studenti {
		fmt.Println(stu.Ime, stu.Priimek, stu.Ocene)
	}
	return
}

// za vsakega studenta izpise njegovo ime in priimek ter povprecno oceno
// uporaba funkcije povprecje
// izpis komentarjev na podlagi povprecne ocene
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for i, stu := range studenti {
		var povp float64 = povprecje(studenti, i)
		var komentar string
		switch {
		case povp >= 9.0:
			komentar = "-> Odlicen student!"
		case povp >= 6.0:
			komentar = "-> Povprecen student"
		default:
			komentar = "-> Neuspesen Student"
		}
		fmt.Println(stu.Ime, stu.Priimek, povp, komentar)
	}
	return
}
