package builder

import "fmt"

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

type NotificationBuilder struct {
	title    string
	icon     string
	priority int
	txt      string
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.title = title
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.priority = priority
}

func (nb *NotificationBuilder) SetTxt(txt string) {
	nb.txt = txt
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	if err := nb.validate(); err != nil {
		return nil, err
	} else {
		return &Notification{
			nb.title,
			nb.icon,
			nb.priority,
			nb.txt,
		}, nil
	}
}

func (nb *NotificationBuilder) validate() error {
	if nb.priority > 5 && nb.priority < 0 {
		return fmt.Errorf("priority can not be lower than 0 or greater than 5")
	}
	return nil
}
