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

	errField := canvas.NewText("", color.NRGBA{255, 0, 0, 255}) // поле для ошибки вводим на начале для удобства
	errField.TextSize = 14

	ic, err := fyne.LoadResourceFromPath("icon.png")
	if err != nil {
		fmt.Println("Error - ", err)
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

	password := widget.NewEntry() // Ввод пароля
	password.SetPlaceHolder("Пароль")

	setmale := widget.NewLabel("Укажите свой пол") // текст "Укажите свой пол"

	male := widget.NewRadioGroup([]string{"Мужской", "Женский"}, func(n string) {}) // Радиогруппа выбора пола

	approval := widget.NewCheck("Даю согласие на обработку персональных данных", func(b bool) {})

	button := widget.NewButton("Зарегистрироваться", func() {
		if setname.Text != "" && setsurn.Text != "" && login.Text != "" && password.Text != "" && male.Selected != "" && approval.Checked {
			errField.Text = ""
			fmt.Printf("Имя %s\n", setname.Text)
			fmt.Printf("Фамилия %s\n", setsurn.Text)
			fmt.Printf("Логин %s\n", login.Text)
			fmt.Printf("Пароль %s\n", password.Text)
			fmt.Printf("Пол %s\n", male.Selected)
			w.Close()
		} else {
			errField.Text = "ОШИБКА! ВЫ ЧТО ТО НЕ ВВЕЛИ"
		}
	})

	url, err := url2.Parse("https://github.com/gurhz") // URl
	if err != nil {                                    // Если ошибка != ничего
		fmt.Println("Error - ", err) // Ошибка
	}
	link := widget.NewHyperlink("Мой гитхаб", url) // Гиперссылка

	w.SetContent(container.NewVBox(
		reg,
		setname,
		setsurn,
		login,
		password,
		setmale,
		male,
		approval,
		button,
		link,
		errField,
	))

	w.ShowAndRun() // Запуск
}
