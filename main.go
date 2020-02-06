package main

func main() {
	a := App{}
	// konfigurasi database disini
	a.Initialize("root", "", "db_go")

	a.Run(":8080")
}
