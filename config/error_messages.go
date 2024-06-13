package config

type ErrorMessages struct {
  UNKNOW_DATABASE_CONNECTOR                  string
  FIELD_TITLE_SHOULD_NOT_BE_EMPTY            string
  FIELD_AUTHOR_AND_TITLE_SHOULD_NOT_BE_EMPTY string
  FIELD_AUTHOR_SHOULD_NOT_BE_EMPTY           string
  USER_REQUIRED                              string
  USER_PASSWORD_REQUIRED                     string
  USER_EMAIL_REQUIRED                        string
  USER_NOT_FOUND                             string
  USER_EMAIL_OR_PASSWORD_REQUIRED            string
}

func NewErrorMessages() *ErrorMessages {
  return &ErrorMessages{
    UNKNOW_DATABASE_CONNECTOR: "Unknow database connector!",
    FIELD_TITLE_SHOULD_NOT_BE_EMPTY: "Field title should not be empty!",
    FIELD_AUTHOR_AND_TITLE_SHOULD_NOT_BE_EMPTY: "Field author and title should not be empty!",
    FIELD_AUTHOR_SHOULD_NOT_BE_EMPTY: "Field author should not be empty!",
    USER_REQUIRED: "Username is required!",
    USER_PASSWORD_REQUIRED: "Password is required!",
    USER_EMAIL_REQUIRED: "Email is required!",
    USER_NOT_FOUND: "User not found!",
    USER_EMAIL_OR_PASSWORD_REQUIRED: "Email or Password is required!",
  }
}
