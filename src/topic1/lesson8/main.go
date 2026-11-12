package main

import (
	"fmt"

	"github.com/ddkhugaev/go-stepik-course/src/topic1/lesson8/testmodule"
)

func main() {
	// Экспорт идентификаторов и соглашения об аббревиатурах

	// заглавная буква -> экспортируемая (публичная сущность)
	// строчная буква -> неэкспортируемая (приватная) сущность

	// Правило регистра работает для всех сущностей:
	// Переменные: PublicVar vs privateVar
	// Функции: ExportedFunc() vs internalFunc()
	// Структуры: PublicStruct vs privateStruct
	// Поля структур: User.Name (public) vs User.password (private)
	// Методы: obj.PublicMethod() vs obj.privateMethod()
	// Константы: MaxValue vs maxValue
	// Интерфейсы: Reader vs reader

	// использование публичных переменных из других пакетов
	fmt.Println(testmodule.PublicVar)

	// в go аббревиатуры пишутся
	// либо полностью заглавными, либо строчными
	// var userID int = 123               ID (не userId)
	// var HTTPServer string = "nginx"    HTTP (не HttpServer)
	// var urlPath string = "/api"        url (не uRLPath)
	// var xmlData string = "<root/>"     xml (не xMLData)
	// var apiKey string = "secret"       api (не aPIKey)

	// Аббревиатура	Правильно (экспорт)	Правильно (приват)	Неправильно
	// ID			UserID				userID				UserId, userId
	// HTTP			HTTPServer			httpServer			HttpServer
	// URL			URLPath				urlPath				UrlPath
	// XML			XMLData				xmlData				XmlData
	// API			APIKey				apiKey				ApiKey
	// JSON			JSONResponse		jsonResponse		JsonResponse
	// HTML			HTMLTemplate		htmlTemplate		HtmlTemplate
	// SQL			SQLQuery			sqlQuery			SqlQuery
	// DB			DBConnection		dbConnection		DbConnection

	// лучшие практики именования
	// camelCase для локальных переменных
	// is/has для булевых переменных
	// описательные имена для больщих scope (область видимости)
	// избегание бессмысленных имен
}
