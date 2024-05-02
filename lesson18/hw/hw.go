package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// statistic("temp.txt")
	// coppy("input.txt", "output.txt")
	// concat([]string{"input.txt", "output.txt", "temp.txt"})
	// allFiles("/home/mirodil/Рабочий стол/src/github/lesson18/hw")
	// fmt.Println(searchInFile("qwerty", "concatenated.txt"))
}

/*
os.Stat() yordamida Fayl nomi,
Fayl hajmi,Ruxsatlar (oktal formatda),
Oxirgi o'zgartirish vaqt ni chiqaring
*/
func statistic(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("size --> ", info.Size(), "\nтname --> ", info.Name(), "\nlastModify --> ", info.ModTime(), "\nmode --> ", info.Mode())

}

/*
Yangi file yarating unga ma’lumot yozing va io.Copy()
yordаmida uni boshqa faylga nusxalab oling
*/

func coppy(input, output string) {
	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}

	file2, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	out, err := io.Copy(file, file2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

/*
Yangi filelar oching ularga ma’lumot yozing.
Ularni o’qish rejimida ochib ulardagi
ma’lumotlarni biriktirib
concatenated.txt nomili yangi filega yozing
*/

func concat(a []string) {
	file, err := os.Create("concatenated.txt")
	if err != nil {
		panic(err)
	}
	for _, val := range a {
		file2, err := os.Open(val)
		if err != nil {
			panic(err)
		}
		defer file2.Close()
		data, err := io.ReadAll(file2)
		if err != nil {
			panic(err)
		}
		temp, err := io.WriteString(file, string(data))
		if err != nil {
			panic(err)
		}
		fmt.Println(temp)
	}
}

/*
os.ReadDir()  yordamida katalogdagi
barcha file va kataloglar ro’yxatini
chiqaradigan dastur tuzing
*/
func allFiles(dirName string) {
	dir, err := os.ReadDir(dirName)
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
}

/*
User qidirmoqchi bo’lgan so’zini
kiritsin va shu so’z fileda bor yoki
yo’qligini aniqlaydigan dastur tuzilsin.
*/

func searchInFile(target, fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)

	if strings.Contains(string(data), target) {
		return true
	}
	return false

}
