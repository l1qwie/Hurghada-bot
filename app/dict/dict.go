package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["sayHello"] = `<b>Инструкция к использованию бота:</b>


<b>Главные команды:</b>

/admin - позволяет просматривать весь функционал доступный организатору
/client - позволяет просматривать весь функционал доступный клиентам
/refresh - после получения этой команды бот нначнет читать таблицу в Google Sheets и если найдет что то новое (не прочитанное ранее) отправит новые данные в телеграм-канал t.me/testdvijhurghada

<b>Update 1.0</b>
 - Добавлены команды /admin и /client. 
 - Разработан и протестирован весь функционал
<b>Update 1.1</b>
 - Добавлены Google Forms для более быстрого заполнения информации
 - Реализованн конект Google Forms и Google Sheets
 - Реализован конект бота с Google Sheets
 - Добавлена команда /refresh
 - Добавлена возможность писать в телеграм-канал
 - Добавлен функционал запоминания клиентов, которые захотели учавстовать
 - Добавлен функционал уведомлений. Первое уведомление приходит за день до, а второе за 2 часа до начала
 - Добавлен интерфейс просмотра клиентов (в admin части). Так же есть возможность перейти в чат с клиентом сразу из бота


<b>Для тестирования 1.1 версии рекомендую:</b>

<b>1.</b> Зайти в канал t.me/testdvijhurghada
<b>2.</b> Воспользоваться Google Forms: https://docs.google.com/forms/d/e/1FAIpQLSfl-nw9Ik1Yl5tgdlGhgH5Sxa1C_Jtf6BMC3rUhcq4olZ361g/viewform?usp=sf_link
<b>3.</b> Отправить /refresh в личный чат с ботом (важно именно в личный!)
<b>4.</b> Нажать на кнопку "Я приду!" в телеграмм канале (кнопка будет в отправленном ботом сообщении)
<b>5.</b> Отправить /admin в личный чат и нажать на кнопку "Список", далее выбрать название активности, которое вы только заполняли в Google Forms и нажать на кнопку. Бот вышлет список всех желающих учавстовать в активности
<b>6.</b> Выберите клиента (нажав на кнопку). Бот пришлет небольшую информацию о клиенте и кнопку для перехода в личную переписку`
	dict["Client"] = "Клиент"
	dict["Admin"] = "Админ"
	dict["lookWhatWeHave"] = "Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас! На этой неделе у нас будут кое-какие активности, и мы будем вас с нетерпением ждать! Ниже нажмите на кнопку, чтобы узнать подробнее"
	dict["chooseOpt"] = "Выберите опцию"
	dict["MainMenu"] = "Главное Меню"
	dict["nameRu"] = "Введите название активности (на русском языке), которую вы добавляете. Это название будет отображаться на кнопках"
	dict["nameEn"] = "Введите название активности (на английском языке), которую вы добавляете. Это название будет отображаться на кнопках"
	dict["textRu"] = "Введите описание активности (на русском языке), которую добавляете. Это описание будет отображаться после нажатия на кнопку, которую вы только что создали"
	dict["textEn"] = "Введите описание активности (на английском языке), которую добавляете. Это описание будет отображаться после нажатия на кнопку, которую вы только что создали"
	dict["almostDone"] = `Вы практически создали новую активность. Осталось ее подтвердить

	Название (ru): %s
	
	Название (en): %s
	
	Описание (ru): %s
	
	Описание (en): %s
	
	Если все верно, то нажмите сохранить, если нет - измените, что пожелаете`
	dict["done"] = "Данные успешно изменены!"
	dict["save"] = "Сохранить"
	dict["change"] = "Изменить"
	dict["create"] = "Создать"
	dict["delete"] = "Удалить"
	dict["chooseAnActiv"] = "Выберите активность, которую хотите удалить"
	dict["cooseAnActivToChange"] = "Выберите активность, которую хотите изменить"
	dict["WhatyouwantChange"] = "Что вы хотите изменить?"
	dict["thisisdata"] = `На данный момент %s у этой активности такое:

	(ru): %s
	(en): %s
	
	Напишите новое название на %s`
	dict["almoseChdone"] = `Вы практически закончили изменения. Осталось все подтвердить

	Название (ru): %s
	
	Название (en): %s
	
	Описание (ru): %s
	
	Описание (en): %s
	
	Если все верно, то нажмите сохранить`
	dict["title"] = "Название"
	dict["discrp"] = "Описание"
	dict["Rus"] = "Русском языке"
	dict["Engl"] = "Английском языке"
	dict["SampleChannel"] = "Название: <b>%s</b>\n\n\nОписание <b>%s</b>\n\nДата и Время: <b>%s</b>\n\nСсылка на место проведения: %s"
	dict["Registred"] = "Вы записались на событие из телеграм канала testDvijHurghada. Мы уведомим вас за сутки и за несколько часов до начала. Вам нужно будет подтвердить, что вы точно придете"
	dict["1 day"] = "До начала активности на которую вы подписались отслось меньше суток! Вы все еще планируете поучавствовать?"
	dict["2 hours"] = "До начала активности на которую вы подписались отслось меньше двух часов! Вы все еще планируете поучавствовать?"
	dict["I will"] = "Да, я буду!"
	dict["I won't"] = "Нет, я не смогу!"
	dict["Cool1time"] = "Мы очень рады, что вы все еще хотите поучаствовать! Будем вас ждать! Мы пришлем вам еще одно уведомление за пару часов до начала"
	dict["Cool2time"] = "Мы очень рады, что вы все еще хотите поучаствовать! Будем вас ждать!"
	dict["upset"] = "Очень жаль!"
	dict["list"] = "Список"
	dict["ChAct"] = "Выберите активность"
	dict["ChClient"] = "Ниже представлены все ваши клиенты, которые ожидают эту активность. Для более подробной информации выберите одного"
	dict["ClientInf"] = "Данные клиента:\n\n\nИмя (если не скрыто): %s\nФамилия (если не скрыто): %s\nНомер телефона (если не скрыт): %s\nПодтверждений обнаружено: %d из 2"
	dict["Chat"] = "Личная переписка"
}

