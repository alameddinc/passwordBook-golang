package record

import (
	"github.com/manifoldco/promptui"
	"strings"
	"time"
)

type Record struct {
	URL       string
	Username  string
	Password  string
	CreatedAt time.Time
}

func (r *Record) Init() {
	r.PromptURL()
	r.PromptUsername()
	r.PromptPassword()
	r.CreatedAt = time.Now()
}

func (r *Record) PromptURL() {
	urlFunc := func(input string) error {
		return nil
	}

	promptURL := promptui.Prompt{
		Label:    "URL",
		Validate: urlFunc,
	}

	url, _ := promptURL.Run()

	arr := strings.Split(url, "://")
	if len(arr) != 1 {
		r.URL = arr[1]
		return
	}
	r.URL = url
}

func (r *Record) PromptUsername() {
	unFunc := func(input string) error {
		return nil
	}

	promptURL := promptui.Prompt{
		Label:    "Username",
		Validate: unFunc,
	}

	r.Username, _ = promptURL.Run()
}

func (r *Record) PromptPassword() {
	passFunc := func(input string) error {
		return nil
	}

	promptURL := promptui.Prompt{
		Label:    "Password",
		Validate: passFunc,
	}

	r.Password, _ = promptURL.Run()
}
