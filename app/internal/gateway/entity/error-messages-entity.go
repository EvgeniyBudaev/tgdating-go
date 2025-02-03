package entity

import "fmt"

type ErrorMessagesEntity struct {
}

func NewErrorMessagesEntity() *ErrorMessagesEntity {
	return &ErrorMessagesEntity{}
}

func (e *ErrorMessagesEntity) GetBadRequest(locale string) string {
	switch locale {
	case "ru":
		return "Некорректные данные в запросе"
	case "en":
		return "Incorrect data in the request"
	case "ar":
		return "البيانات غير صحيحة في الطلب"
	case "be":
		return "Некарэктныя дадзеныя ў запыце"
	case "ca":
		return "Dades incorrectes a la sol·licitud"
	case "cs":
		return "Nesprávné údaje v požadavku"
	case "de":
		return "Falsche daten in der anfrage"
	case "es":
		return "Datos incorrectos en la solicitud"
	case "fi":
		return "Pyynnössä virheelliset tiedot"
	case "fr":
		return "Données incorrectes dans la demande"
	case "he":
		return "נתונים שגויים בבקשה"
	case "hi":
		return "अनुरोध में ग़लत डेटा"
	case "hr":
		return "Netočni podaci u zahtjevu"
	case "hu":
		return "Hibás adatok a kérésben"
	case "id":
		return "Data yang salah dalam permintaan"
	case "it":
		return "Dati errati nella richiesta"
	case "ja":
		return "リクエスト内のデータが正しくありません"
	case "kk":
		return "Сұраудағы деректер дұрыс емес"
	case "ko":
		return "요청에 잘못된 데이터가 있습니다."
	case "nl":
		return "Onjuiste gegevens in de aanvraag"
	case "no":
		return "Feil data i forespørselen"
	case "pt":
		return "Dados incorretos na solicitação"
	case "sv":
		return "Felaktiga uppgifter i begäran"
	case "uk":
		return "Некоректні дані у запиті"
	case "zh":
		return "请求中的数据不正确"
	default:
		return "Incorrect data in the request"
	}
}

func (e *ErrorMessagesEntity) GetLessOrEqualMaxNumber(locale string, max int) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Должно быть меньше или равно %d", max)
	case "en":
		return fmt.Sprintf("Must be less or equal to %d", max)
	case "ar":
		return fmt.Sprintf("يجب أن يكون أقل من أو يساوي %d", max)
	case "be":
		return fmt.Sprintf("Павінна быць менш або роўна %d", max)
	case "ca":
		return fmt.Sprintf("Ha de ser inferior o igual a %d", max)
	case "cs":
		return fmt.Sprintf("Musí být menší nebo roven %d", max)
	case "de":
		return fmt.Sprintf("Muss kleiner oder gleich %d sein", max)
	case "es":
		return fmt.Sprintf("Debe ser menor o igual a %d", max)
	case "fi":
		return fmt.Sprintf("On oltava pienempi tai yhtä suuri kuin %d", max)
	case "fr":
		return fmt.Sprintf("Doit être inférieur ou égal à %d", max)
	case "he":
		return fmt.Sprintf("חייב להיות קטן או שווה ל- %d", max)
	case "hi":
		return fmt.Sprintf("%d से कम या उसके बराबर होना चाहिए", max)
	case "hr":
		return fmt.Sprintf("Mora biti manje od ili jednako %d", max)
	case "hu":
		return fmt.Sprintf("Kisebbnek vagy egyenlőnek kell lennie, mint %d", max)
	case "id":
		return fmt.Sprintf("Harus kurang dari atau sama dengan %d", max)
	case "it":
		return fmt.Sprintf("Deve essere inferiore o uguale a %d", max)
	case "ja":
		return fmt.Sprintf("%d 以下である必要があります", max)
	case "kk":
		return fmt.Sprintf("%d мәнінен кем немесе оған тең болуы керек", max)
	case "ko":
		return fmt.Sprintf("%d 보다 작거나 같아야 합니다.", max)
	case "nl":
		return fmt.Sprintf("Moet kleiner dan of gelijk zijn aan %d", max)
	case "no":
		return fmt.Sprintf("Må være mindre enn eller lik %d", max)
	case "pt":
		return fmt.Sprintf("Deve ser menor ou igual a %d", max)
	case "sv":
		return fmt.Sprintf("Måste vara mindre än eller lika med %d", max)
	case "uk":
		return fmt.Sprintf("Має бути менше або дорівнює %d", max)
	case "zh":
		return fmt.Sprintf("必须小于或等于 %d", max)
	default:
		return fmt.Sprintf("Must be less or equal to %d", max)
	}
}

