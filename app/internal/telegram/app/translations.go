package app

func translationsPrintIntro(instructionMessage, languageCode, welcomeMessage string) (string, string) {
	switch languageCode {
	case "ru":
		welcomeMessage = "Привет! " + EmojiSunglasses
		instructionMessage = "При взаимной симпатии ты получишь уведомление в чат этого бота." +
			" Нажми на кнопку Menu, чтобы начать пользоваться приложением"
	case "en":
		welcomeMessage = "Hello! " + EmojiSunglasses
		instructionMessage = "If you like each other, you will receive a notification in the chat of this bot." +
			" Click on the Menu button to start using the application"
	case "ar":
		welcomeMessage = " مرحبًا! " + EmojiSunglasses
		instructionMessage = "إذا كنت تحب بعضكما البعض، فسوف تتلقى إشعارًا في دردشة هذا الروبوت." +
			" اضغط على زر القائمة لبدء استخدام التطبيق "
	case "be":
		welcomeMessage = "Прывітанне! " + EmojiSunglasses
		instructionMessage = "Пры ўзаемнай сімпатыі ты атрымаеш апавяшчэнне ў чат гэтага робата." +
			" Націсніце на кнопку Menu, каб пачаць карыстацца дадаткам"
	case "ca":
		welcomeMessage = "Hola! " + EmojiSunglasses
		instructionMessage = "Si us agraden, rebreu una notificació al xat d'aquest bot." +
			" Premeu el botó Menú per començar a utilitzar l'aplicació"
	case "cs":
		welcomeMessage = "Ahoj! " + EmojiSunglasses
		instructionMessage = "Pokud se máte rádi, dostanete upozornění na chat tohoto robota." +
			" Stisknutím tlačítka Nabídka začněte používat aplikaci"
	case "de":
		welcomeMessage = "Hallo! " + EmojiSunglasses
		instructionMessage = "Wenn Sie sich mögen, erhalten Sie im Chat dieses Bots eine Benachrichtigung." +
			" Drücken Sie die Menütaste, um die Anwendung zu verwenden"
	case "es":
		welcomeMessage = "¡Hola! " + EmojiSunglasses
		instructionMessage = "Si se gustan, recibirán una notificación en el chat de este bot." +
			" Presione el botón Menú para comenzar a usar la aplicación"
	case "fi":
		welcomeMessage = "Hei! " + EmojiSunglasses
		instructionMessage = "Jos pidätte toisistanne, saat ilmoituksen tämän botin chatissa." +
			" Aloita sovelluksen käyttö painamalla Menu-painiketta"
	case "fr":
		welcomeMessage = "Bonjour! " + EmojiSunglasses
		instructionMessage = "Si vous vous aimez, vous recevrez une notification dans le chat de ce bot." +
			" Appuyez sur le bouton Menu pour commencer à utiliser l'application"
	case "he":
		welcomeMessage = "שלום! " + EmojiSunglasses
		instructionMessage = "אם אתם אוהבים אחד את השני, תקבלו התראה בצ'אט של הבוט הזה." +
			" לחץ על לחצן התפריט כדי להתחיל להשתמש באפליקציה "
	case "hr":
		welcomeMessage = "Zdravo! " + EmojiSunglasses
		instructionMessage = "Ako se sviđate jedno drugome, dobit ćete obavijest u chatu ovog bota." +
			" Pritisnite tipku Izbornik za početak korištenja aplikacije"
	case "hu":
		welcomeMessage = "Helló! " + EmojiSunglasses
		instructionMessage = "Ha kedvelitek egymást, értesítést fog kapni ennek a botnak a chatjében." +
			" Az alkalmazás használatának megkezdéséhez nyomja meg a Menü gombot"
	case "id":
		welcomeMessage = "Halo! " + EmojiSunglasses
		instructionMessage = "Jika Anda menyukai satu sama lain, Anda akan menerima pemberitahuan di obrolan bot ini." +
			" Tekan tombol Menu untuk mulai menggunakan aplikasi"
	case "it":
		welcomeMessage = "Ciao! " + EmojiSunglasses
		instructionMessage = "Se vi piacciono, riceverete una notifica nella chat di questo bot." +
			" Premere il pulsante Menu per iniziare a utilizzare l'applicazione"
	case "ja":
		welcomeMessage = "こんにちは！ " + EmojiSunglasses
		instructionMessage = "お互いにいいねすると、このボットのチャットに通知が届きます。" +
			" メニューボタンを押してアプリケーションの使用を開始します"
	case "kk":
		welcomeMessage = "Сәлем! " + EmojiSunglasses
		instructionMessage = "Егер сіз бір-біріңізді ұнатсаңыз, сіз осы боттың чатында хабарландыру аласыз." +
			" Қолданбаны пайдалануды бастау үшін Мәзір түймесін басыңыз"
	case "ko":
		welcomeMessage = "안녕하세요! " + EmojiSunglasses
		instructionMessage = "서로 좋아요를 누르면 이 봇 채팅을 통해 알림을 받게 됩니다." +
			" 메뉴 버튼을 눌러 애플리케이션 사용을 시작하세요."
	case "nl":
		welcomeMessage = "Hallo! " + EmojiSunglasses
		instructionMessage = "Als je elkaar leuk vindt, ontvang je een melding in de chat van deze bot." +
			" Druk op de Menu-knop om de applicatie te gebruiken"
	case "no":
		welcomeMessage = "Hallo! " + EmojiSunglasses
		instructionMessage = "Hvis du liker hverandre, vil du motta et varsel i chatten til denne boten." +
			" Trykk på Meny-knappen for å begynne å bruke applikasjonen"
	case "pt":
		welcomeMessage = "Olá! " + EmojiSunglasses
		instructionMessage = "Se vocês gostarem, receberão uma notificação no chat deste bot." +
			" Pressione o botão Menu para começar a usar o aplicativo"
	case "sv":
		welcomeMessage = "Hej! " + EmojiSunglasses
		instructionMessage = "Om ni gillar varandra kommer ni att få ett meddelande i chatten för denna bot." +
			" Tryck på menyknappen för att börja använda programmet"
	case "uk":
		welcomeMessage = "Привіт! " + EmojiSunglasses
		instructionMessage = "При взаємній симпатії ти отримаєш повідомлення у чат цього робота." +
			" Натисніть на кнопку Menu, щоб почати користуватися програмою"
	case "zh":
		welcomeMessage = "你好！ " + EmojiSunglasses
		instructionMessage = "如果你们互相喜欢，您将在此机器人的聊天中收到通知。" +
			" 按菜单按钮开始使用该应用程序"
	default:
		welcomeMessage = "Hello! " + EmojiSunglasses
		instructionMessage = "Click the Menu button to start using the application"
	}
	return instructionMessage, welcomeMessage
}
