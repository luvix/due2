package main

//StringArray is string array strcut.
// from https://stackoverflow.com/a/28323276
type StringArray []string

func (sa *StringArray) String() string {
	return "StringArray"
}

//Set is append a string to string array.
func (sa *StringArray) Set(arg string) error {
	*sa = append(*sa, arg)

	return nil
}
