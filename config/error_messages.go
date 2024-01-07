package config

type ErrorMessages struct {
  UNKNOW_DATABASE_CONNECTOR                  string
  FIELD_TITLE_SHOULD_NOT_BE_EMPTY            string
  FIELD_AUTHOR_AND_TITLE_SHOULD_NOT_BE_EMPTY string
  FIELD_AUTHOR_SHOULD_NOT_BE_EMPTY           string
}

func NewErrorMessages() *ErrorMessages {
  return &ErrorMessages{
    UNKNOW_DATABASE_CONNECTOR: "Unknow database connector!",
    FIELD_TITLE_SHOULD_NOT_BE_EMPTY: "Field title should not be empty!",
    FIELD_AUTHOR_AND_TITLE_SHOULD_NOT_BE_EMPTY: "Field author and title should not be empty!",
    FIELD_AUTHOR_SHOULD_NOT_BE_EMPTY: "Field author should not be empty!",
  }
}
