package web

//RouterInit Add by Eric Shi
func RouterInit() {

	APIRouter := E.Group("/api")
	APIRouter.Use()

	WebRouter := E.Group("/")
	WebRouter.Use()

	index := Index{}
	WebRouter.GET("", index.MainPage)
}
