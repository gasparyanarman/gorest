package app

import "Rest_Api/controllers"

func mapUrls() {
	router.GET("/", controllers.GetHomePage)
}
