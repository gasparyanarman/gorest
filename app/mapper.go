package app

import "Rest_Api/controllers"

func mapUrls() {
	router.POST("/authors", controllers.SaveAuthor)
}