func (e *ErrorMessagesEntity) GetMoreOrEqualMinUint64Number(locale string, min uint64) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Должно быть больше или равно %d", min)
	case "en":
		return fmt.Sprintf("Must be more or equal to %d", min)
	case "ar":
		return fmt.Sprintf("يجب أن يكون أكبر من أو يساوي %d", min)
	case "be":
		return fmt.Sprintf("Павінна быць больш або роўна %d", min)
	case "ca":
		return fmt.Sprintf("Ha de ser superior o igual a %d", min)
	case "cs":
		return fmt.Sprintf("Musí být větší nebo rovno %d", min)
	case "de":
		return fmt.Sprintf("Muss größer oder gleich %d sein", min)
	case "es":
		return fmt.Sprintf("Debe ser mayor o igual a %d", min)
	case "fi":
		return fmt.Sprintf("Sen on oltava suurempi tai yhtä suuri kuin %d", min)
	case "fr":
		return fmt.Sprintf("Doit être supérieur ou égal à %d", min)
	case "he":
		return fmt.Sprintf("חייב להיות גדול או שווה ל- %d", min)
	case "hi":
		return fmt.Sprintf("%d से अधिक या उसके बराबर होना चाहिए", min)
	case "hr":
		return fmt.Sprintf("Mora biti veći od ili jednak %d", min)
	case "hu":
		return fmt.Sprintf("Nagyobbnak vagy egyenlőnek kell lennie %d", min)
	case "id":
		return fmt.Sprintf("Harus lebih besar atau sama dengan %d", min)
	case "it":
		return fmt.Sprintf("Deve essere maggiore o uguale a %d", min)
	case "ja":
		return fmt.Sprintf("%d 以上である必要があります", min)
	case "kk":
		return fmt.Sprintf("%d мәнінен үлкен немесе тең болуы керек", min)
	case "ko":
		return fmt.Sprintf("%d 보다 크거나 같아야 합니다.", min)
	case "nl":
		return fmt.Sprintf("Moet groter dan of gelijk zijn aan %d", min)
	case "no":
		return fmt.Sprintf("Må være større enn eller lik %d", min)
	case "pt":
		return fmt.Sprintf("Deve ser maior ou igual a %d", min)
	case "sv":
		return fmt.Sprintf("Måste vara större än eller lika med %d", min)
	case "uk":
		return fmt.Sprintf("Має бути більше або дорівнює %d", min)
	case "zh":
		return fmt.Sprintf("必须大于或等于 %d", min)
	default:
		return fmt.Sprintf("Must be more or equal to %d", min)
	}
}

