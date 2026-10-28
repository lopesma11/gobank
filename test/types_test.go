package main

// TO DO: install testify 
func TestNewAccount(t *Testing.T) {
	acc, err := NewAccount("Teste", "Teste", "hunter")
	assert.Nil(t, err)

	fmt.Printf("%+v\n")
}  	