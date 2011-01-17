package list

import "testing"

func TestInsertEmptyList(t *testing.T) {
	l := New()
	if l.IsEmpty() != true {
		t.Errorf("List should be empty.")
	}
	
	l.Insert(0, "hello")
		
	if l.Length() != 1 {
		t.Errorf("List length should be 1.")
	}
	
	if l.IsEmpty() == true {
		t.Errorf("List should not be empty.")
	}
	
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\".")
	}
	
}

func TestInsertMultipleElements(t *testing.T) {
	l := New()
	l.Insert(0, "Apple")
	l.Insert(1, "Banana")
	l.Insert(2, "Orange")
	
	if l.Length() != 3 {
		t.Errorf("List size should be three.")
	}
	if l.Get(0).Value != "Apple" {
		t.Errorf("1st element should be \"Apple\".")
	}
	if l.Get(1).Value != "Banana" {
		t.Errorf("2nd element should be \"Banana\".")
	}
	if l.Get(2).Value != "Orange" {
		t.Errorf("3rd element should be \"Orange\".")
	}
}

func TestInsertBetweenElements(t *testing.T) {
	l := New()
	l.Insert(0, "Apple")
	l.Insert(1, "Banana")
	l.Insert(1, "Orange")
	
	if l.Length() != 3 {
		t.Errorf("List size should be three.")
	}
	
	if l.Get(0).Value != "Apple" {
		t.Errorf("1st element should be \"Apple\".")
	}
	
	if l.Get(1).Value != "Orange" {
		t.Errorf("2nd element should be \"Orange\".")
 	}
	
	if l.Get(2).Value != "Banana" {
		t.Errorf("3rd element should be \"Banana\".")
	}
}



func TestInsertBeforeFirstElement(t *testing.T) {
	l := New()
	l.Insert(0, "Apple")
	l.Insert(0, "Banana")
	
	if l.Length() != 2 {
		t.Errorf("List size should be two.")
	}
	
	if l.Get(0).Value != "Banana" {
		t.Errorf("1st element should be \"Banana\".")
	}
	if l.Get(1).Value != "Apple" {
		t.Errorf("1st element should be \"Apple\".")
	}
}


func TestInsertAfterLastElement(t *testing.T) {
	l := New()
	l.Insert(0, "Apple")
	l.Insert(1, "Banana")
	
	if l.Length() != 2 {
		t.Errorf("List size should be two.")
	}
	
	if l.Get(0).Value != "Apple" {
		t.Errorf("1st element should be \"Apple\".")
	}
	if l.Get(1).Value != "Banana" {
		t.Errorf("2nd element should be \"Banana\".")
	}
}

func TestInsertOutOfBounds(t *testing.T) {
	l := New()
	
	defer func() {
		if r:= recover(); r!= nil {
			//expected 
            //fmt.Println("Recovered in TestInsertOutOfBounds()")
		}
	}()
	
	l.Insert(-1, "Hello")
	
}

func TestAdd(t *testing.T) {
	l := New()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	
	if l.Length() != 3 {
		t.Errorf("List length should be 3.")
	}
	if l.Get(0).Value != 1 {
		t.Errorf("1st element should be 1.")
	}
	if l.Get(1).Value != 2 {
		t.Errorf("2nd element should be 2.")
	}
	if l.Get(2).Value != 3 {
		t.Errorf("3rd element should be 3.")
	}
}


func TestSet(t *testing.T) {
	l := New()
	l.Insert(0, "hello")
	
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\"")
	}
	
	l.Set(0, "world")
	
	if l.Get(0).Value != "world" {
		t.Errorf("1st element should now be \"world\"")
	}
}

func TestDeleteOnlyElement(t *testing.T) {
	l := New()
	l.Add("hello")
	
	if l.Length() != 1 {
		t.Errorf("List length should be 1.")
	}
	
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\".")
	}
	
	l.Delete(0)
	if l.Length() != 0 {
		t.Errorf("The list length should be 0")
	}

}

func TestDeleteFirstElement(t *testing.T) {
	l := New()
	l.Add("hello")
	l.Add("world")
	l.Add("hello")

	if l.Length() != 3 {
		t.Errorf("List length should be 3.")
	}
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\".")
	}
	if l.Get(1).Value != "world" {
		t.Errorf("2nd element should be \"world\".")
	}
	if l.Get(2).Value != "hello" {
		t.Errorf("3rd element should be \"hello\".")
	}
		
	l.Delete(0)
	if l.Length() != 2 {
		t.Errorf("List length should be 2.")
	}
		
	if l.Get(0).Value != "world" {
		t.Errorf("1st element should be \"world\".")
	}
		
	if l.Get(1).Value != "hello" {
		t.Errorf("2nd element should be \"hello\".")
	}
}

func TestDeleteLastElement(t *testing.T) {
	l := New()
	
	l.Add("hello")
	l.Add("world")
	l.Add("again")
	
	if l.Length() != 3 {
		t.Errorf("List length should be 3.")
	}
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\".")
	}
	if l.Get(1).Value != "world" {
		t.Errorf("2nd element should be \"world\".")
	}
	if l.Get(2).Value != "again" {
		t.Errorf("3rd element should be \"again\".")
	}
	
	l.Delete(2)
	
	if l.Length() != 2 {
		t.Errorf("List length should be 2.")
	}
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\".")
	}
	if l.Get(1).Value != "world" {
		t.Errorf("2nd element should be \"world\".")
	}
	
}

func TestDeleteMiddleElement(t *testing.T) {
	l := New()
	
	l.Add("hello")
	l.Add("world")
	l.Add("again")
	
	if l.Length() != 3 {
		t.Errorf("List length should be 3.")
	}
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\".")
	}
	if l.Get(1).Value != "world" {
		t.Errorf("2nd element should be \"world\".")
	}
	if l.Get(2).Value != "again" {
		t.Errorf("3rd element should be \"again\".")
	}
	
	l.Delete(1)
	
	if l.Length() != 2 {
		t.Errorf("List length should be 2.")
	}
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\".")
	}
	if l.Get(1).Value != "again" {
		t.Errorf("2nd element should be \"again\".")
	}
	
}

func TestDeleteOutOfBounds(t *testing.T) {
	l := New()
	
	defer func() {
		if r:= recover(); r!= nil {
			//expected
            //fmt.Println("Recovered in TestDeleteOutOfBounds()")
			
		}
	}()
	l.Delete(-1)
}

func TestDeleteByValue(t *testing.T) {
	l := New()
	l.Add("hello")
	l.Add("world")
	l.Add("hello")
	
	if l.Length() != 3 {
		t.Errorf("List length should be 3.")
	}
	
	if l.Get(0).Value != "hello" {
		t.Errorf("1st element should be \"hello\".")
	}
	if l.Get(1).Value != "world" {
		t.Errorf("2nd element should be \"world\".")
	}	
	if l.Get(2).Value != "hello" {
		t.Errorf("3rd element should be \"hello\".")
	}
	
	l.DeleteByVal("hello")
	
	if l.Length() != 2 {
		t.Errorf("list length should be 2.")
	}
	
	if l.Get(0).Value != "world" {
		t.Errorf("1st element should be \"world\".")
	}
	if l.Get(1).Value != "hello" {
		t.Errorf("2nd element should be \"hello\".")
	}
}
