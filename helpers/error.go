package helpers

func Error(err error) {
  if err != nil {
    panic(err)
  }
}
