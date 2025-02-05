package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	url2 "net/url"
)

func main() {
	a := app.New()

	w := a.NewWindow("Регистрация")
	w.Resize(fyne.NewSize(300, 360))

	ic, err := fyne.LoadResourceFromPath("icon.png")
	if err != nil {
		fmt.Println("Ошибка")
	}
	w.SetIcon(ic)

	reg := widget.NewLabel("РЕГИСТРАЦИЯ") // Текст посередине "РЕГИСТРАЦИЯ"
	reg.Alignment = fyne.TextAlignCenter

	setname := widget.NewEntry() // Ввод имени
	setname.SetPlaceHolder("Имя")
	setname.Resize(fyne.NewSize(20, 20))

	setsurn := widget.NewEntry() // Ввод фамилии
	setsurn.SetPlaceHolder("Фамилия")

	login := widget.NewEntry() // Ввод логина
	login.SetPlaceHolder("Логин")

	password := widget.NewEntry() // Ввод пароляZ
	password.SetPlaceHolder("Пароль")

	setmale := widget.NewLabel("Укажите свой пол") // текст "Укажите свой пол"

	male := widget.NewRadioGroup([]string{"Мужской", "Женский"}, func(n string) {}) // Радиогруппа выбора пола

	button := widget.NewButton("Зарегистрироваться", func() {
		fmt.Printf("Имя %s\n", setname.Text)
		fmt.Printf("Фамилия %s\n", setsurn.Text)
		fmt.Printf("Логин %s\n", login.Text)
		fmt.Printf("Пароль %s\n", password.Text)
		fmt.Printf("Пол %s\n", male.Selected)
	})

	url, err := url2.Parse("https://github.com/gurhz") // URl
	if err != nil {                                    // Если ошибка != ничего
		fmt.Println("Ошибка! Страница не существует или автор поменял никнейм") // Ошибка
	}
	link := widget.NewHyperlink("Мой гитхаб", url) // Гиперссылка

	errField := canvas.NewText("", color.NRGBA{255, 0, 0, 255})
	errField.TextSize = 14

	w.SetContent(container.NewVBox(
		reg,
		setname,
		setsurn,
		login,
		password,
		setmale,
		male,
		button,
		link,
		errField,
	))

	w.ShowAndRun() // Запуск
}
