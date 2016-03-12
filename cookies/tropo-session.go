// Login to Tropo.com, retrieve cookie, issue AJAX calls
package main


import (
	"net/http"
	"log"
    "encoding/json"

	"fmt"
	"io/ioutil"
	"strings"
)


func main() {
	sessionID, err := login("TROPO_USERNAME", "TROPO_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}

	session, err := next(sessionID)
    if err != nil {
		log.Fatal(err)
	}
    
    fmt.Printf("session: %v", session)
}

// SessionInfo describes a user session
// Properties "FirstName" and "AccountID" are filled if login was successful
type SessionInfo struct {
    LoggedIn    bool
    FirstName   string 
    AccountID   string 
}


func addCookie(req *http.Request, sessionID string) {
    cookie := &http.Cookie {
        Name: "JSESSIONID",
        Value: sessionID,
        Domain: "www.tropo.com",
        Path: "/",     
    }
    req.AddCookie(cookie)
}

func next(sessionID string) (*SessionInfo, error) {

	req, _ := http.NewRequest("GET", "https://www.tropo.com/sessioninfo", nil)
    addCookie(req, sessionID)
    
	res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    } 

	//fmt.Printf("res: %v", res)
	//fmt.Printf("body: %s", string(body))

    // RawSessionInfo correspond to the raw json structure returned by the AJAX call to /sessioninfo
    // ex: {"loggedIn":"true","firstName":"Steve  (5048353)","gravatarHash":"b818098073c27e8c0022a3c724886771"}
    // ex: {"loggedIn":"false"}
    type rawSessionInfo struct {
        LoggedIn string `json:"loggedIn"`
        FirstName string `json:"firstName"`
        GravatarHash string `json:"gravatarHash"`
    }
    var raw rawSessionInfo
    defer res.Body.Close()   
	body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    if err := json.Unmarshal(body, &raw); err != nil {
        return nil, err
    }
    //fmt.Printf("raw: %v", raw)

    switch raw.LoggedIn {
        case "false": // login failed
            return &SessionInfo{
                LoggedIn: false,
            }, nil
        
        case "true": // login successful
            // Extract firstname and accountID from raw string
            // ex: "Steve  (5048353)"
            splitted := strings.Split(raw.FirstName, "(")
            firstname := strings.TrimSpace(splitted[0])
            if len(splitted) != 2 {
                return nil, fmt.Errorf("could not parse firstname: %s", raw.FirstName)
            }
            till := len(splitted[1])-1
            accoundID := splitted[1][0:till]
            
            return &SessionInfo{
                    LoggedIn: true,
                    FirstName: firstname,
                    AccountID: accoundID,
            }, nil
    }
    
    return nil, fmt.Errorf("bad response, login status not recognized: %s", raw.LoggedIn)
}


// Signs in on the Tropo Web Portal and retrieves a JSESSIONID
// The JSESSIONID can be used to issue AJAX requests when encapsulated in a cookie
func login(username string, password string) (string, error) {

	payload := strings.NewReader("password="+password+"&username="+username)
	req, _ := http.NewRequest("POST", "https://www.tropo.com/login", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	statusCode := resp.StatusCode
    if statusCode != http.StatusMovedPermanently {
        return "", fmt.Errorf("status code: %d, but expected: %d", statusCode, http.StatusMovedPermanently)
    } 
    
    // Extract sessionID from location 
    // location: https://www.tropo.com/applications;jsessionid=hssn-2FB4BC666C0DE5F53C5A5EEA2B652101
    // path: /applications;jsessionid=hssn-2FB4BC666C0DE5F53C5A5EEA2B652101
	location, err := resp.Location()
	if err != nil {
		return "", err
	}
	path := location.Path
	splitted := strings.Split(path, "=")
    if len(splitted) != 2 {
        return "", fmt.Errorf("could not parse location: %s", location)
    }
	jsessionID := splitted[1]
    
	//log.Printf("session: %s", jsessionID)

	return jsessionID, nil
}