func (e *ErrorMessagesEntity) GetLessOrEqualMaxUint64Number(locale string, max uint64) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Должно быть меньше или равно %d", max)
	case "en":
		return fmt.Sprintf("Must be less or equal to %d", max)
	case "ar":
		return fmt.Sprintf("يجب أن يكون أقل من أو يساوي %d", max)
	case "be":
		return fmt.Sprintf("Павінна быць менш або роўна %d", max)
	case "ca":
		return fmt.Sprintf("Ha de ser inferior o igual a %d", max)
	case "cs":
		return fmt.Sprintf("Musí být menší nebo roven %d", max)
	case "de":
		return fmt.Sprintf("Muss kleiner oder gleich %d sein", max)
	case "es":
		return fmt.Sprintf("Debe ser menor o igual a %d", max)
	case "fi":
		return fmt.Sprintf("On oltava pienempi tai yhtä suuri kuin %d", max)
	case "fr":
		return fmt.Sprintf("Doit être inférieur ou égal à %d", max)
	case "he":
		return fmt.Sprintf("חייב להיות קטן או שווה ל- %d", max)
	case "hi":
		return fmt.Sprintf("%d से कम या उसके बराबर होना चाहिए", max)
	case "hr":
		return fmt.Sprintf("Mora biti manje od ili jednako %d", max)
	case "hu":
		return fmt.Sprintf("Kisebbnek vagy egyenlőnek kell lennie, mint %d", max)
	case "id":
		return fmt.Sprintf("Harus kurang dari atau sama dengan %d", max)
	case "it":
		return fmt.Sprintf("Deve essere inferiore o uguale a %d", max)
	case "ja":
		return fmt.Sprintf("%d 以下である必要があります", max)
	case "kk":
		return fmt.Sprintf("%d мәнінен кем немесе оған тең болуы керек", max)
	case "ko":
		return fmt.Sprintf("%d 보다 작거나 같아야 합니다.", max)
	case "nl":
		return fmt.Sprintf("Moet kleiner zijn dan of gelijk zijn aan %d", max)
	case "no":
		return fmt.Sprintf("Må være mindre enn eller lik %d", max)
	case "pt":
		return fmt.Sprintf("Deve ser menor ou igual a %d", max)
	case "sv":
		return fmt.Sprintf("Måste vara mindre än eller lika med %d", max)
	case "uk":
		return fmt.Sprintf("Має бути менше або дорівнює %d", max)
	case "zh":
		return fmt.Sprintf("必须小于或等于 %d", max)
	default:
		return fmt.Sprintf("Must be less or equal to %d", max)
	}
}

func (e *ErrorMessagesEntity) GetMaxSymbols(locale string, max int) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Должно быть не более %d символов", max)
	case "en":
		return fmt.Sprintf("Must be no more than %d characters", max)
	case "ar":
		return fmt.Sprintf("يجب ألا يزيد عدد الأحرف عن %d", max)
	case "be":
		return fmt.Sprintf("Павінна быць не больш за %d сімвалаў", max)
	case "ca":
		return fmt.Sprintf("No ha de tenir més de %d caràcters", max)
	case "cs":
		return fmt.Sprintf("Nesmí obsahovat více než %d znaků", max)
	case "de":
		return fmt.Sprintf("Darf nicht mehr als %d Zeichen umfassen", max)
	case "es":
		return fmt.Sprintf("No debe tener más de %d caracteres", max)
	case "fi":
		return fmt.Sprintf("Saa olla enintään %d merkkiä", max)
	case "fr":
		return fmt.Sprintf("Ne doit pas contenir plus de %d caractères", max)
	case "he":
		return fmt.Sprintf("חייב להיות לא יותר מ- %d תווים", max)
	case "hi":
		return fmt.Sprintf("%d वर्णों से अधिक नहीं होना चाहिए", max)
	case "hr":
		return fmt.Sprintf("Ne smije imati više od %d znakova", max)
	case "hu":
		return fmt.Sprintf("Nem lehet több %d karakternél", max)
	case "id":
		return fmt.Sprintf("Tidak boleh lebih dari %d karakter", max)
	case "it":
		return fmt.Sprintf("Non deve contenere più di %d caratteri", max)
	case "ja":
		return fmt.Sprintf("%d 文字以下にする必要があります", max)
	case "kk":
		return fmt.Sprintf("%d таңбадан аспауы керек", max)
	case "ko":
		return fmt.Sprintf("%d 자 이하여야 합니다.", max)
	case "nl":
		return fmt.Sprintf("Mag niet langer zijn dan %d tekens", max)
	case "no":
		return fmt.Sprintf("Må ikke være mer enn %d tegn", max)
	case "pt":
		return fmt.Sprintf("Não deve ter mais que %d caracteres", max)
	case "sv":
		return fmt.Sprintf("Får inte vara fler än %d tecken", max)
	case "uk":
		return fmt.Sprintf("Має бути не більше %d символів", max)
	case "zh":
		return fmt.Sprintf("不得超过 %d 个字符", max)
	default:
		return fmt.Sprintf("Must be no more than %d characters", max)
	}
}

