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
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}

func (e *ErrorMessagesEntity) GetLessOrEqualMaxNumber(locale string, max uint64) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Должно быть меньше или равно %d", max)
	case "en":
		return fmt.Sprintf("Must be less or equal to %d", max)
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}

func (e *ErrorMessagesEntity) GetMaxSymbols(locale string, max uint64) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Должно быть не более %d символов", max)
	case "en":
		return fmt.Sprintf("Must be no more than %d characters", max)
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}

func (e *ErrorMessagesEntity) GetFileMaxAmount(locale string, max int) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Максимальное кол-во файлов %d", max)
	case "en":
		return fmt.Sprintf("Maximum number of files %d", max)
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}

func (e *ErrorMessagesEntity) GetFileMaxSize(locale string, max int) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Максимальный размер файла %dМб", max)
	case "en":
		return fmt.Sprintf("Maximum size file %dMb", max)
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}

func (e *ErrorMessagesEntity) GetMoreOrEqualMinNumber(locale string, min uint64) string {
	switch locale {
	case "ru":
		return fmt.Sprintf("Должно быть больше или равно %d", min)
	case "en":
		return fmt.Sprintf("Must be more or equal to %d", min)
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}

func (e *ErrorMessagesEntity) GetNotEmpty(locale string) string {
	switch locale {
	case "ru":
		return "Поле не может быть пустым"
	case "en":
		return "Field cannot be empty"
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}

func (e *ErrorMessagesEntity) GetNonNegativeNumber(locale string) string {
	switch locale {
	case "ru":
		return "Число должно быть положительным"
	case "en":
		return "Must be a positive number"
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}
