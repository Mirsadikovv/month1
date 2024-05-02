package main

import (
	"fmt"
	"sort"
)

/*
3. Movie library managemt  dasturini amalga tuzing. Dastur quyidagi funktsiyalarni o'z ichiga olishi kerak:
Quyidagi maydonlar bilan Movie  struct tuzing:
Title (string)
Director (string)
Year (int)
Rating(float)

Movie structi quyidagi methodlari bo’lishi kerak:
PrintDetails: Bu usul kinoning hamma fieldlarini chop etishi kerak.

MovieLibrary degan struct yasang (u Movie arrayni saqlashi kerak).
MovieLibrary structi quyidagi methodlari bo’lishi kerak:
AddMovie: Bu usul  MovieLibrary ga yangi kino qo'shishi kerak.
FindMovieYear: Bu usul kutubxonadagi yil bo'yicha qidirishi va kinoni qaytarishi kerak.Topolmasa error qaytarishi kerak
ListMovies: Bu usulda hamma kinolani chiqqan yili bo'yicha kamayish tartibida chiqarib berishi kerak
*/

type Movie struct {
	Title    string
	Director string
	Year     int
	Rating   float64
}

type MovieLibrary struct {
	Movies []Movie
}

func (m *Movie) PrintDetails() {
	fmt.Println(m)
}

func (lib *MovieLibrary) AddMovie(movie Movie) {
	lib.Movies = append(lib.Movies, movie)
	fmt.Println(movie, " successefully added to library")
}
func (lib *MovieLibrary) FindMovieYear(year int) *Movie {
	for i, movie := range lib.Movies {
		if movie.Year == year {
			return &lib.Movies[i]
		}
	}
	return nil
}

func (lib MovieLibrary) ListMovies() (listMovies []Movie) {
	sort.Slice(lib.Movies, func(i, j int) bool {
		return lib.Movies[i].Year > lib.Movies[j].Year
	})

	return lib.Movies

}

func main() {
	movie1 := Movie{Title: "Interstellar", Director: "Nolan", Year: 2018, Rating: 5.00}
	movie2 := Movie{Title: "Forsaj", Director: "Andrey", Year: 2011, Rating: 4.70}
	movie3 := Movie{Title: "NoFilm", Director: "NoName", Year: 2021, Rating: 3.23}
	lib := MovieLibrary{}
	lib.AddMovie(movie1)
	lib.AddMovie(movie2)
	lib.AddMovie(movie3)
	movie1.PrintDetails()
	fmt.Println(lib.FindMovieYear(2021))
	fmt.Println(lib.ListMovies())

}
