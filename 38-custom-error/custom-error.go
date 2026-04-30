package main

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

type notFoundError struct {
	Message string
}

func (e *notFoundError) Error() string {
	return e.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "ID tidak boleh kosong"}
	}
	if id != "abdul" {
		return &notFoundError{Message: "ID tidak ditemukan"}
	}
	if data == nil {
		return &validationError{Message: "Data tidak boleh nil"}
	}
	// Simulasi penyimpanan data
	return nil
}

func main() {
	err := SaveData("abdul", nil)
	if err != nil {
		// pakai switch type assertion
		// switch e := err.(type) {
		// case *validationError:
		// 	println("Validation Error:", e.Error())
		// case *notFoundError:
		// 	println("Not Found Error:", e.Error())
		// default:
		// 	println("Unknown Error:", e.Error())
		// }
		// return
		// pakai if type assertion
		if validationError, ok := err.(*validationError); ok {
			println("Validation Error:", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			println("Not Found Error:", notFoundError.Error())
		} else {
			println("Unknown Error:", err.Error())
		}
		return
	}
	println("Data berhasil disimpan")
}