func en(dict map[string]string) {
	dict["sayHello"] = `<b>Bot Usage Instructions:</b>

<b>Main Commands:</b>
/admin - Allows viewing all functionalities available to the organizer
/client - Allows viewing all functionalities available to clients
/refresh - Upon receiving this command, the bot will start reading the table in Google Sheets. If it finds something new (not previously read), it will send the new data to the Telegram channel t.me/testdvijhurghada

<b>Update 1.0</b>
 - Added commands /admin and /client
 - All functionalities were developed and tested

<b>Update 1.1</b>
 - Added Google Forms for faster information entry
 - Implemented the connection between Google Forms and Google Sheets
 - Implemented the connection between the bot and Google Sheets
 - Added the /refresh command
 - Added the ability to post in the Telegram channel
 - Added functionality to remember clients who wish to participate
 - Added notification functionality. The first notification is sent a day before, and the second one two hours before the start
 - Added a client viewing interface (in the admin part). There is also the option to go directly to a chat with the client from the bot

<b>To test version 1.1, it is recommended to:</b>

<b>1.</b> Join the channel t.me/testdvijhurghada
<b>2.</b> Use the Google Forms: https://docs.google.com/forms/d/e/1FAIpQLSfl-nw9Ik1Yl5tgdlGhgH5Sxa1C_Jtf6BMC3rUhcq4olZ361g/viewform?usp=sf_link
<b>3.</b> Send /refresh in a private chat with the bot (it is important to do this in a private chat!)
<b>4.</b> Click the "Я приду!" button in the Telegram channel (the button will be in the message sent by the bot)
<b>5.</b> Send /admin in a private chat and click the "List" button. Then select the activity name that you just filled in Google Forms and click the button. The bot will send a list of everyone who wants to participate in the activity
<b>6.</b> Select a client (by clicking the button). The bot will send brief information about the client and a button to start a private conversation`
	dict["Client"] = "Client"
	dict["Admin"] = "Admin"
	dict["lookWhatWeHave"] = "Hello! This is the chatbot of Hurghada Church! We are happy to see you! This week, we will have some activities, and we are eagerly waiting for you! Click the button below to learn more"
	dict["chooseOpt"] = "Choose an option"
	dict["MainMenu"] = "Main Menu"
	dict["nameRu"] = "Enter the name of the activity (in Russian) that you are adding. This name will be displayed on the buttons"
	dict["nameEn"] = "Enter the name of the activity (in English) that you are adding. This name will be displayed on the buttons"
	dict["textRu"] = "Enter the description of the activity (in Russian) that you are adding. This description will be displayed after pressing the button you just created"
	dict["textEn"] = "Enter the description of the activity (in English) that you are adding. This description will be displayed after pressing the button you just created"
	dict["almostDone"] = `You have almost created a new activity. Just confirm it

	Title (ru): %s
	
	Title (en): %s
	
	Description (ru): %s
	
	Description (en): %s
	
	If everything is correct, click save. If not, make the desired changes`
	dict["done"] = "Data successfully updated!"
	dict["save"] = "Save"
	dict["change"] = "Change"
	dict["create"] = "Create"
	dict["delete"] = "Delete"
	dict["chooseAnActiv"] = "Select the activity you want to delete"
	dict["cooseAnActivToChange"] = "Select the activity you want to edit"
	dict["WhatyouwantChange"] = "What do you want to change?"
	dict["thisisdata"] = `Currently, the %s of this activity is:

	(ru): %s
	(en): %s
	
	Enter a new name in %s`
	dict["almoseChdone"] = `You have almost finished making changes. Just confirm everything

	Title (ru): %s
	
	Title (en): %s
	
	Description (ru): %s
	
	Description (en): %s
	
	If everything is correct, click save`
	dict["title"] = "Title"
	dict["discrp"] = "Description"
	dict["Rus"] = "Russian"
	dict["Engl"] = "English"
	dict["Registred"] = "You have registered for the event from the Telegram channel testDvijHurghada. We will notify you one day and a few hours before the start. You will need to confirm that you will definitely attend"
	dict["list"] = "List"
	dict["ChAct"] = "Select an activity"
	dict["ChClient"] = "Below are all your clients who are awaiting this activity. To get more details, select one"
	dict["ClientInf"] = "Client Details:\n\nName (if not hidden): %s\nSurname (if not hidden): %s\nPhone Number (if not hidden): %s\nConfirmations Found: %d out of 2"
	dict["Chat"] = "Private Messaging"
}

func init() {
	Dictionary = make(map[string]map[string]string)
	Dictionary["ru"] = make(map[string]string)
	Dictionary["en"] = make(map[string]string)

	ru(Dictionary["ru"])
	en(Dictionary["en"])
}
