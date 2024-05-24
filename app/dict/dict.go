package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["sayHello"] = "Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас!"
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
}

func en(dict map[string]string) {
	dict["sayHello"] = "Hello! This is the chatbot of Hurghada Church! We are happy to see you!"
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
}

func init() {
	Dictionary = make(map[string]map[string]string)
	Dictionary["ru"] = make(map[string]string)
	Dictionary["en"] = make(map[string]string)

	ru(Dictionary["ru"])
	en(Dictionary["en"])
}
