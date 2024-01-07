package config

type constants struct {
  POSTGRES string
}

func NewConstants () *constants {
  return &constants{
    POSTGRES: "postgres",
  }
}