func (e *ErrorMessagesEntity) GetFileMaxAmount(locale string, max int) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Максимальное кол-во файлов %d", max)
	case "en":
		return fmt.Sprintf("Maximum number of files %d", max)
	case "ar":
		return fmt.Sprintf("الحد الأقصى لعدد الملفات %d", max)
	case "be":
		return fmt.Sprintf("Максімальная колькасць файлаў %d", max)
	case "ca":
		return fmt.Sprintf("Nombre màxim de fitxers %d", max)
	case "cs":
		return fmt.Sprintf("Maximální počet souborů %d", max)
	case "de":
		return fmt.Sprintf("Maximale Anzahl Dateien %d", max)
	case "es":
		return fmt.Sprintf("Número máximo de archivos %d", max)
	case "fi":
		return fmt.Sprintf("Tiedostojen enimmäismäärä %d", max)
	case "fr":
		return fmt.Sprintf("Nombre maximum de fichiers %d", max)
	case "he":
		return fmt.Sprintf("מספר מקסימלי של קבצים %d", max)
	case "hi":
		return fmt.Sprintf("फ़ाइलों की अधिकतम संख्या %d", max)
	case "hr":
		return fmt.Sprintf("Maksimalan broj datoteka %d", max)
	case "hu":
		return fmt.Sprintf("Fájlok maximális száma %d", max)
	case "id":
		return fmt.Sprintf("Jumlah maksimum file %d", max)
	case "it":
		return fmt.Sprintf("Numero massimo di file %d", max)
	case "ja":
		return fmt.Sprintf("ファイルの最大数 %d", max)
	case "kk":
		return fmt.Sprintf("Файлдардың ең көп саны %d", max)
	case "ko":
		return fmt.Sprintf("최대 파일 수 %d", max)
	case "nl":
		return fmt.Sprintf("Maximaal aantal bestanden %d", max)
	case "no":
		return fmt.Sprintf("Maksimalt antall filer %d", max)
	case "pt":
		return fmt.Sprintf("Número máximo de arquivos %d", max)
	case "sv":
		return fmt.Sprintf("Maximalt antal filer %d", max)
	case "uk":
		return fmt.Sprintf("Максимальна кількість файлів %d", max)
	case "zh":
		return fmt.Sprintf("最大文件数 %d", max)
	default:
		return fmt.Sprintf("Maximum number of files %d", max)
	}
}

func (e *ErrorMessagesEntity) GetFileMaxSize(locale string, max int) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Максимальный размер файла %dMb", max)
	case "en":
		return fmt.Sprintf("Maximum size file %dMb", max)
	case "ar":
		return fmt.Sprintf("الحد الأقصى لحجم الملف %dMB", max)
	case "be":
		return fmt.Sprintf("Максімальны памер файла %dMb", max)
	case "ca":
		return fmt.Sprintf("Mida màxima del fitxer %dMb", max)
	case "cs":
		return fmt.Sprintf("Maximální velikost souboru %dMb", max)
	case "de":
		return fmt.Sprintf("Maximale Dateigröße %dMb", max)
	case "es":
		return fmt.Sprintf("Tamaño máximo de archivo %dMb", max)
	case "fi":
		return fmt.Sprintf("Tiedoston enimmäiskoko %dMb", max)
	case "fr":
		return fmt.Sprintf("Taille maximale du fichier %dMb", max)
	case "he":
		return fmt.Sprintf("גודל קובץ מקסימלי %dMB", max)
	case "hi":
		return fmt.Sprintf("अधिकतम फ़ाइल आकार %dMb", max)
	case "hr":
		return fmt.Sprintf("Maksimalna veličina datoteke %dMb", max)
	case "hu":
		return fmt.Sprintf("Maximális fájlméret %dMb", max)
	case "id":
		return fmt.Sprintf("Ukuran file maksimum %dMb", max)
	case "it":
		return fmt.Sprintf("Dimensione massima del file %dMb", max)
	case "ja":
		return fmt.Sprintf("最大ファイルサイズ %dMb", max)
	case "kk":
		return fmt.Sprintf("Ең үлкен файл өлшемі %dMb", max)
	case "ko":
		return fmt.Sprintf("최대 파일 크기 %dMb", max)
	case "nl":
		return fmt.Sprintf("Maximale bestandsgrootte %dMb", max)
	case "no":
		return fmt.Sprintf("Maksimal filstørrelse %dMb", max)
	case "pt":
		return fmt.Sprintf("Tamanho máximo do arquivo %dMb", max)
	case "sv":
		return fmt.Sprintf("Maximal filstorlek %dMb", max)
	case "uk":
		return fmt.Sprintf("Максимальний розмір файлу %dMb", max)
	case "zh":
		return fmt.Sprintf("最大文件大小 %dMb", max)
	default:
		return fmt.Sprintf("Maximum size file %dMb", max)
	}
}

