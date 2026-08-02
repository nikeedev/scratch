package main

import (
	// built-in
	"fmt"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"io"

	// vendor
	"github.com/joho/godotenv"
)

// / For Authentication
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

///

type SessionInfo struct {
	User struct {
		Id           int    	  `json:"id"`
		Banned       bool   	  `json:"banned"`
		ShouldVpn	 bool 		  `json:"should_vpn"`
		Username     string 	  `json:"username"`
		Token        string 	  `json:"token"`
		ThumbnailUrl string 	  `json:"thumbnailUrl"`
		DateJoined   string 	  `json:"dateJoined"`
		Email        string 	  `json:"email"`
		BirthYear    int 	 	  `json:"birth_year"`
		BirthMonth   int 		  `json:"birth_month"`
		Gender       string       `json:"gender"`
		Country      string   	  `json:"country"`
		State        string 	  `json:"state"`
		MembershipAvatarBadge int `json:"membership_avatar_badge"`
	} `json:"user"`

	Permissions struct {
		Admin            bool `json:"admin"`
		Scratcher        bool `json:"scratcher"`
		NewScratcher     bool `json:"new_scrather"`
		InvitedScratcher bool `json:"invited_scratcher"`
		Social           bool `json:"social"`
		Educator         bool `json:"educator"`
		EducatorInvitee  bool `json:"educator_invitee"`
		Student          bool `json:"student"`
	} `json:"permissions"`

	Flags struct {
		MustResetPassword                bool `json:"must_reset_password"`
		MustCompleteRegistration         bool `json:"must_complete_registration"`
		HasOutstandingEmailConfirmation  bool `json:"has_outstanding_email_confirmation"`
		ShowWelcome                      bool `json:"show_welcome"`
		ConfirmEmailBanner               bool `json:"confirm_email_banner"`
		UnsupportedBrowserBanner         bool `json:"unsupported_browser_banner"`
		ProjectCommentsEnabled           bool `json:"project_comments_enabled"`
		GalleryCommentsEnabled           bool `json:"gallery_comments_enabled"`
		UserprofileCommentsEnabled       bool `json:"userprofile_comments_enabled"`
		EverythingIsTotallyNormal        bool `json:"everything_is_totally_normal"`
	} `json:"flags"`
}

type ApiInfo struct {
	id          int
	username    string
	scratchteam bool
	history     struct {
		joined string
	}

	profile struct {
		id      int
		status  string
		bio     string
		country string
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	userInfo := &User{Username: os.Getenv("s_username"), Password: os.Getenv("s_password")}

	userJSON, err := json.Marshal(userInfo)

	if err != nil {
		log.Fatal("JSON error:", err)
	}

	// log.Println(bytes.NewBuffer(userJSON))

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://scratch.mit.edu/login/", bytes.NewBuffer(userJSON))
	if err != nil {
		log.Fatal("Error: ", err)
	}

	req.Header.Add("Referer", "https://scratch.mit.edu")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("X-CSRFToken", "a")
	req.Header.Add("Cookie", "scratchcsrftoken=a;")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Error: ", err)
	}

 	sessionId := strings.Split(strings.TrimSpace(resp.Cookies()[0].Value), ";")[0]
	
	fmt.Println(sessionId)

	req, err = http.NewRequest("GET", "https://scratch.mit.edu/session/", nil)

	if err != nil {
		log.Fatal("Error: ", err)
	}
	req.Header.Add("X-CSRFToken", "a")
	req.Header.Add("Cookie", fmt.Sprintf("scratchcsrftoken=a;scratchsessionsid=${%s};", sessionId))
	req.Header.Add("Referer", "https://scratch.mit.edu")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	resp, err = client.Do(req)

	if err != nil {
		log.Fatal("Error: ", err)
	}

 	defer resp.Body.Close()

	/* var session SessionInfo

	err = json.NewDecoder(resp.Body).Decode(&session)
	if err != nil {
    	log.Fatal(err)
	}  */ 
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("JSON error:", err)
	}

	// 0644 sets standard read/write permissions for the file owner
	err = os.WriteFile("output.txt", body, 0644)
	if err != nil {
		log.Fatal("Error writing file:", err)
		return
	}
	
	fmt.Println(string(body))
}
