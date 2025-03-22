package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	url2 "net/url"
	"strings"
	"unicode"
)

func removeSpaces(entry *widget.Entry) {
	entry.OnChanged = func(s string) {
		if strings.Contains(s, " ") {
			// Удаляем пробелы
			entry.SetText(strings.ReplaceAll(s, " ", ""))
		}
	}
}

func isValidLoginOrPassword(input string) bool {
	// Задаем набор недопустимых символов
	for _, char := range input {
		if !isAllowed(char) {
			return false
		}
	}
	return true
}

func isAllowed(char rune) bool {
	// Разрешаем только буквы, цифры и некоторые специальные символы
	return unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_' || char == '-'
}

func main() {
	a := app.New()

	w := a.NewWindow("Регистрация")
	w.Resize(fyne.NewSize(400, 460))

	LightTheme := fyne.NewMenuItem("Светлая", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	DarkTheme := fyne.NewMenuItem("Тёмная", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})
	menu := fyne.NewMenu("Тема", LightTheme, DarkTheme)
	mainMenu := fyne.NewMainMenu(menu)
	w.SetMainMenu(mainMenu)

	ic, err := fyne.LoadResourceFromPath("icon.png")
	if err != nil {
		fmt.Println("Error - ", err)
	}
	w.SetIcon(ic)

	reg := widget.NewLabel("РЕГИСТРАЦИЯ") // Текст посередине "РЕГИСТРАЦИЯ"
	reg.Alignment = fyne.TextAlignCenter

	username := widget.NewEntry() // Ввод имени
	username.SetPlaceHolder("Имя пользователя")
	removeSpaces(username)

	password := widget.NewPasswordEntry() // Ввод пароля
	password.SetPlaceHolder("Пароль")
	removeSpaces(password)

	email := widget.NewEntry() // Ввод логина
	email.SetPlaceHolder("Почта")
	removeSpaces(email)

	emailOptions := []string{"@gmail.com", "@mail.ru", "@yandex.ru"}
	mails := widget.NewSelect(emailOptions, nil)
	mails.SetSelected(emailOptions[0])

	setmale := widget.NewLabel("Укажите свой пол") // текст "Укажите свой пол"

	male := widget.NewRadioGroup([]string{"Мужской", "Женский"}, func(n string) {}) // Радиогруппа выбора пола

	approval := widget.NewCheck("Даю согласие на обработку персональных данных", func(b bool) {})

	button := widget.NewButton("Зарегистрироваться", func() {
		if username.Text != "" && email.Text != "" && password.Text != "" && male.Selected != "" &&
			approval.Checked && isValidLoginOrPassword(password.Text) && isValidLoginOrPassword(email.Text) {
			fmt.Printf("Имя %s\n", username.Text)
			fmt.Printf("Почта %s\n", email.Text+mails.Selected)
			fmt.Printf("Пароль %s\n", password.Text)
			fmt.Printf("Пол %s\n", male.Selected)
			w.Close()
		} else {
			dialog.ShowInformation("Упс!", "По всей видимости вы что-то забыли записать!", w)
		}
	})

	url, err := url2.Parse("https://github.com/gurhz") // URl
	if err != nil {                                    // Если ошибка != ничего
		log.Println("Error - ", err) // Ошибка
	}
	link := widget.NewHyperlink("Мой гитхаб", url) // Гиперссылка

	w.SetContent(container.NewVBox(
		reg,
		username,
		password,
		container.New(layout.NewGridLayout(2), email, mails),
		setmale,
		male,
		approval,
		button,
		link,
	))

	w.ShowAndRun() // Запуск
}
