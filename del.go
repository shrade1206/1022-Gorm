package main

func del() {
	var pic []Pic
	Main_DB.Where("Name = ? ", "test").Delete(&pic)

	Main_DB.Unscoped().Where("Name = ?", "Test").Delete(&pic)

}
