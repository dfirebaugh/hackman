package modal

import "testing"

func TestPush(t *testing.T) {
	ms := &ModalStack{}
	ms.Push(Modal{
		Text:    "some text",
		Buttons: []ModalButton{},
	})
	if ms.Len() != 1 {
		t.Error("modal stack should have 1 after push")
	}
}

func TestPop(t *testing.T) {
	ms := &ModalStack{}
	text := "some pop text"
	ms.Push(Modal{
		Text: text,
	})

	if ms.Len() != 1 {
		t.Error("stack should have one element")
	}

	m := ms.Pop()

	if m.Text != text {
		t.Error("should pop off the top of the stack")
	}

	if ms.Len() != 0 {
		t.Error("should remove item from stack")
	}

	ms.Push(Modal{
		Text: "a",
	})
	ms.Push(Modal{
		Text: "b",
	})
	ms.Push(Modal{
		Text: "c",
	})

	if ms.Len() != 3 {
		t.Error("should have added 3 modals")
	}

	top := ms.Pop()
	if top.Text != "c" {
		t.Error("should have popped the last item")
	}

	if ms.Len() != 2 {
		t.Error("should have removed one element")
	}
}