func (e *ErrorMessagesEntity) GetNotEmpty(locale string) string {
	switch locale {
	case "ru":
		return "Поле не может быть пустым"
	case "en":
		return "Field cannot be empty"
	case "ar":
		return "لا يمكن أن يكون الحقل فارغاً"
	case "be":
		return "Поле не можа быць пустым"
	case "ca":
		return "El camp no pot estar buit"
	case "cs":
		return "Pole nemůže být prázdné"
	case "de":
		return "Das feld darf nicht leer sein"
	case "es":
		return "El campo no puede estar vacío"
	case "fi":
		return "Kenttä ei voi olla tyhjä"
	case "fr":
		return "Le champ ne peut pas être vide"
	case "he":
		return "שדה לא יכול להיות ריק"
	case "hi":
		return "फ़ील्ड खाली नहीं हो सकती"
	case "hr":
		return "Polje ne može biti prazno"
	case "hu":
		return "A mező nem lehet üres"
	case "id":
		return "Bidang tidak boleh kosong"
	case "it":
		return "Il campo non può essere vuoto"
	case "ja":
		return "フィールドを空にすることはできません"
	case "kk":
		return "Өріс бос болмауы керек"
	case "ko":
		return "필드는 비워둘 수 없습니다"
	case "nl":
		return "Veld mag niet leeg zijn"
	case "no":
		return "Feltet kan ikke være tomt"
	case "pt":
		return "O campo não pode ficar vazio"
	case "sv":
		return "Fältet får inte vara tomt"
	case "uk":
		return "Поле не може бути порожнім"
	case "zh":
		return "字段不能为空"
	default:
		return "Field cannot be empty"
	}
}

func (e *ErrorMessagesEntity) GetNonNegativeNumber(locale string) string {
	switch locale {
	case "ru":
		return "Число должно быть положительным"
	case "en":
		return "Must be a positive number"
	case "ar":
		return "يجب أن يكون الرقم موجبًا"
	case "be":
		return "Лік павінен быць станоўчым"
	case "ca":
		return "El nombre ha de ser positiu"
	case "cs":
		return "Číslo musí být kladné"
	case "de":
		return "Die Zahl muss positiv sein"
	case "es":
		return "El numero debe ser positivo"
	case "fi":
		return "Numeron on oltava positiivinen"
	case "fr":
		return "Le nombre doit être positif"
	case "he":
		return "המספר חייב להיות חיובי"
	case "hi":
		return "संख्या धनात्मक होनी चाहिए"
	case "hr":
		return "Broj mora biti pozitivan"
	case "hu":
		return "A számnak pozitívnak kell lennie"
	case "id":
		return "Angkanya harus positif"
	case "it":
		return "Il numero deve essere positivo"
	case "ja":
		return "数値は正でなければなりません"
	case "kk":
		return "Сан оң болуы керек"
	case "ko":
		return "숫자는 양수여야 합니다"
	case "nl":
		return "Het getal moet positief zijn"
	case "no":
		return "Tallet må være positivt"
	case "pt":
		return "O número deve ser positivo"
	case "sv":
		return "Siffran måste vara positiv"
	case "uk":
		return "Число має бути позитивним"
	case "zh":
		return "数字必须为正数"
	default:
		return "Must be a positive number"
	}
}
