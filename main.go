package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "dbname"
)

func checkError(err error) { // прописываем ошибки
	if err != nil {
		panic(err)
	}
}

func connectTo() string { // конектимся к базе данных

	sqlConn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
	return sqlConn
}

func openDb() *sql.DB { // открываем базу данных

	db, err := sql.Open("postgres", connectTo())
	checkError(err)

	return db
}

func show_table_rez() { // показываем таблицу результаты
	rows, err := openDb().Query(`SELECT "забитые_голы", "пропущеные_голы", "дата_чемпионата" FROM "BdChemp"."результаты"`) // селект запрос
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var gol, les_gol int
		var date string

		err = rows.Scan(&gol, &les_gol, &date)
		checkError(err)

		fmt.Println(gol, les_gol, date)
	}

	checkError(err)
}

func show_table_sost_team() { // показываем таблицу состав команд
	rows, err := openDb().Query(`SELECT "фио_футболиста", "номер_футболиста", "позиция" FROM "BdChemp"."состав_команд"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var number int
		var fiO, pozc string

		err = rows.Scan(&fiO, &number, &pozc)
		checkError(err)

		fmt.Println(fiO, number, pozc)
	}

	checkError(err)
}

func show_table_team() { // показываем таблицу список команд
	rows, err := openDb().Query(`SELECT "название_команды", "фио_тренера", "страна" FROM "BdChemp"."список_команд"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {

		var name_tm, fio_tren, country string

		err = rows.Scan(&name_tm, &fio_tren, &country)
		checkError(err)

		fmt.Println(name_tm, fio_tren, country)
	}

	checkError(err)
}

func show_table_chemp() { // показываем таблицу чемпионаты мира
	rows, err := openDb().Query(`SELECT "название_чемпионата", "год_чемпионата", "страна_чемпионата" FROM "BdChemp"."чемпионаты_мира"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var nazv_ch, year, country_ch string

		err = rows.Scan(&nazv_ch, &year, &country_ch)
		checkError(err)

		fmt.Println(nazv_ch, year, country_ch)
	}

	checkError(err)
}

func delete_key_rez() { // удаляем из таблицы результаты
	var check string
	fmt.Print("какую запись удалить? введите дату чемпионата:\n") // вводим с клавиатуры значения в переменные
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdChemp"."результаты" where "дата_чемпионата"=$1` // запрос удаления из бд
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_sost_team() { // удаляем из таблицы состав
	var check int
	fmt.Print("какую запись удалить? введите номер фио футболиста:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdChemp"."состав_команд" where "фио_футболиста"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_team() { // удаляем из таблицы список
	var check string
	fmt.Print("какую запись удалить? введите фио тренера:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdChemp"."список_команд" where "фио_тренера"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_chemp() { // удаляем из таблицы чемпионат
	var check int
	fmt.Print("какую запись удалить? номер объявления:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdChemp"."чемпионаты_мира" where "название_чемпионата"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func update_table_rez() { // обновляем запись в таблице результаты
	var kol_gol, les_gol int
	var date string

	fmt.Print("Введите количество забитых голов:\n") // вводим значения в переменные с клавиатуры
	fmt.Fscan(os.Stdin, &kol_gol)

	fmt.Print("Введите пропущенных голов:\n")
	fmt.Fscan(os.Stdin, &les_gol)

	fmt.Print("введите дату чемпионата: \n")
	fmt.Fscan(os.Stdin, &date)

	updateStm := `update "BdChemp"."результаты" set "забитые_голы"=$1, "пропущеные_голы"=$2 where "дата_чемпионата"=$3` // запрос обновления данных
	_, e := openDb().Exec(updateStm, &kol_gol, &les_gol, &date)
	checkError(e)
}

func update_table_sost_team() { // обновляем запись в таблице состав команды
	var fiO, pozc string
	var number int

	fmt.Print("Введите фио футболиста для обновления:\n")
	fmt.Fscan(os.Stdin, &fiO)

	fmt.Print("Введите новый номер футболиста:\n")
	fmt.Fscan(os.Stdin, &number)

	fmt.Print("введите позицию футболиста: \n")
	fmt.Fscan(os.Stdin, &pozc)

	updateStm := `update "BdChemp"."состав_команд" set "номер_футболиста"=$1, "позиция"=$2 where "фио_футболиста"=$3`
	_, e := openDb().Exec(updateStm, &number, &pozc, &fiO)
	checkError(e)
}

func update_table_team() { // обновляем запись в таблице список команд
	var name_tm, fio_tren, country string

	fmt.Print("Введите фио тренера для обновления:\n")
	fmt.Fscan(os.Stdin, &fio_tren)

	fmt.Print("Введите новое название команды:\n")
	fmt.Fscan(os.Stdin, &name_tm)

	fmt.Print("введите новцю страну: \n")
	fmt.Fscan(os.Stdin, &country)

	updateStm := `update "BdChemp"."список_команд" set "название_команды"=$1, "страна"=$2 where "фио_тренера"=$3`
	_, e := openDb().Exec(updateStm, &name_tm, &country, &fio_tren)
	checkError(e)
}

func update_table_chemp() { // обновляем запись в таблице чемпионаты
	var nazv_ch, year, country_ch string

	fmt.Print("Введите название чемпионата:\n")
	fmt.Fscan(os.Stdin, &nazv_ch)

	fmt.Print("Введите новый год чемпионата:\n")
	fmt.Fscan(os.Stdin, &year)

	fmt.Print("введите новую страну чемпионата: \n")
	fmt.Fscan(os.Stdin, &country_ch)

	updateStm := `update "BdChemp"."чемпионаты_мира" set "год_чемпионата"=$1, "страна_чемпионата"=$2 where "название_чемпионата"=$3`
	_, e := openDb().Exec(updateStm, &year, &country_ch, &nazv_ch)
	checkError(e)
}

func add_znach_to_rez() { // добавляем значения в таблицу результат
	var kol_gol, les_gol int
	var date string
	fmt.Print("Введите количество забитых голов:\n")
	fmt.Fscan(os.Stdin, &kol_gol) // вводим значения в переменные

	fmt.Print("Введите количество пропущеных голов:\n")
	fmt.Fscan(os.Stdin, &les_gol)

	fmt.Print("Введите дату чемпионата:\n")
	fmt.Fscan(os.Stdin, &date)

	insertToDyn := `insert into "BdChemp"."результаты"("забитые_голы","пропущеные_голы","дата_чемпионата") values($1, $2, $3)` // инсерт запрос куда -> что -> какие значения
	_, e := openDb().Exec(insertToDyn, &kol_gol, &les_gol, &date)
	checkError(e)

}

func add_znach_to_sost_team() { // добавляем значения в таблицу состав команд
	var fiO, pozc string
	var number int
	fmt.Print("Введите фио футболиста: \n")
	fmt.Fscan(os.Stdin, &fiO)

	fmt.Print("Введите номер футболиста: \n")
	fmt.Fscan(os.Stdin, &number)

	fmt.Print("Введите позицию футболиста: \n")
	fmt.Fscan(os.Stdin, &pozc)

	insertToDyn := `insert into "BdChemp"."состав_команд"("фио_футболиста","номер_футболиста","позиция") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &fiO, &number, &pozc)
	checkError(e)

}

func add_znach_to_team() { // добавялем значения в таблицу список команд
	var name_tm, fio_tren, country string
	fmt.Print("Введите название команды: \n")
	fmt.Fscan(os.Stdin, &name_tm)

	fmt.Print("Введите фио тренера: \n")
	fmt.Fscan(os.Stdin, &fio_tren)

	fmt.Print("Введите страну команды: \n")
	fmt.Fscan(os.Stdin, &country)

	insertToDyn := `insert into "BdChemp"."список_команд"("название_команды","фио_тренера","страна") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &name_tm, &fio_tren, &country)
	checkError(e)

}

func add_znach_to_chemp() { // добавляем значения в таблицу чемпионаты мира
	var nazv_ch, year, country_ch string
	fmt.Print("введите название чемпионата: \n")
	fmt.Fscan(os.Stdin, &nazv_ch)

	fmt.Print("введите год чемпионата: \n")
	fmt.Fscan(os.Stdin, &year)

	fmt.Print("введите страну чемпионата: \n")
	fmt.Fscan(os.Stdin, &country_ch)

	insertToDyn := `insert into "BdChemp"."чемпионаты_мира"("название_чемпионата","год_чемпионата","страна_чемпионата") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &nazv_ch, &year, &country_ch)
	checkError(e)

}

func add_switch_case() { // функция выбора таблицы для добавления записи
	var vibor string
	fmt.Print("выберите дейсвтие: addRez -  чтобы добавить данные в таблицу дом\n addSostTeam - добавить данные в таблицу квартира\n addTeam добавить данные в таблицу квартира\n addChemp - добавить данные в таблицу объявление\n add_prod - добавить данные в таблицу продавец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "addRez":
		add_znach_to_rez()
	case "addSostTeam":
		add_znach_to_sost_team()
	case "addTeam":
		add_znach_to_team()
	case "addChemp":
		add_znach_to_chemp()
	}
}

func update_switch_case() { // функция выбора таблицы для обновления
	var vibor string
	fmt.Print("выберите дейсвтие: updateRez -  чтобы обновить данные в таблице результаты\n updateSostTeam - обновить данные в таблице состав команд\n updateTeam обновить данные в таблице список команд\n updateChemp - обновить данные в таблице чемпионаты\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "updateRez":
		update_table_rez()
	case "updateSostTeam":
		update_table_sost_team()
	case "updateTeam":
		update_table_team()
	case "updateChemp":
		update_table_chemp()
	}
}

func delete_switch_case() { // функция выбора таблицы для удаления записи
	var vibor string
	fmt.Print("выберите дейсвтие: delRez -  чтобы удалить данные из таблицы результаты\n delSostTeam - удалить данные из таблицы состав команд\n delTeam удалить данные из таблицы список команд\n delChemp - удалить данные из таблицы чемпионаты\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "delDom":
		delete_key_rez()
	case "delKvartira":
		delete_key_sost_team()
	case "delCompany":
		delete_key_team()
	case "delOble":
		delete_key_chemp()
	}
}

func show_switch_case() { // функция выбора таблицы для выводы данных
	var vibor string
	fmt.Print("выберите дейсвтие: showRez -  чтобы показать данные из таблицы результаты\n showSostTeam - показать данные из таблицы состав команд\n showTeam показать данные из таблицы список команд\n showChemp - показать данные из таблицы чемпионаты\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "showRez":
		show_table_rez()
	case "showSostTeam":
		show_table_sost_team()
	case "showTeam":
		show_table_team()
	case "showChemp":
		show_table_chemp()
	}
}

func main() {
	var v1 string
	//close db
	defer openDb().Close()
	//check db
	err := openDb().Ping()
	checkError(err)

	fmt.Print("Что вы хотите сделать?\n Чтобы добавить значения в таблицу введите addTable\n Чтобы обновить запись введите updateTable\n Чтобы удалить запись из таблицы введите deleteFromTable\n Чтобы показать данные в таблице введите showTable\n")
	fmt.Scanf("%s\n", &v1)

	switch v1 {
	case "addTable":
		add_switch_case()
	case "updateTable":
		update_switch_case()
	case "deleteFromTable":
		delete_switch_case()
	case "showTable":
		show_switch_case()
	}
}
