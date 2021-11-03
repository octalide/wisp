package wisp

import (
	"fmt"
	"os"
	"testing"
)

/*
var (
	testH  *Handler
	testHB *Handler
	testHC *Handler
)
*/

func TestMain(m *testing.M) {
	Init()

	/*
		testH = &Handler{
			Callback: func(data interface{}) bool {
				fmt.Println("data:", data)
				return false
			},
			Tags:     []string{"test"},
			Blocking: false,
		}
		testHB = &Handler{
			Callback: func(data interface{}) bool {
				fmt.Println("data (B):", data)
				return false
			},
			Tags:     []string{"testb"},
			Blocking: true,
		}
		testHC = &Handler{
			Callback: func(data interface{}) bool {
				fmt.Println("data (C):", data)
				return true
			},
			Tags:     []string{"testc"},
			Blocking: true,
		}

		AddHandler(testH)
		AddHandler(testHB)
		AddHandler(testHC)
	*/

	res := m.Run()

	Stop()

	os.Exit(res)
}

func TestAddEvent(t *testing.T) {
	type args struct {
		event *Event
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "add event",
			args: args{
				event: NewEvent("test", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Broadcast(tt.args.event)
		})
	}
}

func TestAddHandler(t *testing.T) {
	type args struct {
		handler *Handler
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "add handler",
			args: args{
				handler: &Handler{
					Callback: func(e *Event) bool {
						fmt.Println("data (A):", e.Data)
						return true
					},
					Tags:     []string{"testa"},
					Blocking: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := len(Handlers())
			AddHandler(tt.args.handler)
			if start >= len(Handlers()) {
				t.Fatalf("handler size unchanged after add")
			}
		})
	}
}

func TestDelHandler(t *testing.T) {
	testHD := &Handler{}
	AddHandler(testHD)

	type args struct {
		handler *Handler
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "del handler",
			args: args{
				handler: testHD,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := len(Handlers())
			DelHandler(tt.args.handler)
			if start <= len(Handlers()) {
				t.Fatalf("handler size unchanged after del")
			}
		})
	}
}
