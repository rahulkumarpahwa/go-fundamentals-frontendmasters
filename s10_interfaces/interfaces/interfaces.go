package interfaces


// here the interface is the list of the methods and we can use this as a type.
// every struct which follows the interface needs to have this method, inside them like the SignUp() method, with definition of thier own but same signature.

type Signable interface{
	SignUp() bool; // without func keyword.
}