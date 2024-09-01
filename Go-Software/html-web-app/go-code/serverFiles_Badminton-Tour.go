# .. Badminton Tour ° // - Website App ~

package main



import (


)


func app_welcome_center_page() {


}

func app_infornation_booth_page() {


}

func app_login_center_page() {


}


func app_about_app_page() {


}

func app_dashboard_interface_page() {



}

func app_player_profile_page() {


}

func app_contact_list_page() {


}

func app_team_chat_page() {


}

func app_date_planner_page() {


}

func app_location_map_page() {


}

func app_schedule_calendar_page() {


}

func app_course_map_page() {


}

func app_score_card_page() {


}

func app_match_history_page() {


}


// . indexHandler,  ~ for Public Pages °
func indexHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "Badminton Tour ° // - Website App ~"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  if pagePath == "/" {
      pageTitle = "Index Page"
      pageList = pageList
  }
  
    if pagePath == "/front" {
      pageTitle = "Front Page"
      pageList = pageList
  }
  
    if pagePath == "/main" {
      pageTitle = "Main Page"
      pageList = pageList
  }
  
    if pagePath == "/home" {
      pageTitle = "Home Page"
      pageList = pageList
  }
  
    if pagePath == "/start" {
      pageTitle = "Start Page"
      pageList = pageList
  }
  
  
  
  // ~ http.HandleFunc("/what", indexHandler)
// - What - // Match Type °

    if pagePath == "/what" {
      pageTitle = "What Page"
      pageList = pageList
  }

// ~ http.HandleFunc("/when", indexHandler)
// - When - // Tee Time ° 

    if pagePath == "/when" {
      pageTitle = "When Page"
      pageList = pageList
  }
  
// ~ http.HandleFunc("/where", indexHandler)
// - Where - // Course Location °

    if pagePath == "/where" {
      pageTitle = "Where Page"
      pageList = pageList
  }
  
// ~ http.HandleFunc("/who", indexHandler)
// - Who - // Player Buddy Contact List °

    if pagePath == "/who" {
      pageTitle = "Who Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/why", indexHandler)
// - Why - // Reason for Match °

    if pagePath == "/why" {
      pageTitle = "Why Page"
      pageList = pageList
  }
  
// ~ http.HandleFunc("/how", indexHandler)
// - How - // Price and Payment Info °

    if pagePath == "/how" {
      pageTitle = "How Page"
      pageList = pageList
  }


// ,  ° . +
### _ non-user

// ~ http.HandleFunc("/welcome_screen", indexHandler)
// + Welcome Screen

    if pagePath == "/welcome_screen" {
      pageTitle = "Welcome Screen Page"
      pageList = pageList
  }
  
// ~ http.HandleFunc("/information_booth", indexHandler)
// + Information Booth 

    if pagePath == "/information_booth" {
      pageTitle = "Information Booth Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/login_center", indexHandler)
// + Login Center

    if pagePath == "/login_center" {
      pageTitle = "Login Center Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/about_app", indexHandler)
// + About App

    if pagePath == "/about_app" {
      pageTitle = "About App Page"
      pageList = pageList
  }

// ,  ° . +
### _ setting match

// ~ http.HandleFunc("/dashboard_interface", indexHandler)
// + Dashboard Interface 

    if pagePath == "/dashboard_interface" {
      pageTitle = "Dashboard Interface Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/player_profile", indexHandler)
// + Player Profile 

    if pagePath == "/player_profile" {
      pageTitle = "Player Profile Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/contact_list", indexHandler)
// + Contact List

    if pagePath == "/contact_list" {
      pageTitle = "Contact List Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/team_chat", indexHandler)
// + Team Chat

    if pagePath == "/team_chat" {
      pageTitle = "Team Chat Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/date_planner", indexHandler)
// + Date Planner

    if pagePath == "/date_planner" {
      pageTitle = "Date Planner Page"
      pageList = pageList
  }

// ,  ° . +
### _ tee time

// ~ http.HandleFunc("/location_map", indexHandler)
// + Location Map

    if pagePath == "/location_map" {
      pageTitle = "Location Map Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/schedule_calendar", indexHandler)
// + Schedule Calendar 

    if pagePath == "/schedule_calendar" {
      pageTitle = "Schedule Calendar Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/course_map", indexHandler)
// + Course Map

    if pagePath == "/course_map" {
      pageTitle = "Course Map Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/score_card", indexHandler)
// + Score Card

    if pagePath == "/score_card" {
      pageTitle = "Score Card Page"
      pageList = pageList
  }
  
  
// ~ http.HandleFunc("/match_history", indexHandler)
// + Match History

    if pagePath == "/match_history" {
      pageTitle = "Match History Page"
      pageList = pageList
  }



// ,  ° . +
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler


//  .  html url routes 
//  .  as well as terminal cli logs

func main() {
// ,  ° . +
  appName := "Badminton Tour ° // - Website App ~"

// ,  ° . +
  http.HandleFunc("/", indexHandler)
  
  // ,  ° . +
  http.HandleFunc("/user", indexHandler)
  http.HandleFunc("/account", indexHandler)
  http.HandleFunc("/profile", indexHandler)
  
  // ,  ° . +
  http.HandleFunc("/portfolio", indexHandler)
  http.HandleFunc("/resume", indexHandler)
  
  // ,  ° . +
  http.HandleFunc("/settings", indexHandler)
  
  
  // ,  ° . +
http.HandleFunc("/what", indexHandler)
- What - // Match Type °

http.HandleFunc("/when", indexHandler)
// - When - // Tee Time ° 
http.HandleFunc("/where", indexHandler)
// - Where - // Course Location °
http.HandleFunc("/who", indexHandler)
// - Who - // Player Buddy Contact List °
http.HandleFunc("/why", indexHandler)
// - Why - // Reason for Match °
http.HandleFunc("/how", indexHandler)
// - How - // Price and Payment Info °



// ,  ° . +
### _ non-user

http.HandleFunc("/welcome_screen", indexHandler)
// + Welcome Screen
http.HandleFunc("/information_booth", indexHandler)
// + Information Booth 
http.HandleFunc("/login_center", indexHandler)
// + Login Center
http.HandleFunc("/about_app", indexHandler)
// + About App

// ,  ° . +
### _ setting match

http.HandleFunc("/dashboard_interface", indexHandler)
// + Dashboard Interface 
http.HandleFunc("/player_profile", indexHandler)
// + Player Profile 
http.HandleFunc("/contact_list", indexHandler)
// + Contact List
http.HandleFunc("/team_chat", indexHandler)
// + Team Chat
http.HandleFunc("/date_planner", indexHandler)
// + Date Planner

// ,  ° . +
### _ tee time

http.HandleFunc("/location_map", indexHandler)
// + Location Map
http.HandleFunc("/schedule_calendar", indexHandler)
// + Schedule Calendar 
http.HandleFunc("/course_map", indexHandler)
// + Course Map
http.HandleFunc("/score_card", indexHandler)
// + Score Card
http.HandleFunc("/match_history", indexHandler)
// + Match History






// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
  
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }

} // ~ func main()
