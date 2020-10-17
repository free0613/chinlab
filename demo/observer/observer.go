package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Updatemsg(msg string)
}

type Subject struct {
	obersers []IObserver
}

func (sub *Subject) Register(observer IObserver) {
	sub.obersers = append(sub.obersers, observer)
}

func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.obersers {
		if ob == observer {
			sub.obersers = append(sub.obersers[:i], sub.obersers[i+1:]...)
		}
	}
}

func (sub *Subject) Notify(msg string) {
	for _, ob := range sub.obersers {
		ob.Updatemsg(msg)
	}
}

type Observer1 struct{}
type Observer2 struct{}

func (ob Observer1) Updatemsg(msg string) {
	fmt.Printf("recive1 msg: %s \n", msg)
}

func (Observer2) Updatemsg(msg string) {
	fmt.Printf("recive2 msg: %s \n", msg)
}
