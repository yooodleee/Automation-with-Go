package advanced

import (
	"advanced/config"
	"advanced/model"
)


func main() {
	config.Init()
	config.DB.AutoMigrate(&model.User{})
}