package main

// TO DO: install testify 
func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("Teste", "Teste", "132165132")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}  	